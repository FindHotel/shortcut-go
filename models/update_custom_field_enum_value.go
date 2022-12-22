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

// UpdateCustomFieldEnumValue update custom field enum value
//
// swagger:model UpdateCustomFieldEnumValue
type UpdateCustomFieldEnumValue struct {

	// A color key associated with this EnumValue within the CustomField's domain.
	ColorKey *string `json:"color_key,omitempty"`

	// Whether this EnumValue is enabled for its CustomField or not. Leaving this key out of the request leaves the current enabled state untouched.
	Enabled bool `json:"enabled,omitempty"`

	// The unique ID of an existing EnumValue within the CustomField's domain.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// A string value within the domain of this Custom Field.
	// Max Length: 63
	// Min Length: 1
	Value string `json:"value,omitempty"`
}

// Validate validates this update custom field enum value
func (m *UpdateCustomFieldEnumValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCustomFieldEnumValue) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCustomFieldEnumValue) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.MinLength("value", "body", m.Value, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", m.Value, 63); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update custom field enum value based on context it is used
func (m *UpdateCustomFieldEnumValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCustomFieldEnumValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCustomFieldEnumValue) UnmarshalBinary(b []byte) error {
	var res UpdateCustomFieldEnumValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}