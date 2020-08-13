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

package routes

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/michaelquigley/pfxlog"
	"github.com/mitchellh/mapstructure"
	"github.com/openziti/edge/controller/apierror"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/internal/permissions"
	"github.com/openziti/edge/controller/model"
	"github.com/openziti/edge/controller/response"
	"github.com/openziti/edge/rest_model"
	"github.com/openziti/edge/rest_server/operations/authentication"
	"github.com/openziti/foundation/util/stringz"
	"net/http"
	"time"
)

func init() {
	r := NewAuthRouter()
	env.AddRouter(r)
}

type AuthRouter struct {
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (ro *AuthRouter) Register(ae *env.AppEnv) {
	ae.Api.AuthenticationAuthenticateHandler = authentication.AuthenticateHandlerFunc(func(params authentication.AuthenticateParams) middleware.Responder {
		return ae.IsAllowed(func(ae *env.AppEnv, rc *response.RequestContext) { ro.authHandler(ae, rc, params) }, params.HTTPRequest, "", "", permissions.Always())
	})
}

func (ro *AuthRouter) authHandler(ae *env.AppEnv, rc *response.RequestContext, params authentication.AuthenticateParams) {
	logger := pfxlog.Logger()

	startTime := time.Now()

	authContext := model.NewAuthContextHttp(params.HTTPRequest, params.Method, params.Body)

	identity, err := ae.Handlers.Authenticator.IsAuthorized(authContext)

	if err != nil {
		rc.RespondWithError(err)
		return
	}

	if identity == nil {
		rc.RespondWithApiError(apierror.NewUnauthorized())
		return
	}

	if identity.EnvInfo == nil {
		identity.EnvInfo = &model.EnvInfo{}
	}

	if identity.SdkInfo == nil {
		identity.SdkInfo = &model.SdkInfo{}
	}

	if dataMap := authContext.GetData(); dataMap != nil {
		shouldUpdate := false

		if envInfoInterface := dataMap["envInfo"]; envInfoInterface != nil {
			if envInfo := envInfoInterface.(map[string]interface{}); envInfo != nil {
				if err := mapstructure.Decode(envInfo, &identity.EnvInfo); err != nil {
					logger.WithError(err).Error("error processing env info")
				}
				shouldUpdate = true
			}
		}

		if sdkInfoInterface := dataMap["sdkInfo"]; sdkInfoInterface != nil {
			if sdkInfo := sdkInfoInterface.(map[string]interface{}); sdkInfo != nil {
				if err := mapstructure.Decode(sdkInfo, &identity.SdkInfo); err != nil {
					logger.WithError(err).Error("error processing sdk info")
				}
				shouldUpdate = true
			}
		}

		if shouldUpdate {
			beforePathInfo := time.Now()
			if err := ae.GetHandlers().Identity.PatchInfo(identity); err != nil {
				logger.WithError(err).Errorf("failed to update sdk/env info on identity [%s] auth", identity.Id)
			}
			patchInfoDelta := time.Now().Sub(beforePathInfo)
			logger.Debugf("auth patch info timing (micro-sec): %v\n", patchInfoDelta.Microseconds())
		}
	}

	token := uuid.New().String()
	configTypes := map[string]struct{}{}

	if params.Body != nil {
		configTypes = mapConfigTypeNamesToIds(ae, params.Body.ConfigTypes, identity.Id)
	}

	logger.Debugf("client %v requesting configTypes: %v", identity.Name, configTypes)
	s := &model.ApiSession{
		IdentityId:  identity.Id,
		Token:       token,
		ConfigTypes: configTypes,
	}
	beforeSessionCreate := time.Now()
	sessionId, err := ae.Handlers.ApiSession.Create(s)

	if err != nil {
		rc.RespondWithError(err)
		return
	}
	sessionCreateDuration := time.Now().Sub(beforeSessionCreate)
	logger.Debugf("auth create session timing (micro-sec): %v\n", sessionCreateDuration.Microseconds())

	session, err := ae.Handlers.ApiSession.Read(sessionId)

	if err != nil {
		logger.WithField("cause", err).Error("loading session by id resulted in an error")
		rc.RespondWithApiError(apierror.NewUnauthorized())
	}

	apiSession := MapToCurrentApiSessionRestModel(session, ae.Config.SessionTimeoutDuration())

	envelope := &rest_model.CurrentAPISessionDetailEnvelope{Data: apiSession, Meta: &rest_model.Meta{}}

	expiration := time.Time(*apiSession.ExpiresAt)
	cookie := http.Cookie{Name: ae.AuthCookieName, Value: token, Expires: expiration}

	rc.ResponseWriter.Header().Set(ae.AuthHeaderName, session.Token)
	http.SetCookie(rc.ResponseWriter, &cookie)

	rc.Respond(envelope, http.StatusOK)

	taken := time.Now().Sub(startTime).Microseconds()
	logger.Debugf("auth timing (micro-sec): %v\n", taken)
}

func mapConfigTypeNamesToIds(ae *env.AppEnv, values []string, identityId string) map[string]struct{} {
	var result []string
	if stringz.Contains(values, "all") {
		result = []string{"all"}
	} else {
		for _, val := range values {
			if configType, _ := ae.GetHandlers().ConfigType.Read(val); configType != nil {
				result = append(result, val)
			} else if configType, _ := ae.GetHandlers().ConfigType.ReadByName(val); configType != nil {
				result = append(result, configType.Id)
			} else {
				pfxlog.Logger().Debugf("user %v submitted %v as a config type of interest, but no matching records found", identityId, val)
			}
		}
	}
	return stringz.SliceToSet(result)
}
