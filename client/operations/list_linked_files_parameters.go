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

// NewListLinkedFilesParams creates a new ListLinkedFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLinkedFilesParams() *ListLinkedFilesParams {
	return &ListLinkedFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLinkedFilesParamsWithTimeout creates a new ListLinkedFilesParams object
// with the ability to set a timeout on a request.
func NewListLinkedFilesParamsWithTimeout(timeout time.Duration) *ListLinkedFilesParams {
	return &ListLinkedFilesParams{
		timeout: timeout,
	}
}

// NewListLinkedFilesParamsWithContext creates a new ListLinkedFilesParams object
// with the ability to set a context for a request.
func NewListLinkedFilesParamsWithContext(ctx context.Context) *ListLinkedFilesParams {
	return &ListLinkedFilesParams{
		Context: ctx,
	}
}

// NewListLinkedFilesParamsWithHTTPClient creates a new ListLinkedFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLinkedFilesParamsWithHTTPClient(client *http.Client) *ListLinkedFilesParams {
	return &ListLinkedFilesParams{
		HTTPClient: client,
	}
}

/*
ListLinkedFilesParams contains all the parameters to send to the API endpoint

	for the list linked files operation.

	Typically these are written to a http.Request.
*/
type ListLinkedFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list linked files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLinkedFilesParams) WithDefaults() *ListLinkedFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list linked files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLinkedFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list linked files params
func (o *ListLinkedFilesParams) WithTimeout(timeout time.Duration) *ListLinkedFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list linked files params
func (o *ListLinkedFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list linked files params
func (o *ListLinkedFilesParams) WithContext(ctx context.Context) *ListLinkedFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list linked files params
func (o *ListLinkedFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list linked files params
func (o *ListLinkedFilesParams) WithHTTPClient(client *http.Client) *ListLinkedFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list linked files params
func (o *ListLinkedFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListLinkedFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
