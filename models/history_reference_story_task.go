// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryReferenceStoryTask A reference to a Story Task.
//
// swagger:model HistoryReferenceStoryTask
type HistoryReferenceStoryTask struct {

	// The description of the Story Task.
	// Required: true
	Description *string `json:"description"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID interface{} `json:"id"`
}

// Validate validates this history reference story task
func (m *HistoryReferenceStoryTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceStoryTask) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceStoryTask) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceStoryTask) validateID(formats strfmt.Registry) error {

	if m.ID == nil {
		return errors.Required("id", "body", nil)
	}

	return nil
}

// ContextValidate validates this history reference story task based on context it is used
func (m *HistoryReferenceStoryTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceStoryTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceStoryTask) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceStoryTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}