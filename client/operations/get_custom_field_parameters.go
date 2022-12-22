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
)

// NewGetCustomFieldParams creates a new GetCustomFieldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomFieldParams() *GetCustomFieldParams {
	return &GetCustomFieldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomFieldParamsWithTimeout creates a new GetCustomFieldParams object
// with the ability to set a timeout on a request.
func NewGetCustomFieldParamsWithTimeout(timeout time.Duration) *GetCustomFieldParams {
	return &GetCustomFieldParams{
		timeout: timeout,
	}
}

// NewGetCustomFieldParamsWithContext creates a new GetCustomFieldParams object
// with the ability to set a context for a request.
func NewGetCustomFieldParamsWithContext(ctx context.Context) *GetCustomFieldParams {
	return &GetCustomFieldParams{
		Context: ctx,
	}
}

// NewGetCustomFieldParamsWithHTTPClient creates a new GetCustomFieldParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomFieldParamsWithHTTPClient(client *http.Client) *GetCustomFieldParams {
	return &GetCustomFieldParams{
		HTTPClient: client,
	}
}

/*
GetCustomFieldParams contains all the parameters to send to the API endpoint

	for the get custom field operation.

	Typically these are written to a http.Request.
*/
type GetCustomFieldParams struct {

	/* CustomFieldPublicID.

	   The unique ID of the CustomField.

	   Format: uuid
	*/
	CustomFieldPublicID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get custom field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomFieldParams) WithDefaults() *GetCustomFieldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get custom field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomFieldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get custom field params
func (o *GetCustomFieldParams) WithTimeout(timeout time.Duration) *GetCustomFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom field params
func (o *GetCustomFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom field params
func (o *GetCustomFieldParams) WithContext(ctx context.Context) *GetCustomFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom field params
func (o *GetCustomFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom field params
func (o *GetCustomFieldParams) WithHTTPClient(client *http.Client) *GetCustomFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom field params
func (o *GetCustomFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomFieldPublicID adds the customFieldPublicID to the get custom field params
func (o *GetCustomFieldParams) WithCustomFieldPublicID(customFieldPublicID strfmt.UUID) *GetCustomFieldParams {
	o.SetCustomFieldPublicID(customFieldPublicID)
	return o
}

// SetCustomFieldPublicID adds the customFieldPublicId to the get custom field params
func (o *GetCustomFieldParams) SetCustomFieldPublicID(customFieldPublicID strfmt.UUID) {
	o.CustomFieldPublicID = customFieldPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param custom-field-public-id
	if err := r.SetPathParam("custom-field-public-id", o.CustomFieldPublicID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}