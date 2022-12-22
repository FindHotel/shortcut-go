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

// HistoryReferenceGeneral A default reference for entity types that don't have extra fields.
//
// swagger:model HistoryReferenceGeneral
type HistoryReferenceGeneral struct {

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID interface{} `json:"id"`

	// The name of the entity referenced.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this history reference general
func (m *HistoryReferenceGeneral) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceGeneral) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceGeneral) validateID(formats strfmt.Registry) error {

	if m.ID == nil {
		return errors.Required("id", "body", nil)
	}

	return nil
}

func (m *HistoryReferenceGeneral) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history reference general based on context it is used
func (m *HistoryReferenceGeneral) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceGeneral) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
