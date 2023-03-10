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

// NewListLabelEpicsParams creates a new ListLabelEpicsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLabelEpicsParams() *ListLabelEpicsParams {
	return &ListLabelEpicsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLabelEpicsParamsWithTimeout creates a new ListLabelEpicsParams object
// with the ability to set a timeout on a request.
func NewListLabelEpicsParamsWithTimeout(timeout time.Duration) *ListLabelEpicsParams {
	return &ListLabelEpicsParams{
		timeout: timeout,
	}
}

// NewListLabelEpicsParamsWithContext creates a new ListLabelEpicsParams object
// with the ability to set a context for a request.
func NewListLabelEpicsParamsWithContext(ctx context.Context) *ListLabelEpicsParams {
	return &ListLabelEpicsParams{
		Context: ctx,
	}
}

// NewListLabelEpicsParamsWithHTTPClient creates a new ListLabelEpicsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLabelEpicsParamsWithHTTPClient(client *http.Client) *ListLabelEpicsParams {
	return &ListLabelEpicsParams{
		HTTPClient: client,
	}
}

/*
ListLabelEpicsParams contains all the parameters to send to the API endpoint

	for the list label epics operation.

	Typically these are written to a http.Request.
*/
type ListLabelEpicsParams struct {

	/* LabelPublicID.

	   The unique ID of the Label.

	   Format: int64
	*/
	LabelPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list label epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelEpicsParams) WithDefaults() *ListLabelEpicsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list label epics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelEpicsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list label epics params
func (o *ListLabelEpicsParams) WithTimeout(timeout time.Duration) *ListLabelEpicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list label epics params
func (o *ListLabelEpicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list label epics params
func (o *ListLabelEpicsParams) WithContext(ctx context.Context) *ListLabelEpicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list label epics params
func (o *ListLabelEpicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list label epics params
func (o *ListLabelEpicsParams) WithHTTPClient(client *http.Client) *ListLabelEpicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list label epics params
func (o *ListLabelEpicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabelPublicID adds the labelPublicID to the list label epics params
func (o *ListLabelEpicsParams) WithLabelPublicID(labelPublicID int64) *ListLabelEpicsParams {
	o.SetLabelPublicID(labelPublicID)
	return o
}

// SetLabelPublicID adds the labelPublicId to the list label epics params
func (o *ListLabelEpicsParams) SetLabelPublicID(labelPublicID int64) {
	o.LabelPublicID = labelPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListLabelEpicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param label-public-id
	if err := r.SetPathParam("label-public-id", swag.FormatInt64(o.LabelPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
