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

// HistoryActionTaskCreate An action representing a Task being created.
//
// swagger:model HistoryActionTaskCreate
type HistoryActionTaskCreate struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: [create]
	Action *string `json:"action"`

	// Whether or not the Task is complete.
	// Required: true
	Complete *bool `json:"complete"`

	// A timestamp that represent's the Task's deadline.
	Deadline string `json:"deadline,omitempty"`

	// The description of the Task.
	// Required: true
	Description *string `json:"description"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// An array of Groups IDs that represent which have been mentioned in the Task.
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// An array of Member IDs that represent who has been mentioned in the Task.
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// An array of Member IDs that represent the Task's owners.
	OwnerIds []strfmt.UUID `json:"owner_ids"`
}

// Validate validates this history action task create
func (m *HistoryActionTaskCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionIds(formats); err != nil {
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

var historyActionTaskCreateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionTaskCreateTypeActionPropEnum = append(historyActionTaskCreateTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionTaskCreateActionCreate captures enum value "create"
	HistoryActionTaskCreateActionCreate string = "create"
)

// prop value enum
func (m *HistoryActionTaskCreate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionTaskCreateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionTaskCreate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskCreate) validateComplete(formats strfmt.Registry) error {

	if err := validate.Required("complete", "body", m.Complete); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskCreate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskCreate) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskCreate) validateGroupMentionIds(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupMentionIds) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupMentionIds); i++ {

		if err := validate.FormatOf("group_mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupMentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *HistoryActionTaskCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionTaskCreate) validateMentionIds(formats strfmt.Registry) error {
	if swag.IsZero(m.MentionIds) { // not required
		return nil
	}

	for i := 0; i < len(m.MentionIds); i++ {

		if err := validate.FormatOf("mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *HistoryActionTaskCreate) validateOwnerIds(formats strfmt.Registry) error {
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

// ContextValidate validates this history action task create based on context it is used
func (m *HistoryActionTaskCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionTaskCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionTaskCreate) UnmarshalBinary(b []byte) error {
	var res HistoryActionTaskCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
