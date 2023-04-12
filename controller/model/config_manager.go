/*
	Copyright NetFoundry Inc.

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
	"encoding/json"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/edge/pb/edge_cmd_pb"
	"github.com/openziti/fabric/controller/change"
	"github.com/openziti/fabric/controller/command"
	"github.com/openziti/fabric/controller/fields"
	"github.com/openziti/fabric/controller/models"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/storage/boltz"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
	"strings"
)

func NewConfigManager(env Env) *ConfigManager {
	manager := &ConfigManager{
		baseEntityManager: newBaseEntityManager[*Config, *persistence.Config](env, env.GetStores().Config),
	}
	manager.impl = manager

	network.RegisterManagerDecoder[*Config](env.GetHostController().GetNetwork().Managers, manager)

	return manager
}

type ConfigManager struct {
	baseEntityManager[*Config, *persistence.Config]
}

func (self *ConfigManager) newModelEntity() *Config {
	return &Config{}
}

func (self *ConfigManager) Create(entity *Config, ctx *change.Context) error {
	return network.DispatchCreate[*Config](self, entity, ctx)
}

func (self *ConfigManager) ApplyCreate(cmd *command.CreateEntityCommand[*Config]) error {
	_, err := self.createEntity(cmd.Entity, cmd.Context)
	return err
}

func (self *ConfigManager) Update(entity *Config, checker fields.UpdatedFields, ctx *change.Context) error {
	return network.DispatchUpdate[*Config](self, entity, checker, ctx)
}

func (self *ConfigManager) ApplyUpdate(cmd *command.UpdateEntityCommand[*Config]) error {
	var checker boltz.FieldChecker = self
	if cmd.UpdatedFields != nil {
		checker = &AndFieldChecker{first: self, second: cmd.UpdatedFields}
	}
	return self.updateEntity(cmd.Entity, checker, cmd.Context)
}

func (self *ConfigManager) Read(id string) (*Config, error) {
	modelEntity := &Config{}
	if err := self.readEntity(id, modelEntity); err != nil {
		return nil, err
	}
	return modelEntity, nil
}

func (self *ConfigManager) readInTx(tx *bbolt.Tx, id string) (*Config, error) {
	modelEntity := &Config{}
	if err := self.readEntityInTx(tx, id, modelEntity); err != nil {
		return nil, err
	}
	return modelEntity, nil
}

func (self *ConfigManager) IsUpdated(field string) bool {
	return !strings.EqualFold(field, "type")
}

func (self *ConfigManager) Marshall(entity *Config) ([]byte, error) {
	tags, err := edge_cmd_pb.EncodeTags(entity.Tags)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(entity.Data)
	if err != nil {
		return nil, err
	}

	msg := &edge_cmd_pb.Config{
		Id:           entity.Id,
		Name:         entity.Name,
		ConfigTypeId: entity.TypeId,
		Data:         data,
		Tags:         tags,
	}

	return proto.Marshal(msg)
}

func (self *ConfigManager) Unmarshall(bytes []byte) (*Config, error) {
	msg := &edge_cmd_pb.Config{}
	if err := proto.Unmarshal(bytes, msg); err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		return nil, err
	}

	return &Config{
		BaseEntity: models.BaseEntity{
			Id:   msg.Id,
			Tags: edge_cmd_pb.DecodeTags(msg.Tags),
		},
		Name:   msg.Name,
		TypeId: msg.ConfigTypeId,
		Data:   data,
	}, nil
}
