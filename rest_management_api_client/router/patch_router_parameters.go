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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// NewPatchRouterParams creates a new PatchRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRouterParams() *PatchRouterParams {
	return &PatchRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRouterParamsWithTimeout creates a new PatchRouterParams object
// with the ability to set a timeout on a request.
func NewPatchRouterParamsWithTimeout(timeout time.Duration) *PatchRouterParams {
	return &PatchRouterParams{
		timeout: timeout,
	}
}

// NewPatchRouterParamsWithContext creates a new PatchRouterParams object
// with the ability to set a context for a request.
func NewPatchRouterParamsWithContext(ctx context.Context) *PatchRouterParams {
	return &PatchRouterParams{
		Context: ctx,
	}
}

// NewPatchRouterParamsWithHTTPClient creates a new PatchRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRouterParamsWithHTTPClient(client *http.Client) *PatchRouterParams {
	return &PatchRouterParams{
		HTTPClient: client,
	}
}

/* PatchRouterParams contains all the parameters to send to the API endpoint
   for the patch router operation.

   Typically these are written to a http.Request.
*/
type PatchRouterParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* Router.

	   A router patch object
	*/
	Router *rest_model.RouterPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRouterParams) WithDefaults() *PatchRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch router params
func (o *PatchRouterParams) WithTimeout(timeout time.Duration) *PatchRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch router params
func (o *PatchRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch router params
func (o *PatchRouterParams) WithContext(ctx context.Context) *PatchRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch router params
func (o *PatchRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch router params
func (o *PatchRouterParams) WithHTTPClient(client *http.Client) *PatchRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch router params
func (o *PatchRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch router params
func (o *PatchRouterParams) WithID(id string) *PatchRouterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch router params
func (o *PatchRouterParams) SetID(id string) {
	o.ID = id
}

// WithRouter adds the router to the patch router params
func (o *PatchRouterParams) WithRouter(router *rest_model.RouterPatch) *PatchRouterParams {
	o.SetRouter(router)
	return o
}

// SetRouter adds the router to the patch router params
func (o *PatchRouterParams) SetRouter(router *rest_model.RouterPatch) {
	o.Router = router
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Router != nil {
		if err := r.SetBodyParam(o.Router); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
