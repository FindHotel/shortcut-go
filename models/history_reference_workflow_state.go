// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryReferenceWorkflowState A references to a Story Workflow State.
//
// swagger:model HistoryReferenceWorkflowState
type HistoryReferenceWorkflowState struct {

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID interface{} `json:"id"`

	// The name of the entity referenced.
	// Required: true
	Name *string `json:"name"`

	// Either "unstarted", "started", or "done".
	// Required: true
	// Enum: [started unstarted done]
	Type *string `json:"type"`
}

// Validate validates this history reference workflow state
func (m *HistoryReferenceWorkflowState) Validate(formats strfmt.Registry) error {
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

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryReferenceWorkflowState) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryReferenceWorkflowState) validateID(formats strfmt.Registry) error {

	if m.ID == nil {
		return errors.Required("id", "body", nil)
	}

	return nil
}

func (m *HistoryReferenceWorkflowState) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var historyReferenceWorkflowStateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["started","unstarted","done"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyReferenceWorkflowStateTypeTypePropEnum = append(historyReferenceWorkflowStateTypeTypePropEnum, v)
	}
}

const (

	// HistoryReferenceWorkflowStateTypeStarted captures enum value "started"
	HistoryReferenceWorkflowStateTypeStarted string = "started"

	// HistoryReferenceWorkflowStateTypeUnstarted captures enum value "unstarted"
	HistoryReferenceWorkflowStateTypeUnstarted string = "unstarted"

	// HistoryReferenceWorkflowStateTypeDone captures enum value "done"
	HistoryReferenceWorkflowStateTypeDone string = "done"
)

// prop value enum
func (m *HistoryReferenceWorkflowState) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyReferenceWorkflowStateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryReferenceWorkflowState) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history reference workflow state based on context it is used
func (m *HistoryReferenceWorkflowState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryReferenceWorkflowState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryReferenceWorkflowState) UnmarshalBinary(b []byte) error {
	var res HistoryReferenceWorkflowState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
