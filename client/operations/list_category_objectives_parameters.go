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

// NewListCategoryObjectivesParams creates a new ListCategoryObjectivesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCategoryObjectivesParams() *ListCategoryObjectivesParams {
	return &ListCategoryObjectivesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCategoryObjectivesParamsWithTimeout creates a new ListCategoryObjectivesParams object
// with the ability to set a timeout on a request.
func NewListCategoryObjectivesParamsWithTimeout(timeout time.Duration) *ListCategoryObjectivesParams {
	return &ListCategoryObjectivesParams{
		timeout: timeout,
	}
}

// NewListCategoryObjectivesParamsWithContext creates a new ListCategoryObjectivesParams object
// with the ability to set a context for a request.
func NewListCategoryObjectivesParamsWithContext(ctx context.Context) *ListCategoryObjectivesParams {
	return &ListCategoryObjectivesParams{
		Context: ctx,
	}
}

// NewListCategoryObjectivesParamsWithHTTPClient creates a new ListCategoryObjectivesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCategoryObjectivesParamsWithHTTPClient(client *http.Client) *ListCategoryObjectivesParams {
	return &ListCategoryObjectivesParams{
		HTTPClient: client,
	}
}

/*
ListCategoryObjectivesParams contains all the parameters to send to the API endpoint

	for the list category objectives operation.

	Typically these are written to a http.Request.
*/
type ListCategoryObjectivesParams struct {

	/* CategoryPublicID.

	   The unique ID of the Category.

	   Format: int64
	*/
	CategoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list category objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoryObjectivesParams) WithDefaults() *ListCategoryObjectivesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list category objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoryObjectivesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list category objectives params
func (o *ListCategoryObjectivesParams) WithTimeout(timeout time.Duration) *ListCategoryObjectivesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list category objectives params
func (o *ListCategoryObjectivesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list category objectives params
func (o *ListCategoryObjectivesParams) WithContext(ctx context.Context) *ListCategoryObjectivesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list category objectives params
func (o *ListCategoryObjectivesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list category objectives params
func (o *ListCategoryObjectivesParams) WithHTTPClient(client *http.Client) *ListCategoryObjectivesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list category objectives params
func (o *ListCategoryObjectivesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryPublicID adds the categoryPublicID to the list category objectives params
func (o *ListCategoryObjectivesParams) WithCategoryPublicID(categoryPublicID int64) *ListCategoryObjectivesParams {
	o.SetCategoryPublicID(categoryPublicID)
	return o
}

// SetCategoryPublicID adds the categoryPublicId to the list category objectives params
func (o *ListCategoryObjectivesParams) SetCategoryPublicID(categoryPublicID int64) {
	o.CategoryPublicID = categoryPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListCategoryObjectivesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category-public-id
	if err := r.SetPathParam("category-public-id", swag.FormatInt64(o.CategoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}