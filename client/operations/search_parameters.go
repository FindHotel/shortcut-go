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

// NewSearchParams creates a new SearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchParams() *SearchParams {
	return &SearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the ability to set a timeout on a request.
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	return &SearchParams{
		timeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the ability to set a context for a request.
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	return &SearchParams{
		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	return &SearchParams{
		HTTPClient: client,
	}
}

/*
SearchParams contains all the parameters to send to the API endpoint

	for the search operation.

	Typically these are written to a http.Request.
*/
type SearchParams struct {

	// Search.
	Search *models.Search

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) WithDefaults() *SearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout time.Duration) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearch adds the search to the search params
func (o *SearchParams) WithSearch(search *models.Search) *SearchParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the search params
func (o *SearchParams) SetSearch(search *models.Search) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Search != nil {
		if err := r.SetBodyParam(o.Search); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
