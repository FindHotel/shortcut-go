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

// NewCreateStoryParams creates a new CreateStoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStoryParams() *CreateStoryParams {
	return &CreateStoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStoryParamsWithTimeout creates a new CreateStoryParams object
// with the ability to set a timeout on a request.
func NewCreateStoryParamsWithTimeout(timeout time.Duration) *CreateStoryParams {
	return &CreateStoryParams{
		timeout: timeout,
	}
}

// NewCreateStoryParamsWithContext creates a new CreateStoryParams object
// with the ability to set a context for a request.
func NewCreateStoryParamsWithContext(ctx context.Context) *CreateStoryParams {
	return &CreateStoryParams{
		Context: ctx,
	}
}

// NewCreateStoryParamsWithHTTPClient creates a new CreateStoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStoryParamsWithHTTPClient(client *http.Client) *CreateStoryParams {
	return &CreateStoryParams{
		HTTPClient: client,
	}
}

/*
CreateStoryParams contains all the parameters to send to the API endpoint

	for the create story operation.

	Typically these are written to a http.Request.
*/
type CreateStoryParams struct {

	/* CreateStoryParams.

	   Request parameters for creating a story.
	*/
	CreateStoryParams *models.CreateStoryParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create story params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryParams) WithDefaults() *CreateStoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create story params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create story params
func (o *CreateStoryParams) WithTimeout(timeout time.Duration) *CreateStoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create story params
func (o *CreateStoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create story params
func (o *CreateStoryParams) WithContext(ctx context.Context) *CreateStoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create story params
func (o *CreateStoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create story params
func (o *CreateStoryParams) WithHTTPClient(client *http.Client) *CreateStoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create story params
func (o *CreateStoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateStoryParams adds the createStoryParams to the create story params
func (o *CreateStoryParams) WithCreateStoryParams(createStoryParams *models.CreateStoryParams) *CreateStoryParams {
	o.SetCreateStoryParams(createStoryParams)
	return o
}

// SetCreateStoryParams adds the createStoryParams to the create story params
func (o *CreateStoryParams) SetCreateStoryParams(createStoryParams *models.CreateStoryParams) {
	o.CreateStoryParams = createStoryParams
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateStoryParams != nil {
		if err := r.SetBodyParam(o.CreateStoryParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}