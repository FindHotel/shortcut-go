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

// CreateObjective create objective
//
// swagger:model CreateObjective
type CreateObjective struct {

	// An array of IDs of Categories attached to the Objective.
	Categories []*CreateCategoryParams `json:"categories"`

	// A manual override for the time/date the Objective was completed.
	// Format: date-time
	CompletedAtOverride strfmt.DateTime `json:"completed_at_override,omitempty"`

	// The Objective's description.
	// Max Length: 100000
	Description string `json:"description,omitempty"`

	// The name of the Objective.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name"`

	// A manual override for the time/date the Objective was started.
	// Format: date-time
	StartedAtOverride strfmt.DateTime `json:"started_at_override,omitempty"`

	// The workflow state that the Objective is in.
	// Enum: ["in progress","to do","done"]
	State string `json:"state,omitempty"`
}

// Validate validates this create objective
func (m *CreateObjective) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateObjective) validateCategories(formats strfmt.Registry) error {
	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateObjective) validateCompletedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateObjective) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 100000); err != nil {
		return err
	}

	return nil
}

func (m *CreateObjective) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *CreateObjective) validateStartedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

var createObjectiveTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in progress","to do","done"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createObjectiveTypeStatePropEnum = append(createObjectiveTypeStatePropEnum, v)
	}
}

const (

	// CreateObjectiveStateInProgress captures enum value "in progress"
	CreateObjectiveStateInProgress string = "in progress"

	// CreateObjectiveStateToDo captures enum value "to do"
	CreateObjectiveStateToDo string = "to do"

	// CreateObjectiveStateDone captures enum value "done"
	CreateObjectiveStateDone string = "done"
)

// prop value enum
func (m *CreateObjective) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createObjectiveTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateObjective) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create objective based on the context it is used
func (m *CreateObjective) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateObjective) contextValidateCategories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Categories); i++ {

		if m.Categories[i] != nil {

			if swag.IsZero(m.Categories[i]) { // not required
				return nil
			}

			if err := m.Categories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateObjective) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateObjective) UnmarshalBinary(b []byte) error {
	var res CreateObjective
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
