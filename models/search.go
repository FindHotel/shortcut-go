// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Search search
//
// swagger:model Search
type Search struct {

	// A collection of entity_types to search. Defaults to story and epic. Suports: epic, iteration, story.
	EntityTypes []string `json:"entity_types"`

	// include
	// Enum: [cursors]
	Include string `json:"include,omitempty"`

	// The next page token.
	Next string `json:"next,omitempty"`

	// The number of search results to include in a page. Minimum of 1 and maximum of 25.
	PageSize int64 `json:"page_size,omitempty"`

	// See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators)
	// Required: true
	// Min Length: 1
	Query *string `json:"query"`
}

// Validate validates this search
func (m *Search) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var searchEntityTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["story","milestone","epic","iteration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchEntityTypesItemsEnum = append(searchEntityTypesItemsEnum, v)
	}
}

func (m *Search) validateEntityTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchEntityTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Search) validateEntityTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityTypes); i++ {

		// value enum
		if err := m.validateEntityTypesItemsEnum("entity_types"+"."+strconv.Itoa(i), "body", m.EntityTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var searchTypeIncludePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cursors"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchTypeIncludePropEnum = append(searchTypeIncludePropEnum, v)
	}
}

const (

	// SearchIncludeCursors captures enum value "cursors"
	SearchIncludeCursors string = "cursors"
)

// prop value enum
func (m *Search) validateIncludeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchTypeIncludePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Search) validateInclude(formats strfmt.Registry) error {
	if swag.IsZero(m.Include) { // not required
		return nil
	}

	// value enum
	if err := m.validateIncludeEnum("include", "body", m.Include); err != nil {
		return err
	}

	return nil
}

func (m *Search) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("query", "body", *m.Query, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this search based on context it is used
func (m *Search) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Search) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Search) UnmarshalBinary(b []byte) error {
	var res Search
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
