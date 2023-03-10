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

// NewDeleteEntityTemplateParams creates a new DeleteEntityTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEntityTemplateParams() *DeleteEntityTemplateParams {
	return &DeleteEntityTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEntityTemplateParamsWithTimeout creates a new DeleteEntityTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteEntityTemplateParamsWithTimeout(timeout time.Duration) *DeleteEntityTemplateParams {
	return &DeleteEntityTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteEntityTemplateParamsWithContext creates a new DeleteEntityTemplateParams object
// with the ability to set a context for a request.
func NewDeleteEntityTemplateParamsWithContext(ctx context.Context) *DeleteEntityTemplateParams {
	return &DeleteEntityTemplateParams{
		Context: ctx,
	}
}

// NewDeleteEntityTemplateParamsWithHTTPClient creates a new DeleteEntityTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEntityTemplateParamsWithHTTPClient(client *http.Client) *DeleteEntityTemplateParams {
	return &DeleteEntityTemplateParams{
		HTTPClient: client,
	}
}

/*
DeleteEntityTemplateParams contains all the parameters to send to the API endpoint

	for the delete entity template operation.

	Typically these are written to a http.Request.
*/
type DeleteEntityTemplateParams struct {

	/* EntityTemplatePublicID.

	   The unique ID of the entity template.

	   Format: uuid
	*/
	EntityTemplatePublicID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete entity template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntityTemplateParams) WithDefaults() *DeleteEntityTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete entity template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntityTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete entity template params
func (o *DeleteEntityTemplateParams) WithTimeout(timeout time.Duration) *DeleteEntityTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete entity template params
func (o *DeleteEntityTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete entity template params
func (o *DeleteEntityTemplateParams) WithContext(ctx context.Context) *DeleteEntityTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete entity template params
func (o *DeleteEntityTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete entity template params
func (o *DeleteEntityTemplateParams) WithHTTPClient(client *http.Client) *DeleteEntityTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete entity template params
func (o *DeleteEntityTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityTemplatePublicID adds the entityTemplatePublicID to the delete entity template params
func (o *DeleteEntityTemplateParams) WithEntityTemplatePublicID(entityTemplatePublicID strfmt.UUID) *DeleteEntityTemplateParams {
	o.SetEntityTemplatePublicID(entityTemplatePublicID)
	return o
}

// SetEntityTemplatePublicID adds the entityTemplatePublicId to the delete entity template params
func (o *DeleteEntityTemplateParams) SetEntityTemplatePublicID(entityTemplatePublicID strfmt.UUID) {
	o.EntityTemplatePublicID = entityTemplatePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEntityTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity-template-public-id
	if err := r.SetPathParam("entity-template-public-id", o.EntityTemplatePublicID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
