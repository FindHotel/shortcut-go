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

// NewGetMilestoneParams creates a new GetMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMilestoneParams() *GetMilestoneParams {
	return &GetMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMilestoneParamsWithTimeout creates a new GetMilestoneParams object
// with the ability to set a timeout on a request.
func NewGetMilestoneParamsWithTimeout(timeout time.Duration) *GetMilestoneParams {
	return &GetMilestoneParams{
		timeout: timeout,
	}
}

// NewGetMilestoneParamsWithContext creates a new GetMilestoneParams object
// with the ability to set a context for a request.
func NewGetMilestoneParamsWithContext(ctx context.Context) *GetMilestoneParams {
	return &GetMilestoneParams{
		Context: ctx,
	}
}

// NewGetMilestoneParamsWithHTTPClient creates a new GetMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMilestoneParamsWithHTTPClient(client *http.Client) *GetMilestoneParams {
	return &GetMilestoneParams{
		HTTPClient: client,
	}
}

/*
GetMilestoneParams contains all the parameters to send to the API endpoint

	for the get milestone operation.

	Typically these are written to a http.Request.
*/
type GetMilestoneParams struct {

	/* MilestonePublicID.

	   The ID of the Milestone.

	   Format: int64
	*/
	MilestonePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMilestoneParams) WithDefaults() *GetMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get milestone params
func (o *GetMilestoneParams) WithTimeout(timeout time.Duration) *GetMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get milestone params
func (o *GetMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get milestone params
func (o *GetMilestoneParams) WithContext(ctx context.Context) *GetMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get milestone params
func (o *GetMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get milestone params
func (o *GetMilestoneParams) WithHTTPClient(client *http.Client) *GetMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get milestone params
func (o *GetMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestonePublicID adds the milestonePublicID to the get milestone params
func (o *GetMilestoneParams) WithMilestonePublicID(milestonePublicID int64) *GetMilestoneParams {
	o.SetMilestonePublicID(milestonePublicID)
	return o
}

// SetMilestonePublicID adds the milestonePublicId to the get milestone params
func (o *GetMilestoneParams) SetMilestonePublicID(milestonePublicID int64) {
	o.MilestonePublicID = milestonePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestone-public-id
	if err := r.SetPathParam("milestone-public-id", swag.FormatInt64(o.MilestonePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
