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

// HistoryReferenceBranch A reference to a VCS Branch.
//
// swagger:model HistoryReferenceBranch
type HistoryReferenceBranch struct {

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID interface{} `json:"id"`

	// The name of the entity referenced.
	// Required: true
	Name *string `json:"name"`

	// The external URL for the Branch.
	// Required: true
	// Max Length: 2048
	// Pattern: ^https?://.+$
	URL *string `json:"url"`
}

// Validate validates this history reference branch
func (m *HistoryReferenceBranch) Validate(formats strfmt.Registry) error {
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

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceBranch) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceBranch) validateID(formats strfmt.Registry) error {

	if m.ID == nil {
		return errors.Required("id", "body", nil)
	}

	return nil
}

func (m *HistoryReferenceBranch) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceBranch) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MaxLength("url", "body", *m.URL, 2048); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", *m.URL, `^https?://.+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history reference branch based on context it is used
func (m *HistoryReferenceBranch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceBranch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceBranch) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceBranch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}