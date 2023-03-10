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

// NewDeleteEpicParams creates a new DeleteEpicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEpicParams() *DeleteEpicParams {
	return &DeleteEpicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEpicParamsWithTimeout creates a new DeleteEpicParams object
// with the ability to set a timeout on a request.
func NewDeleteEpicParamsWithTimeout(timeout time.Duration) *DeleteEpicParams {
	return &DeleteEpicParams{
		timeout: timeout,
	}
}

// NewDeleteEpicParamsWithContext creates a new DeleteEpicParams object
// with the ability to set a context for a request.
func NewDeleteEpicParamsWithContext(ctx context.Context) *DeleteEpicParams {
	return &DeleteEpicParams{
		Context: ctx,
	}
}

// NewDeleteEpicParamsWithHTTPClient creates a new DeleteEpicParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEpicParamsWithHTTPClient(client *http.Client) *DeleteEpicParams {
	return &DeleteEpicParams{
		HTTPClient: client,
	}
}

/*
DeleteEpicParams contains all the parameters to send to the API endpoint

	for the delete epic operation.

	Typically these are written to a http.Request.
*/
type DeleteEpicParams struct {

	/* EpicPublicID.

	   The unique ID of the Epic.

	   Format: int64
	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEpicParams) WithDefaults() *DeleteEpicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete epic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEpicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete epic params
func (o *DeleteEpicParams) WithTimeout(timeout time.Duration) *DeleteEpicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete epic params
func (o *DeleteEpicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete epic params
func (o *DeleteEpicParams) WithContext(ctx context.Context) *DeleteEpicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete epic params
func (o *DeleteEpicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete epic params
func (o *DeleteEpicParams) WithHTTPClient(client *http.Client) *DeleteEpicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete epic params
func (o *DeleteEpicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEpicPublicID adds the epicPublicID to the delete epic params
func (o *DeleteEpicParams) WithEpicPublicID(epicPublicID int64) *DeleteEpicParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the delete epic params
func (o *DeleteEpicParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEpicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
