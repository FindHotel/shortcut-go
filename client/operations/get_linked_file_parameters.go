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

// NewGetLinkedFileParams creates a new GetLinkedFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLinkedFileParams() *GetLinkedFileParams {
	return &GetLinkedFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLinkedFileParamsWithTimeout creates a new GetLinkedFileParams object
// with the ability to set a timeout on a request.
func NewGetLinkedFileParamsWithTimeout(timeout time.Duration) *GetLinkedFileParams {
	return &GetLinkedFileParams{
		timeout: timeout,
	}
}

// NewGetLinkedFileParamsWithContext creates a new GetLinkedFileParams object
// with the ability to set a context for a request.
func NewGetLinkedFileParamsWithContext(ctx context.Context) *GetLinkedFileParams {
	return &GetLinkedFileParams{
		Context: ctx,
	}
}

// NewGetLinkedFileParamsWithHTTPClient creates a new GetLinkedFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLinkedFileParamsWithHTTPClient(client *http.Client) *GetLinkedFileParams {
	return &GetLinkedFileParams{
		HTTPClient: client,
	}
}

/*
GetLinkedFileParams contains all the parameters to send to the API endpoint

	for the get linked file operation.

	Typically these are written to a http.Request.
*/
type GetLinkedFileParams struct {

	/* LinkedFilePublicID.

	   The unique identifier of the linked file.

	   Format: int64
	*/
	LinkedFilePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get linked file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkedFileParams) WithDefaults() *GetLinkedFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get linked file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkedFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get linked file params
func (o *GetLinkedFileParams) WithTimeout(timeout time.Duration) *GetLinkedFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get linked file params
func (o *GetLinkedFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get linked file params
func (o *GetLinkedFileParams) WithContext(ctx context.Context) *GetLinkedFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get linked file params
func (o *GetLinkedFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get linked file params
func (o *GetLinkedFileParams) WithHTTPClient(client *http.Client) *GetLinkedFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get linked file params
func (o *GetLinkedFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkedFilePublicID adds the linkedFilePublicID to the get linked file params
func (o *GetLinkedFileParams) WithLinkedFilePublicID(linkedFilePublicID int64) *GetLinkedFileParams {
	o.SetLinkedFilePublicID(linkedFilePublicID)
	return o
}

// SetLinkedFilePublicID adds the linkedFilePublicId to the get linked file params
func (o *GetLinkedFileParams) SetLinkedFilePublicID(linkedFilePublicID int64) {
	o.LinkedFilePublicID = linkedFilePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLinkedFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linked-file-public-id
	if err := r.SetPathParam("linked-file-public-id", swag.FormatInt64(o.LinkedFilePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
