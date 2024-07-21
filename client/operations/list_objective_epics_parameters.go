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

// NewListObjectiveEpicsParams creates a new ListObjectiveEpicsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListObjectiveEpicsParams() *ListObjectiveEpicsParams {
	return &ListObjectiveEpicsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListObjectiveEpicsParamsWithTimeout creates a new ListObjectiveEpicsParams object
// with the ability to set a timeout on a request.
func NewListObjectiveEpicsParamsWithTimeout(timeout time.Duration) *ListObjectiveEpicsParams {
	return &ListObjectiveEpicsParams{
		timeout: timeout,
	}
}

// NewListObjectiveEpicsParamsWithContext creates a new ListObjectiveEpicsParams object
// with the ability to set a context for a request.
func NewListObjectiveEpicsParamsWithContext(ctx context.Context) *ListObjectiveEpicsParams {
	return &ListObjectiveEpicsParams{
		Context: ctx,
	}
}

// NewListObjectiveEpicsParamsWithHTTPClient creates a new ListObjectiveEpicsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListObjectiveEpicsParamsWithHTTPClient(client *http.Client) *ListObjectiveEpicsParams {
	return &ListObjectiveEpicsParams{
		HTTPClient: client,
	}
}

/*
ListObjectiveEpicsParams contains all the parameters to send to the API endpoint

	for the list objective epics operation.

	Typically these are written to a http.Request.
*/
type ListObjectiveEpicsParams struct {

	/* ObjectivePublicID.

	   The ID of the Objective.

	   Format: int64
	*/
	ObjectivePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list objective epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectiveEpicsParams) WithDefaults() *ListObjectiveEpicsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list objective epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectiveEpicsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list objective epics params
func (o *ListObjectiveEpicsParams) WithTimeout(timeout time.Duration) *ListObjectiveEpicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list objective epics params
func (o *ListObjectiveEpicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list objective epics params
func (o *ListObjectiveEpicsParams) WithContext(ctx context.Context) *ListObjectiveEpicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list objective epics params
func (o *ListObjectiveEpicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list objective epics params
func (o *ListObjectiveEpicsParams) WithHTTPClient(client *http.Client) *ListObjectiveEpicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list objective epics params
func (o *ListObjectiveEpicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithObjectivePublicID adds the objectivePublicID to the list objective epics params
func (o *ListObjectiveEpicsParams) WithObjectivePublicID(objectivePublicID int64) *ListObjectiveEpicsParams {
	o.SetObjectivePublicID(objectivePublicID)
	return o
}

// SetObjectivePublicID adds the objectivePublicId to the list objective epics params
func (o *ListObjectiveEpicsParams) SetObjectivePublicID(objectivePublicID int64) {
	o.ObjectivePublicID = objectivePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListObjectiveEpicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
