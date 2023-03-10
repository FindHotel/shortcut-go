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

// NewUpdateMilestoneParams creates a new UpdateMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMilestoneParams() *UpdateMilestoneParams {
	return &UpdateMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMilestoneParamsWithTimeout creates a new UpdateMilestoneParams object
// with the ability to set a timeout on a request.
func NewUpdateMilestoneParamsWithTimeout(timeout time.Duration) *UpdateMilestoneParams {
	return &UpdateMilestoneParams{
		timeout: timeout,
	}
}

// NewUpdateMilestoneParamsWithContext creates a new UpdateMilestoneParams object
// with the ability to set a context for a request.
func NewUpdateMilestoneParamsWithContext(ctx context.Context) *UpdateMilestoneParams {
	return &UpdateMilestoneParams{
		Context: ctx,
	}
}

// NewUpdateMilestoneParamsWithHTTPClient creates a new UpdateMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMilestoneParamsWithHTTPClient(client *http.Client) *UpdateMilestoneParams {
	return &UpdateMilestoneParams{
		HTTPClient: client,
	}
}

/*
UpdateMilestoneParams contains all the parameters to send to the API endpoint

	for the update milestone operation.

	Typically these are written to a http.Request.
*/
type UpdateMilestoneParams struct {

	// UpdateMilestone.
	UpdateMilestone *models.UpdateMilestone

	/* MilestonePublicID.

	   The ID of the Milestone.

	   Format: int64
	*/
	MilestonePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMilestoneParams) WithDefaults() *UpdateMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update milestone params
func (o *UpdateMilestoneParams) WithTimeout(timeout time.Duration) *UpdateMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update milestone params
func (o *UpdateMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update milestone params
func (o *UpdateMilestoneParams) WithContext(ctx context.Context) *UpdateMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update milestone params
func (o *UpdateMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update milestone params
func (o *UpdateMilestoneParams) WithHTTPClient(client *http.Client) *UpdateMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update milestone params
func (o *UpdateMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateMilestone adds the updateMilestone to the update milestone params
func (o *UpdateMilestoneParams) WithUpdateMilestone(updateMilestone *models.UpdateMilestone) *UpdateMilestoneParams {
	o.SetUpdateMilestone(updateMilestone)
	return o
}

// SetUpdateMilestone adds the updateMilestone to the update milestone params
func (o *UpdateMilestoneParams) SetUpdateMilestone(updateMilestone *models.UpdateMilestone) {
	o.UpdateMilestone = updateMilestone
}

// WithMilestonePublicID adds the milestonePublicID to the update milestone params
func (o *UpdateMilestoneParams) WithMilestonePublicID(milestonePublicID int64) *UpdateMilestoneParams {
	o.SetMilestonePublicID(milestonePublicID)
	return o
}

// SetMilestonePublicID adds the milestonePublicId to the update milestone params
func (o *UpdateMilestoneParams) SetMilestonePublicID(milestonePublicID int64) {
	o.MilestonePublicID = milestonePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateMilestone != nil {
		if err := r.SetBodyParam(o.UpdateMilestone); err != nil {
			return err
		}
	}

	// path param milestone-public-id
	if err := r.SetPathParam("milestone-public-id", swag.FormatInt64(o.MilestonePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
