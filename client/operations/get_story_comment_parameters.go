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

// NewGetStoryCommentParams creates a new GetStoryCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStoryCommentParams() *GetStoryCommentParams {
	return &GetStoryCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoryCommentParamsWithTimeout creates a new GetStoryCommentParams object
// with the ability to set a timeout on a request.
func NewGetStoryCommentParamsWithTimeout(timeout time.Duration) *GetStoryCommentParams {
	return &GetStoryCommentParams{
		timeout: timeout,
	}
}

// NewGetStoryCommentParamsWithContext creates a new GetStoryCommentParams object
// with the ability to set a context for a request.
func NewGetStoryCommentParamsWithContext(ctx context.Context) *GetStoryCommentParams {
	return &GetStoryCommentParams{
		Context: ctx,
	}
}

// NewGetStoryCommentParamsWithHTTPClient creates a new GetStoryCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStoryCommentParamsWithHTTPClient(client *http.Client) *GetStoryCommentParams {
	return &GetStoryCommentParams{
		HTTPClient: client,
	}
}

/*
GetStoryCommentParams contains all the parameters to send to the API endpoint

	for the get story comment operation.

	Typically these are written to a http.Request.
*/
type GetStoryCommentParams struct {

	/* CommentPublicID.

	   The ID of the Comment.

	   Format: int64
	*/
	CommentPublicID int64

	/* StoryPublicID.

	   The ID of the Story that the Comment is in.

	   Format: int64
	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoryCommentParams) WithDefaults() *GetStoryCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get story comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoryCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get story comment params
func (o *GetStoryCommentParams) WithTimeout(timeout time.Duration) *GetStoryCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get story comment params
func (o *GetStoryCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get story comment params
func (o *GetStoryCommentParams) WithContext(ctx context.Context) *GetStoryCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get story comment params
func (o *GetStoryCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get story comment params
func (o *GetStoryCommentParams) WithHTTPClient(client *http.Client) *GetStoryCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get story comment params
func (o *GetStoryCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentPublicID adds the commentPublicID to the get story comment params
func (o *GetStoryCommentParams) WithCommentPublicID(commentPublicID int64) *GetStoryCommentParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the get story comment params
func (o *GetStoryCommentParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithStoryPublicID adds the storyPublicID to the get story comment params
func (o *GetStoryCommentParams) WithStoryPublicID(storyPublicID int64) *GetStoryCommentParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the get story comment params
func (o *GetStoryCommentParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoryCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment-public-id
	if err := r.SetPathParam("comment-public-id", swag.FormatInt64(o.CommentPublicID)); err != nil {
		return err
	}

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
