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

// NewGetEpicParams creates a new GetEpicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEpicParams() *GetEpicParams {
	return &GetEpicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEpicParamsWithTimeout creates a new GetEpicParams object
// with the ability to set a timeout on a request.
func NewGetEpicParamsWithTimeout(timeout time.Duration) *GetEpicParams {
	return &GetEpicParams{
		timeout: timeout,
	}
}

// NewGetEpicParamsWithContext creates a new GetEpicParams object
// with the ability to set a context for a request.
func NewGetEpicParamsWithContext(ctx context.Context) *GetEpicParams {
	return &GetEpicParams{
		Context: ctx,
	}
}

// NewGetEpicParamsWithHTTPClient creates a new GetEpicParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEpicParamsWithHTTPClient(client *http.Client) *GetEpicParams {
	return &GetEpicParams{
		HTTPClient: client,
	}
}

/*
GetEpicParams contains all the parameters to send to the API endpoint

	for the get epic operation.

	Typically these are written to a http.Request.
*/
type GetEpicParams struct {

	/* EpicPublicID.

	   The unique ID of the Epic.

	   Format: int64
	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEpicParams) WithDefaults() *GetEpicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEpicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get epic params
func (o *GetEpicParams) WithTimeout(timeout time.Duration) *GetEpicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get epic params
func (o *GetEpicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get epic params
func (o *GetEpicParams) WithContext(ctx context.Context) *GetEpicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get epic params
func (o *GetEpicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get epic params
func (o *GetEpicParams) WithHTTPClient(client *http.Client) *GetEpicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get epic params
func (o *GetEpicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEpicPublicID adds the epicPublicID to the get epic params
func (o *GetEpicParams) WithEpicPublicID(epicPublicID int64) *GetEpicParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the get epic params
func (o *GetEpicParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEpicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param epic-public-id
	if err := r.SetPathParam("epic-public-id", swag.FormatInt64(o.EpicPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
