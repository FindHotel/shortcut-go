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

// WorkflowState Workflow State is any of the at least 3 columns. Workflow States correspond to one of 3 types: Unstarted, Started, or Done.
//
// swagger:model WorkflowState
type WorkflowState struct {

	// The hex color for this Workflow State.
	// Min Length: 1
	// Pattern: ^#[a-fA-F0-9]{6}$
	Color string `json:"color,omitempty"`

	// The time/date the Workflow State was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The description of what sort of Stories belong in that Workflow state.
	// Required: true
	Description *string `json:"description"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// global id
	// Required: true
	GlobalID *string `json:"global_id"`

	// The unique ID of the Workflow State.
	// Required: true
	ID *int64 `json:"id"`

	// The Workflow State's name.
	// Required: true
	Name *string `json:"name"`

	// The number of Stories currently in that Workflow State.
	// Required: true
	NumStories *int64 `json:"num_stories"`

	// The number of Story Templates associated with that Workflow State.
	// Required: true
	NumStoryTemplates *int64 `json:"num_story_templates"`

	// The position that the Workflow State is in, starting with 0 at the left.
	// Required: true
	Position *int64 `json:"position"`

	// The type of Workflow State (Unstarted, Started, or Finished)
	// Required: true
	Type *string `json:"type"`

	// When the Workflow State was last updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The verb that triggers a move to that Workflow State when making VCS commits.
	// Required: true
	Verb *string `json:"verb"`
}

// Validate validates this workflow state
func (m *WorkflowState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoryTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowState) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MinLength("color", "body", m.Color, 1); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^#[a-fA-F0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateGlobalID(formats strfmt.Registry) error {

	if err := validate.Required("global_id", "body", m.GlobalID); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateNumStories(formats strfmt.Registry) error {

	if err := validate.Required("num_stories", "body", m.NumStories); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateNumStoryTemplates(formats strfmt.Registry) error {

	if err := validate.Required("num_story_templates", "body", m.NumStoryTemplates); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowState) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("verb", "body", m.Verb); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow state based on context it is used
func (m *WorkflowState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowState) UnmarshalBinary(b []byte) error {
	var res WorkflowState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
