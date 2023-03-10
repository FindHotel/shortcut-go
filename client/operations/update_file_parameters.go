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

// NewUpdateFileParams creates a new UpdateFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFileParams() *UpdateFileParams {
	return &UpdateFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFileParamsWithTimeout creates a new UpdateFileParams object
// with the ability to set a timeout on a request.
func NewUpdateFileParamsWithTimeout(timeout time.Duration) *UpdateFileParams {
	return &UpdateFileParams{
		timeout: timeout,
	}
}

// NewUpdateFileParamsWithContext creates a new UpdateFileParams object
// with the ability to set a context for a request.
func NewUpdateFileParamsWithContext(ctx context.Context) *UpdateFileParams {
	return &UpdateFileParams{
		Context: ctx,
	}
}

// NewUpdateFileParamsWithHTTPClient creates a new UpdateFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFileParamsWithHTTPClient(client *http.Client) *UpdateFileParams {
	return &UpdateFileParams{
		HTTPClient: client,
	}
}

/*
UpdateFileParams contains all the parameters to send to the API endpoint

	for the update file operation.

	Typically these are written to a http.Request.
*/
type UpdateFileParams struct {

	// UpdateFile.
	UpdateFile *models.UpdateFile

	/* FilePublicID.

	   The unique ID assigned to the file in Shortcut.

	   Format: int64
	*/
	FilePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFileParams) WithDefaults() *UpdateFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update file params
func (o *UpdateFileParams) WithTimeout(timeout time.Duration) *UpdateFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update file params
func (o *UpdateFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update file params
func (o *UpdateFileParams) WithContext(ctx context.Context) *UpdateFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update file params
func (o *UpdateFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update file params
func (o *UpdateFileParams) WithHTTPClient(client *http.Client) *UpdateFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update file params
func (o *UpdateFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateFile adds the updateFile to the update file params
func (o *UpdateFileParams) WithUpdateFile(updateFile *models.UpdateFile) *UpdateFileParams {
	o.SetUpdateFile(updateFile)
	return o
}

// SetUpdateFile adds the updateFile to the update file params
func (o *UpdateFileParams) SetUpdateFile(updateFile *models.UpdateFile) {
	o.UpdateFile = updateFile
}

// WithFilePublicID adds the filePublicID to the update file params
func (o *UpdateFileParams) WithFilePublicID(filePublicID int64) *UpdateFileParams {
	o.SetFilePublicID(filePublicID)
	return o
}

// SetFilePublicID adds the filePublicId to the update file params
func (o *UpdateFileParams) SetFilePublicID(filePublicID int64) {
	o.FilePublicID = filePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateFile != nil {
		if err := r.SetBodyParam(o.UpdateFile); err != nil {
			return err
		}
	}

	// path param file-public-id
	if err := r.SetPathParam("file-public-id", swag.FormatInt64(o.FilePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
