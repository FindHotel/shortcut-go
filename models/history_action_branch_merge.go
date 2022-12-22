// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryActionBranchMerge An action representing a VCS Branch being merged.
//
// swagger:model HistoryActionBranchMerge
type HistoryActionBranchMerge struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: [merge]
	Action *string `json:"action"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// The name of the VCS Branch that was pushed
	// Required: true
	Name *string `json:"name"`

	// The URL from the provider of the VCS Branch that was pushed
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this history action branch merge
func (m *HistoryActionBranchMerge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historyActionBranchMergeTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["merge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionBranchMergeTypeActionPropEnum = append(historyActionBranchMergeTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionBranchMergeActionMerge captures enum value "merge"
	HistoryActionBranchMergeActionMerge string = "merge"
)

// prop value enum
func (m *HistoryActionBranchMerge) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionBranchMergeTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionBranchMerge) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionBranchMerge) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionBranchMerge) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionBranchMerge) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionBranchMerge) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history action branch merge based on context it is used
func (m *HistoryActionBranchMerge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionBranchMerge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionBranchMerge) UnmarshalBinary(b []byte) error {
	var res HistoryActionBranchMerge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}