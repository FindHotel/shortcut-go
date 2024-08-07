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

// NewSearchStoriesParams creates a new SearchStoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchStoriesParams() *SearchStoriesParams {
	return &SearchStoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchStoriesParamsWithTimeout creates a new SearchStoriesParams object
// with the ability to set a timeout on a request.
func NewSearchStoriesParamsWithTimeout(timeout time.Duration) *SearchStoriesParams {
	return &SearchStoriesParams{
		timeout: timeout,
	}
}

// NewSearchStoriesParamsWithContext creates a new SearchStoriesParams object
// with the ability to set a context for a request.
func NewSearchStoriesParamsWithContext(ctx context.Context) *SearchStoriesParams {
	return &SearchStoriesParams{
		Context: ctx,
	}
}

// NewSearchStoriesParamsWithHTTPClient creates a new SearchStoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchStoriesParamsWithHTTPClient(client *http.Client) *SearchStoriesParams {
	return &SearchStoriesParams{
		HTTPClient: client,
	}
}

/*
SearchStoriesParams contains all the parameters to send to the API endpoint

	for the search stories operation.

	Typically these are written to a http.Request.
*/
type SearchStoriesParams struct {

	/* Detail.

	     The amount of detail included in each result item.
	   "full" will include all descriptions and comments and more fields on
	   related items such as pull requests, branches and tasks.
	   "slim" omits larger fulltext fields such as descriptions and comments
	   and only references related items by id.
	   The default is "full".
	*/
	Detail *string

	/* EntityTypes.

	   A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story.
	*/
	EntityTypes []string

	/* Next.

	   The next page token.
	*/
	Next *string

	/* PageSize.

	   The number of search results to include in a page. Minimum of 1 and maximum of 25.

	   Format: int64
	*/
	PageSize *int64

	/* Query.

	   See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators)
	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchStoriesParams) WithDefaults() *SearchStoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search stories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchStoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search stories params
func (o *SearchStoriesParams) WithTimeout(timeout time.Duration) *SearchStoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search stories params
func (o *SearchStoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search stories params
func (o *SearchStoriesParams) WithContext(ctx context.Context) *SearchStoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search stories params
func (o *SearchStoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search stories params
func (o *SearchStoriesParams) WithHTTPClient(client *http.Client) *SearchStoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search stories params
func (o *SearchStoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetail adds the detail to the search stories params
func (o *SearchStoriesParams) WithDetail(detail *string) *SearchStoriesParams {
	o.SetDetail(detail)
	return o
}

// SetDetail adds the detail to the search stories params
func (o *SearchStoriesParams) SetDetail(detail *string) {
	o.Detail = detail
}

// WithEntityTypes adds the entityTypes to the search stories params
func (o *SearchStoriesParams) WithEntityTypes(entityTypes []string) *SearchStoriesParams {
	o.SetEntityTypes(entityTypes)
	return o
}

// SetEntityTypes adds the entityTypes to the search stories params
func (o *SearchStoriesParams) SetEntityTypes(entityTypes []string) {
	o.EntityTypes = entityTypes
}

// WithNext adds the next to the search stories params
func (o *SearchStoriesParams) WithNext(next *string) *SearchStoriesParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the search stories params
func (o *SearchStoriesParams) SetNext(next *string) {
	o.Next = next
}

// WithPageSize adds the pageSize to the search stories params
func (o *SearchStoriesParams) WithPageSize(pageSize *int64) *SearchStoriesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the search stories params
func (o *SearchStoriesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQuery adds the query to the search stories params
func (o *SearchStoriesParams) WithQuery(query string) *SearchStoriesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search stories params
func (o *SearchStoriesParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SearchStoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Detail != nil {

		// query param detail
		var qrDetail string

		if o.Detail != nil {
			qrDetail = *o.Detail
		}
		qDetail := qrDetail
		if qDetail != "" {

			if err := r.SetQueryParam("detail", qDetail); err != nil {
				return err
			}
		}
	}

	if o.EntityTypes != nil {

		// binding items for entity_types
		joinedEntityTypes := o.bindParamEntityTypes(reg)

		// query array param entity_types
		if err := r.SetQueryParam("entity_types", joinedEntityTypes...); err != nil {
			return err
		}
	}

	if o.Next != nil {

		// query param next
		var qrNext string

		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {

			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {

		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchStories binds the parameter entity_types
func (o *SearchStoriesParams) bindParamEntityTypes(formats strfmt.Registry) []string {
	entityTypesIR := o.EntityTypes

	var entityTypesIC []string
	for _, entityTypesIIR := range entityTypesIR { // explode []string

		entityTypesIIV := entityTypesIIR // string as string
		entityTypesIC = append(entityTypesIC, entityTypesIIV)
	}

	// items.CollectionFormat: "multi"
	entityTypesIS := swag.JoinByFormat(entityTypesIC, "multi")

	return entityTypesIS
}
