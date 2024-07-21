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

// NewListStoryCommentParams creates a new ListStoryCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListStoryCommentParams() *ListStoryCommentParams {
	return &ListStoryCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListStoryCommentParamsWithTimeout creates a new ListStoryCommentParams object
// with the ability to set a timeout on a request.
func NewListStoryCommentParamsWithTimeout(timeout time.Duration) *ListStoryCommentParams {
	return &ListStoryCommentParams{
		timeout: timeout,
	}
}

// NewListStoryCommentParamsWithContext creates a new ListStoryCommentParams object
// with the ability to set a context for a request.
func NewListStoryCommentParamsWithContext(ctx context.Context) *ListStoryCommentParams {
	return &ListStoryCommentParams{
		Context: ctx,
	}
}

// NewListStoryCommentParamsWithHTTPClient creates a new ListStoryCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewListStoryCommentParamsWithHTTPClient(client *http.Client) *ListStoryCommentParams {
	return &ListStoryCommentParams{
		HTTPClient: client,
	}
}

/*
ListStoryCommentParams contains all the parameters to send to the API endpoint

	for the list story comment operation.

	Typically these are written to a http.Request.
*/
type ListStoryCommentParams struct {

	/* StoryPublicID.

	   The ID of the Story that the Comment is in.

	   Format: int64
	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStoryCommentParams) WithDefaults() *ListStoryCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStoryCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list story comment params
func (o *ListStoryCommentParams) WithTimeout(timeout time.Duration) *ListStoryCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list story comment params
func (o *ListStoryCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list story comment params
func (o *ListStoryCommentParams) WithContext(ctx context.Context) *ListStoryCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list story comment params
func (o *ListStoryCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list story comment params
func (o *ListStoryCommentParams) WithHTTPClient(client *http.Client) *ListStoryCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list story comment params
func (o *ListStoryCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the list story comment params
func (o *ListStoryCommentParams) WithStoryPublicID(storyPublicID int64) *ListStoryCommentParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the list story comment params
func (o *ListStoryCommentParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *ListStoryCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
