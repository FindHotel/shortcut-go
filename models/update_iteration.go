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

// UpdateIteration update iteration
//
// swagger:model UpdateIteration
type UpdateIteration struct {

	// The description of the Iteration.
	// Max Length: 100000
	Description string `json:"description,omitempty"`

	// The date this Iteration ends, e.g. 2019-07-05.
	// Min Length: 1
	EndDate string `json:"end_date,omitempty"`

	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []strfmt.UUID `json:"group_ids"`

	// An array of Labels attached to the Iteration.
	Labels []*CreateLabelParams `json:"labels"`

	// The name of this Iteration
	// Max Length: 256
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// The date this Iteration begins, e.g. 2019-07-01
	// Min Length: 1
	StartDate string `json:"start_date,omitempty"`
}

// Validate validates this update iteration
func (m *UpdateIteration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateIteration) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 100000); err != nil {
		return err
	}

	return nil
}

func (m *UpdateIteration) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.MinLength("end_date", "body", m.EndDate, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateIteration) validateFollowerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.FollowerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateIteration) validateGroupIds(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupIds) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupIds); i++ {

		if err := validate.FormatOf("group_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateIteration) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateIteration) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *UpdateIteration) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.MinLength("start_date", "body", m.StartDate, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update iteration based on the context it is used
func (m *UpdateIteration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateIteration) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {

			if swag.IsZero(m.Labels[i]) { // not required
				return nil
			}

			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateIteration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateIteration) UnmarshalBinary(b []byte) error {
	var res UpdateIteration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
