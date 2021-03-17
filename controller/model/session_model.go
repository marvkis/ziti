/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package model

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/apierror"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/fabric/controller/models"
	"github.com/openziti/foundation/storage/boltz"
	"github.com/openziti/foundation/util/errorz"
	"github.com/openziti/foundation/util/stringz"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"reflect"
	"time"
)

type Session struct {
	models.BaseEntity
	Token        string
	ApiSessionId string
	ServiceId    string
	Type         string
	SessionCerts []*SessionCert
}

type SessionCert struct {
	Cert        string
	Fingerprint string
	ValidFrom   time.Time
	ValidTo     time.Time
}

func (entity *Session) toBoltEntityForCreate(tx *bbolt.Tx, handler Handler) (boltz.Entity, error) {
	apiSession, err := handler.GetEnv().GetStores().ApiSession.LoadOneById(tx, entity.ApiSessionId)
	if err != nil {
		return nil, err
	}
	if apiSession == nil {
		return nil, errorz.NewFieldError("api session not found", "ApiSessionId", entity.ApiSessionId)
	}

	service, err := handler.GetEnv().GetHandlers().EdgeService.ReadForIdentityInTx(tx, entity.ServiceId, apiSession.IdentityId, nil)
	if err != nil {
		return nil, err
	}

	if entity.Type == "" {
		entity.Type = persistence.SessionTypeDial
	}

	if persistence.SessionTypeDial == entity.Type && !stringz.Contains(service.Permissions, persistence.PolicyTypeDialName) {
		return nil, errorz.NewFieldError("service not found", "ServiceId", entity.ServiceId)
	}

	if persistence.SessionTypeBind == entity.Type && !stringz.Contains(service.Permissions, persistence.PolicyTypeBindName) {
		return nil, errorz.NewFieldError("service not found", "ServiceId", entity.ServiceId)
	}

	checkCache := map[string]bool{} //cache individual check status
	validPosture := false
	hasMatchingPolicies := false

	postureCheckMap := handler.GetEnv().GetHandlers().EdgeService.GetPostureChecks(apiSession.IdentityId, entity.ServiceId)

	for policyId, postureChecks := range postureCheckMap {
		policy, err := handler.GetEnv().GetHandlers().ServicePolicy.Read(policyId)

		if err != nil {
			continue
		}

		if policy.PolicyType != entity.Type {
			continue
		}
		hasMatchingPolicies = true
		isPolicyPassing := true

		for _, postureCheck := range postureChecks {

			isCheckPassing := true
			found := false
			if isCheckPassing, found = checkCache[postureCheck.Id]; !found {
				isCheckPassing = handler.GetEnv().GetHandlers().PostureResponse.Evaluate(apiSession.IdentityId, apiSession.Id, postureCheck)
				checkCache[postureCheck.Id] = isCheckPassing
			}

			if !isCheckPassing {
				isPolicyPassing = false //failed, move to next policy
				break
			}
		}
		if isPolicyPassing {
			validPosture = true
			break
		}
	}

	if hasMatchingPolicies && !validPosture {
		var policyIds []string

		dataMap := map[string]interface{}{}

		for policyId, checks := range postureCheckMap {
			policyIds = append(policyIds, policyId)

			var queries []interface{}

			for _, check := range checks {
				queries = append(queries, map[string]interface{}{
					"id":        check.Id,
					"typeId":    check.TypeId,
					"isPassing": checkCache[check.Id],
				})
			}

			dataMap[policyId] = queries
		}

		cause := apierror.GenericCauseError{
			Message: fmt.Sprintf("Failed to pass posture checks for service policies: %v", policyIds),
			DataMap: dataMap,
		}
		return nil, apierror.NewInvalidPosture(cause)
	}

	maxRows := 1
	result, err := handler.GetEnv().GetHandlers().EdgeRouter.ListForIdentityAndServiceWithTx(tx, apiSession.IdentityId, entity.ServiceId, &maxRows)
	if err != nil {
		return nil, err
	}
	if result.Count < 1 {
		return nil, apierror.NewNoEdgeRoutersAvailable()
	}

	boltEntity := &persistence.Session{
		BaseExtEntity: *boltz.NewExtEntity(entity.Id, entity.Tags),
		Token:         entity.Token,
		ApiSessionId:  entity.ApiSessionId,
		ServiceId:     entity.ServiceId,
		Type:          entity.Type,
		ApiSession:    apiSession,
	}

	identity, err := handler.GetEnv().GetStores().Identity.LoadOneById(tx, apiSession.IdentityId)

	if err != nil {
		return nil, err
	}

	fingerprints := map[string]string{}

	for _, authenticatorId := range identity.Authenticators {
		authenticator, err := handler.GetEnv().GetStores().Authenticator.LoadOneById(tx, authenticatorId)
		if err != nil {
			pfxlog.Logger().Errorf("encountered error retrieving fingerprints for authenticator [%s]", authenticatorId)
			continue
		}
		if certAuth := authenticator.ToCert(); certAuth != nil {
			fingerprints[certAuth.Fingerprint] = certAuth.Pem
		}
	}

	for fingerprint, cert := range fingerprints {
		validFrom := time.Now()
		validTo := time.Now().AddDate(1, 0, 0)

		boltEntity.Certs = append(boltEntity.Certs, &persistence.SessionCert{
			Cert:        cert,
			Fingerprint: fingerprint,
			ValidFrom:   validFrom,
			ValidTo:     validTo,
		})
	}

	return boltEntity, nil
}

func (entity *Session) toBoltEntityForUpdate(*bbolt.Tx, Handler) (boltz.Entity, error) {
	return &persistence.Session{
		BaseExtEntity: *boltz.NewExtEntity(entity.Id, entity.Tags),
		Token:         entity.Token,
		ApiSessionId:  entity.ApiSessionId,
		ServiceId:     entity.ServiceId,
		Type:          entity.Type,
	}, nil
}

func (entity *Session) toBoltEntityForPatch(tx *bbolt.Tx, handler Handler, checker boltz.FieldChecker) (boltz.Entity, error) {
	return entity.toBoltEntityForUpdate(tx, handler)
}

func (entity *Session) fillFrom(_ Handler, _ *bbolt.Tx, boltEntity boltz.Entity) error {
	boltSession, ok := boltEntity.(*persistence.Session)
	if !ok {
		return errors.Errorf("unexpected type %v when filling model Session", reflect.TypeOf(boltEntity))
	}
	entity.FillCommon(boltSession)
	entity.Token = boltSession.Token
	entity.ApiSessionId = boltSession.ApiSessionId
	entity.ServiceId = boltSession.ServiceId
	entity.Type = boltSession.Type
	return nil
}

func (entity *SessionCert) FillFrom(_ Handler, _ *bbolt.Tx, boltEntity boltz.Entity) error {
	boltSessionCert, ok := boltEntity.(*persistence.SessionCert)
	if !ok {
		return errors.Errorf("unexpected type %v when filling model SessionCert", reflect.TypeOf(boltEntity))
	}
	entity.Fingerprint = boltSessionCert.Fingerprint
	entity.Cert = boltSessionCert.Cert
	entity.ValidFrom = boltSessionCert.ValidFrom
	entity.ValidTo = boltSessionCert.ValidTo
	return nil
}
