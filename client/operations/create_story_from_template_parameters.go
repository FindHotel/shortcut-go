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

// NewCreateStoryFromTemplateParams creates a new CreateStoryFromTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStoryFromTemplateParams() *CreateStoryFromTemplateParams {
	return &CreateStoryFromTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStoryFromTemplateParamsWithTimeout creates a new CreateStoryFromTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateStoryFromTemplateParamsWithTimeout(timeout time.Duration) *CreateStoryFromTemplateParams {
	return &CreateStoryFromTemplateParams{
		timeout: timeout,
	}
}

// NewCreateStoryFromTemplateParamsWithContext creates a new CreateStoryFromTemplateParams object
// with the ability to set a context for a request.
func NewCreateStoryFromTemplateParamsWithContext(ctx context.Context) *CreateStoryFromTemplateParams {
	return &CreateStoryFromTemplateParams{
		Context: ctx,
	}
}

// NewCreateStoryFromTemplateParamsWithHTTPClient creates a new CreateStoryFromTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStoryFromTemplateParamsWithHTTPClient(client *http.Client) *CreateStoryFromTemplateParams {
	return &CreateStoryFromTemplateParams{
		HTTPClient: client,
	}
}

/*
CreateStoryFromTemplateParams contains all the parameters to send to the API endpoint

	for the create story from template operation.

	Typically these are written to a http.Request.
*/
type CreateStoryFromTemplateParams struct {

	/* CreateStoryFromTemplateParams.

	   Request parameters for creating a story from a story template. These parameters are merged with the values derived from the template.
	*/
	CreateStoryFromTemplateParams *models.CreateStoryFromTemplateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create story from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryFromTemplateParams) WithDefaults() *CreateStoryFromTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create story from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoryFromTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create story from template params
func (o *CreateStoryFromTemplateParams) WithTimeout(timeout time.Duration) *CreateStoryFromTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create story from template params
func (o *CreateStoryFromTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create story from template params
func (o *CreateStoryFromTemplateParams) WithContext(ctx context.Context) *CreateStoryFromTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create story from template params
func (o *CreateStoryFromTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create story from template params
func (o *CreateStoryFromTemplateParams) WithHTTPClient(client *http.Client) *CreateStoryFromTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create story from template params
func (o *CreateStoryFromTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateStoryFromTemplateParams adds the createStoryFromTemplateParams to the create story from template params
func (o *CreateStoryFromTemplateParams) WithCreateStoryFromTemplateParams(createStoryFromTemplateParams *models.CreateStoryFromTemplateParams) *CreateStoryFromTemplateParams {
	o.SetCreateStoryFromTemplateParams(createStoryFromTemplateParams)
	return o
}

// SetCreateStoryFromTemplateParams adds the createStoryFromTemplateParams to the create story from template params
func (o *CreateStoryFromTemplateParams) SetCreateStoryFromTemplateParams(createStoryFromTemplateParams *models.CreateStoryFromTemplateParams) {
	o.CreateStoryFromTemplateParams = createStoryFromTemplateParams
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStoryFromTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateStoryFromTemplateParams != nil {
		if err := r.SetBodyParam(o.CreateStoryFromTemplateParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}