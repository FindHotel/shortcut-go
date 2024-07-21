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

// ObjectiveSearchResult A Milestone in search results. This is typed differently from Milestone because the details=slim search argument will omit some fields.
//
// swagger:model ObjectiveSearchResult
type ObjectiveSearchResult struct {

	// The Shortcut application url for the Milestone.
	// Required: true
	AppURL *string `json:"app_url"`

	// A boolean indicating whether the Milestone has been archived or not.
	// Required: true
	Archived *bool `json:"archived"`

	// An array of Categories attached to the Milestone.
	// Required: true
	Categories []*Category `json:"categories"`

	// A true/false boolean indicating if the Milestone has been completed.
	// Required: true
	Completed *bool `json:"completed"`

	// The time/date the Milestone was completed.
	// Required: true
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at"`

	// A manual override for the time/date the Milestone was completed.
	// Required: true
	// Format: date-time
	CompletedAtOverride *strfmt.DateTime `json:"completed_at_override"`

	// The time/date the Milestone was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The Milestone's description.
	Description string `json:"description,omitempty"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// global id
	// Required: true
	GlobalID *string `json:"global_id"`

	// The unique ID of the Milestone.
	// Required: true
	ID *int64 `json:"id"`

	// The IDs of the Key Results associated with the Objective.
	// Required: true
	KeyResultIds []strfmt.UUID `json:"key_result_ids"`

	// The name of the Milestone.
	// Required: true
	Name *string `json:"name"`

	// A number representing the position of the Milestone in relation to every other Milestone within the Workspace.
	// Required: true
	Position *int64 `json:"position"`

	// A true/false boolean indicating if the Milestone has been started.
	// Required: true
	Started *bool `json:"started"`

	// The time/date the Milestone was started.
	// Required: true
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at"`

	// A manual override for the time/date the Milestone was started.
	// Required: true
	// Format: date-time
	StartedAtOverride *strfmt.DateTime `json:"started_at_override"`

	// The workflow state that the Milestone is in.
	// Required: true
	State *string `json:"state"`

	// stats
	// Required: true
	Stats *MilestoneStats `json:"stats"`

	// The time/date the Milestone was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this objective search result
func (m *ObjectiveSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
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

	if err := m.validateKeyResultIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectiveSearchResult) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateCategories(formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
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

func (m *ObjectiveSearchResult) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateCompletedAt(formats strfmt.Registry) error {

	if err := validate.Required("completed_at", "body", m.CompletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateCompletedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("completed_at_override", "body", m.CompletedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateGlobalID(formats strfmt.Registry) error {

	if err := validate.Required("global_id", "body", m.GlobalID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateKeyResultIds(formats strfmt.Registry) error {

	if err := validate.Required("key_result_ids", "body", m.KeyResultIds); err != nil {
		return err
	}

	for i := 0; i < len(m.KeyResultIds); i++ {

		if err := validate.FormatOf("key_result_ids"+"."+strconv.Itoa(i), "body", "uuid", m.KeyResultIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectiveSearchResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateStarted(formats strfmt.Registry) error {

	if err := validate.Required("started", "body", m.Started); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", m.StartedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateStartedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("started_at_override", "body", m.StartedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveSearchResult) validateStats(formats strfmt.Registry) error {

	if err := validate.Required("stats", "body", m.Stats); err != nil {
		return err
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectiveSearchResult) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this objective search result based on the context it is used
func (m *ObjectiveSearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectiveSearchResult) contextValidateCategories(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ObjectiveSearchResult) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectiveSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectiveSearchResult) UnmarshalBinary(b []byte) error {
	var res ObjectiveSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
