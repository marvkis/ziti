// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/rest_model"
)

// CreateServiceReader is a Reader for the CreateService structure.
type CreateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateServiceCreated creates a CreateServiceCreated with default headers values
func NewCreateServiceCreated() *CreateServiceCreated {
	return &CreateServiceCreated{}
}

/*
CreateServiceCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateServiceCreated struct {
	Payload *rest_model.CreateEnvelope
}

// IsSuccess returns true when this create service created response has a 2xx status code
func (o *CreateServiceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service created response has a 3xx status code
func (o *CreateServiceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service created response has a 4xx status code
func (o *CreateServiceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service created response has a 5xx status code
func (o *CreateServiceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create service created response a status code equal to that given
func (o *CreateServiceCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create service created response
func (o *CreateServiceCreated) Code() int {
	return 201
}

func (o *CreateServiceCreated) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceCreated  %+v", 201, o.Payload)
}

func (o *CreateServiceCreated) String() string {
	return fmt.Sprintf("[POST /services][%d] createServiceCreated  %+v", 201, o.Payload)
}

func (o *CreateServiceCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateServiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceBadRequest creates a CreateServiceBadRequest with default headers values
func NewCreateServiceBadRequest() *CreateServiceBadRequest {
	return &CreateServiceBadRequest{}
}

/*
CreateServiceBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateServiceBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create service bad request response has a 2xx status code
func (o *CreateServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service bad request response has a 3xx status code
func (o *CreateServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service bad request response has a 4xx status code
func (o *CreateServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service bad request response has a 5xx status code
func (o *CreateServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create service bad request response a status code equal to that given
func (o *CreateServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create service bad request response
func (o *CreateServiceBadRequest) Code() int {
	return 400
}

func (o *CreateServiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceBadRequest) String() string {
	return fmt.Sprintf("[POST /services][%d] createServiceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceUnauthorized creates a CreateServiceUnauthorized with default headers values
func NewCreateServiceUnauthorized() *CreateServiceUnauthorized {
	return &CreateServiceUnauthorized{}
}

/*
CreateServiceUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type CreateServiceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create service unauthorized response has a 2xx status code
func (o *CreateServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service unauthorized response has a 3xx status code
func (o *CreateServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service unauthorized response has a 4xx status code
func (o *CreateServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service unauthorized response has a 5xx status code
func (o *CreateServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create service unauthorized response a status code equal to that given
func (o *CreateServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create service unauthorized response
func (o *CreateServiceUnauthorized) Code() int {
	return 401
}

func (o *CreateServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateServiceUnauthorized) String() string {
	return fmt.Sprintf("[POST /services][%d] createServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateServiceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
