// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BaseTaskParams Request parameters for specifying how to pre-populate a task through a template.
//
// swagger:model BaseTaskParams
type BaseTaskParams struct {

	// True/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete,omitempty"`

	// The Task description.
	// Required: true
	// Max Length: 2048
	// Min Length: 1
	Description *string `json:"description"`

	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`

	// An array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIds []strfmt.UUID `json:"owner_ids"`
}

// Validate validates this base task params
func (m *BaseTaskParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseTaskParams) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 2048); err != nil {
		return err
	}

	return nil
}

func (m *BaseTaskParams) validateOwnerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this base task params based on context it is used
func (m *BaseTaskParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseTaskParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseTaskParams) UnmarshalBinary(b []byte) error {
	var res BaseTaskParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}