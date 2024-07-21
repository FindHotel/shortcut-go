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

// HistoryReferenceCustomFieldEnumValue A reference to a CustomField value asserted on a Story.
//
// swagger:model HistoryReferenceCustomFieldEnumValue
type HistoryReferenceCustomFieldEnumValue struct {

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// Whether or not the custom-field enum value is enabled.
	// Required: true
	EnumValueEnabled *bool `json:"enum_value_enabled"`

	// Whether or not the custom-field is enabled.
	// Required: true
	FieldEnabled *bool `json:"field_enabled"`

	// The public-id of the parent custom-field of this enum value.
	// Required: true
	// Format: uuid
	FieldID *strfmt.UUID `json:"field_id"`

	// The name as it is displayed to the user of the parent custom-field of this enum value.
	// Required: true
	FieldName *string `json:"field_name"`

	// The type variety of the parent custom-field of this enum value.
	// Required: true
	FieldType *string `json:"field_type"`

	// The ID of the entity referenced.
	// Required: true
	ID interface{} `json:"id"`

	// The custom-field enum value as a string.
	// Required: true
	StringValue *string `json:"string_value"`
}

// Validate validates this history reference custom field enum value
func (m *HistoryReferenceCustomFieldEnumValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnumValueEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStringValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateEnumValueEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enum_value_enabled", "body", m.EnumValueEnabled); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateFieldEnabled(formats strfmt.Registry) error {

	if err := validate.Required("field_enabled", "body", m.FieldEnabled); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateFieldID(formats strfmt.Registry) error {

	if err := validate.Required("field_id", "body", m.FieldID); err != nil {
		return err
	}

	if err := validate.FormatOf("field_id", "body", "uuid", m.FieldID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateFieldName(formats strfmt.Registry) error {

	if err := validate.Required("field_name", "body", m.FieldName); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateFieldType(formats strfmt.Registry) error {

	if err := validate.Required("field_type", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateID(formats strfmt.Registry) error {

	if m.ID == nil {
		return errors.Required("id", "body", nil)
	}

	return nil
}

func (m *HistoryReferenceCustomFieldEnumValue) validateStringValue(formats strfmt.Registry) error {

	if err := validate.Required("string_value", "body", m.StringValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history reference custom field enum value based on context it is used
func (m *HistoryReferenceCustomFieldEnumValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceCustomFieldEnumValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceCustomFieldEnumValue) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceCustomFieldEnumValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
