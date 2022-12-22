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

	"github.com/FindHotel/shortcut-go/models"
)

// NewListLabelStoriesParams creates a new ListLabelStoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLabelStoriesParams() *ListLabelStoriesParams {
	return &ListLabelStoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLabelStoriesParamsWithTimeout creates a new ListLabelStoriesParams object
// with the ability to set a timeout on a request.
func NewListLabelStoriesParamsWithTimeout(timeout time.Duration) *ListLabelStoriesParams {
	return &ListLabelStoriesParams{
		timeout: timeout,
	}
}

// NewListLabelStoriesParamsWithContext creates a new ListLabelStoriesParams object
// with the ability to set a context for a request.
func NewListLabelStoriesParamsWithContext(ctx context.Context) *ListLabelStoriesParams {
	return &ListLabelStoriesParams{
		Context: ctx,
	}
}

// NewListLabelStoriesParamsWithHTTPClient creates a new ListLabelStoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLabelStoriesParamsWithHTTPClient(client *http.Client) *ListLabelStoriesParams {
	return &ListLabelStoriesParams{
		HTTPClient: client,
	}
}

/*
ListLabelStoriesParams contains all the parameters to send to the API endpoint

	for the list label stories operation.

	Typically these are written to a http.Request.
*/
type ListLabelStoriesParams struct {

	// GetLabelStories.
	GetLabelStories *models.GetLabelStories

	/* LabelPublicID.

	   The unique ID of the Label.

	   Format: int64
	*/
	LabelPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list label stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelStoriesParams) WithDefaults() *ListLabelStoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list label stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelStoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list label stories params
func (o *ListLabelStoriesParams) WithTimeout(timeout time.Duration) *ListLabelStoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list label stories params
func (o *ListLabelStoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list label stories params
func (o *ListLabelStoriesParams) WithContext(ctx context.Context) *ListLabelStoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list label stories params
func (o *ListLabelStoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list label stories params
func (o *ListLabelStoriesParams) WithHTTPClient(client *http.Client) *ListLabelStoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list label stories params
func (o *ListLabelStoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetLabelStories adds the getLabelStories to the list label stories params
func (o *ListLabelStoriesParams) WithGetLabelStories(getLabelStories *models.GetLabelStories) *ListLabelStoriesParams {
	o.SetGetLabelStories(getLabelStories)
	return o
}

// SetGetLabelStories adds the getLabelStories to the list label stories params
func (o *ListLabelStoriesParams) SetGetLabelStories(getLabelStories *models.GetLabelStories) {
	o.GetLabelStories = getLabelStories
}

// WithLabelPublicID adds the labelPublicID to the list label stories params
func (o *ListLabelStoriesParams) WithLabelPublicID(labelPublicID int64) *ListLabelStoriesParams {
	o.SetLabelPublicID(labelPublicID)
	return o
}

// SetLabelPublicID adds the labelPublicId to the list label stories params
func (o *ListLabelStoriesParams) SetLabelPublicID(labelPublicID int64) {
	o.LabelPublicID = labelPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListLabelStoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GetLabelStories != nil {
		if err := r.SetBodyParam(o.GetLabelStories); err != nil {
			return err
		}
	}

	// path param label-public-id
	if err := r.SetPathParam("label-public-id", swag.FormatInt64(o.LabelPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
