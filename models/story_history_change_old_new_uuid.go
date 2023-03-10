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

// StoryHistoryChangeOldNewUUID The Team ID for the Story.
//
// swagger:model StoryHistoryChangeOldNewUuid
type StoryHistoryChangeOldNewUUID struct {

	// The new value.
	// Format: uuid
	New strfmt.UUID `json:"new,omitempty"`

	// The old value.
	// Format: uuid
	Old strfmt.UUID `json:"old,omitempty"`
}

// Validate validates this story history change old new Uuid
func (m *StoryHistoryChangeOldNewUUID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNew(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOld(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoryHistoryChangeOldNewUUID) validateNew(formats strfmt.Registry) error {
	if swag.IsZero(m.New) { // not required
		return nil
	}

	if err := validate.FormatOf("new", "body", "uuid", m.New.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StoryHistoryChangeOldNewUUID) validateOld(formats strfmt.Registry) error {
	if swag.IsZero(m.Old) { // not required
		return nil
	}

	if err := validate.FormatOf("old", "body", "uuid", m.Old.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this story history change old new Uuid based on context it is used
func (m *StoryHistoryChangeOldNewUUID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewUUID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryHistoryChangeOldNewUUID) UnmarshalBinary(b []byte) error {
	var res StoryHistoryChangeOldNewUUID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
