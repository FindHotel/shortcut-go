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
	"github.com/go-openapi/swag"
)

// NewDeleteObjectiveParams creates a new DeleteObjectiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteObjectiveParams() *DeleteObjectiveParams {
	return &DeleteObjectiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteObjectiveParamsWithTimeout creates a new DeleteObjectiveParams object
// with the ability to set a timeout on a request.
func NewDeleteObjectiveParamsWithTimeout(timeout time.Duration) *DeleteObjectiveParams {
	return &DeleteObjectiveParams{
		timeout: timeout,
	}
}

// NewDeleteObjectiveParamsWithContext creates a new DeleteObjectiveParams object
// with the ability to set a context for a request.
func NewDeleteObjectiveParamsWithContext(ctx context.Context) *DeleteObjectiveParams {
	return &DeleteObjectiveParams{
		Context: ctx,
	}
}

// NewDeleteObjectiveParamsWithHTTPClient creates a new DeleteObjectiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteObjectiveParamsWithHTTPClient(client *http.Client) *DeleteObjectiveParams {
	return &DeleteObjectiveParams{
		HTTPClient: client,
	}
}

/*
DeleteObjectiveParams contains all the parameters to send to the API endpoint

	for the delete objective operation.

	Typically these are written to a http.Request.
*/
type DeleteObjectiveParams struct {

	/* ObjectivePublicID.

	   The ID of the Objective.

	   Format: int64
	*/
	ObjectivePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete objective params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectiveParams) WithDefaults() *DeleteObjectiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete objective params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete objective params
func (o *DeleteObjectiveParams) WithTimeout(timeout time.Duration) *DeleteObjectiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete objective params
func (o *DeleteObjectiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete objective params
func (o *DeleteObjectiveParams) WithContext(ctx context.Context) *DeleteObjectiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete objective params
func (o *DeleteObjectiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete objective params
func (o *DeleteObjectiveParams) WithHTTPClient(client *http.Client) *DeleteObjectiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete objective params
func (o *DeleteObjectiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithObjectivePublicID adds the objectivePublicID to the delete objective params
func (o *DeleteObjectiveParams) WithObjectivePublicID(objectivePublicID int64) *DeleteObjectiveParams {
	o.SetObjectivePublicID(objectivePublicID)
	return o
}

// SetObjectivePublicID adds the objectivePublicId to the delete objective params
func (o *DeleteObjectiveParams) SetObjectivePublicID(objectivePublicID int64) {
	o.ObjectivePublicID = objectivePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteObjectiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param objective-public-id
	if err := r.SetPathParam("objective-public-id", swag.FormatInt64(o.ObjectivePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
