// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/FindHotel/shortcut-go/models"
)

// NewCreateGroupParams creates a new CreateGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGroupParams() *CreateGroupParams {
	return &CreateGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGroupParamsWithTimeout creates a new CreateGroupParams object
// with the ability to set a timeout on a request.
func NewCreateGroupParamsWithTimeout(timeout time.Duration) *CreateGroupParams {
	return &CreateGroupParams{
		timeout: timeout,
	}
}

// NewCreateGroupParamsWithContext creates a new CreateGroupParams object
// with the ability to set a context for a request.
func NewCreateGroupParamsWithContext(ctx context.Context) *CreateGroupParams {
	return &CreateGroupParams{
		Context: ctx,
	}
}

// NewCreateGroupParamsWithHTTPClient creates a new CreateGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGroupParamsWithHTTPClient(client *http.Client) *CreateGroupParams {
	return &CreateGroupParams{
		HTTPClient: client,
	}
}

/*
CreateGroupParams contains all the parameters to send to the API endpoint

	for the create group operation.

	Typically these are written to a http.Request.
*/
type CreateGroupParams struct {

	// CreateGroup.
	CreateGroup *models.CreateGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGroupParams) WithDefaults() *CreateGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create group params
func (o *CreateGroupParams) WithTimeout(timeout time.Duration) *CreateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create group params
func (o *CreateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create group params
func (o *CreateGroupParams) WithContext(ctx context.Context) *CreateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create group params
func (o *CreateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) WithHTTPClient(client *http.Client) *CreateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateGroup adds the createGroup to the create group params
func (o *CreateGroupParams) WithCreateGroup(createGroup *models.CreateGroup) *CreateGroupParams {
	o.SetCreateGroup(createGroup)
	return o
}

// SetCreateGroup adds the createGroup to the create group params
func (o *CreateGroupParams) SetCreateGroup(createGroup *models.CreateGroup) {
	o.CreateGroup = createGroup
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateGroup != nil {
		if err := r.SetBodyParam(o.CreateGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
