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

// NewListLabelsParams creates a new ListLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLabelsParams() *ListLabelsParams {
	return &ListLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLabelsParamsWithTimeout creates a new ListLabelsParams object
// with the ability to set a timeout on a request.
func NewListLabelsParamsWithTimeout(timeout time.Duration) *ListLabelsParams {
	return &ListLabelsParams{
		timeout: timeout,
	}
}

// NewListLabelsParamsWithContext creates a new ListLabelsParams object
// with the ability to set a context for a request.
func NewListLabelsParamsWithContext(ctx context.Context) *ListLabelsParams {
	return &ListLabelsParams{
		Context: ctx,
	}
}

// NewListLabelsParamsWithHTTPClient creates a new ListLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLabelsParamsWithHTTPClient(client *http.Client) *ListLabelsParams {
	return &ListLabelsParams{
		HTTPClient: client,
	}
}

/*
ListLabelsParams contains all the parameters to send to the API endpoint

	for the list labels operation.

	Typically these are written to a http.Request.
*/
type ListLabelsParams struct {

	// ListLabels.
	ListLabels *models.ListLabels

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelsParams) WithDefaults() *ListLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list labels params
func (o *ListLabelsParams) WithTimeout(timeout time.Duration) *ListLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list labels params
func (o *ListLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list labels params
func (o *ListLabelsParams) WithContext(ctx context.Context) *ListLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list labels params
func (o *ListLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list labels params
func (o *ListLabelsParams) WithHTTPClient(client *http.Client) *ListLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list labels params
func (o *ListLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListLabels adds the listLabels to the list labels params
func (o *ListLabelsParams) WithListLabels(listLabels *models.ListLabels) *ListLabelsParams {
	o.SetListLabels(listLabels)
	return o
}

// SetListLabels adds the listLabels to the list labels params
func (o *ListLabelsParams) SetListLabels(listLabels *models.ListLabels) {
	o.ListLabels = listLabels
}

// WriteToRequest writes these params to a swagger request
func (o *ListLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ListLabels != nil {
		if err := r.SetBodyParam(o.ListLabels); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
