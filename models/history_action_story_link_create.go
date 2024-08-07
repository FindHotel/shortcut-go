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

// HistoryActionStoryLinkCreate An action representing a Story Link being created.
//
// swagger:model HistoryActionStoryLinkCreate
type HistoryActionStoryLinkCreate struct {

	// The action of the entity referenced.
	// Required: true
	// Enum: ["create"]
	Action *string `json:"action"`

	// The type of entity referenced.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the entity referenced.
	// Required: true
	ID *int64 `json:"id"`

	// The Story ID of the object Story.
	// Required: true
	ObjectID *int64 `json:"object_id"`

	// The Story ID of the subject Story.
	// Required: true
	SubjectID *int64 `json:"subject_id"`

	// The verb describing the link's relationship.
	// Required: true
	// Enum: ["blocks","duplicates","relates to"]
	Verb *string `json:"verb"`
}

// Validate validates this history action story link create
func (m *HistoryActionStoryLinkCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectID(formats); err != nil {
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

var historyActionStoryLinkCreateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionStoryLinkCreateTypeActionPropEnum = append(historyActionStoryLinkCreateTypeActionPropEnum, v)
	}
}

const (

	// HistoryActionStoryLinkCreateActionCreate captures enum value "create"
	HistoryActionStoryLinkCreateActionCreate string = "create"
)

// prop value enum
func (m *HistoryActionStoryLinkCreate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionStoryLinkCreateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionStoryLinkCreate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryLinkCreate) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryLinkCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryLinkCreate) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryActionStoryLinkCreate) validateSubjectID(formats strfmt.Registry) error {

	if err := validate.Required("subject_id", "body", m.SubjectID); err != nil {
		return err
	}

	return nil
}

var historyActionStoryLinkCreateTypeVerbPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocks","duplicates","relates to"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyActionStoryLinkCreateTypeVerbPropEnum = append(historyActionStoryLinkCreateTypeVerbPropEnum, v)
	}
}

const (

	// HistoryActionStoryLinkCreateVerbBlocks captures enum value "blocks"
	HistoryActionStoryLinkCreateVerbBlocks string = "blocks"

	// HistoryActionStoryLinkCreateVerbDuplicates captures enum value "duplicates"
	HistoryActionStoryLinkCreateVerbDuplicates string = "duplicates"

	// HistoryActionStoryLinkCreateVerbRelatesTo captures enum value "relates to"
	HistoryActionStoryLinkCreateVerbRelatesTo string = "relates to"
)

// prop value enum
func (m *HistoryActionStoryLinkCreate) validateVerbEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyActionStoryLinkCreateTypeVerbPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryActionStoryLinkCreate) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("verb", "body", m.Verb); err != nil {
		return err
	}

	// value enum
	if err := m.validateVerbEnum("verb", "body", *m.Verb); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history action story link create based on context it is used
func (m *HistoryActionStoryLinkCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryActionStoryLinkCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryActionStoryLinkCreate) UnmarshalBinary(b []byte) error {
	var res HistoryActionStoryLinkCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
