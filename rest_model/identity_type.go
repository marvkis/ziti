// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IdentityType identity type
//
// swagger:model identityType
type IdentityType string

func NewIdentityType(value IdentityType) *IdentityType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IdentityType.
func (m IdentityType) Pointer() *IdentityType {
	return &m
}

const (

	// IdentityTypeUser captures enum value "User"
	IdentityTypeUser IdentityType = "User"

	// IdentityTypeDevice captures enum value "Device"
	IdentityTypeDevice IdentityType = "Device"

	// IdentityTypeService captures enum value "Service"
	IdentityTypeService IdentityType = "Service"

	// IdentityTypeRouter captures enum value "Router"
	IdentityTypeRouter IdentityType = "Router"
)

// for schema
var identityTypeEnum []interface{}

func init() {
	var res []IdentityType
	if err := json.Unmarshal([]byte(`["User","Device","Service","Router"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		identityTypeEnum = append(identityTypeEnum, v)
	}
}

func (m IdentityType) validateIdentityTypeEnum(path, location string, value IdentityType) error {
	if err := validate.EnumCase(path, location, value, identityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this identity type
func (m IdentityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIdentityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this identity type based on context it is used
func (m IdentityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
