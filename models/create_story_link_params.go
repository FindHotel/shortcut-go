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

// CreateStoryLinkParams Request parameters for creating a Story Link within a Story.
//
// swagger:model CreateStoryLinkParams
type CreateStoryLinkParams struct {

	// The unique ID of the Story defined as object.
	ObjectID int64 `json:"object_id,omitempty"`

	// The unique ID of the Story defined as subject.
	SubjectID int64 `json:"subject_id,omitempty"`

	// How the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	// Required: true
	// Enum: ["blocks","duplicates","relates to"]
	Verb *string `json:"verb"`
}

// Validate validates this create story link params
func (m *CreateStoryLinkParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createStoryLinkParamsTypeVerbPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocks","duplicates","relates to"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createStoryLinkParamsTypeVerbPropEnum = append(createStoryLinkParamsTypeVerbPropEnum, v)
	}
}

const (

	// CreateStoryLinkParamsVerbBlocks captures enum value "blocks"
	CreateStoryLinkParamsVerbBlocks string = "blocks"

	// CreateStoryLinkParamsVerbDuplicates captures enum value "duplicates"
	CreateStoryLinkParamsVerbDuplicates string = "duplicates"

	// CreateStoryLinkParamsVerbRelatesTo captures enum value "relates to"
	CreateStoryLinkParamsVerbRelatesTo string = "relates to"
)

// prop value enum
func (m *CreateStoryLinkParams) validateVerbEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createStoryLinkParamsTypeVerbPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateStoryLinkParams) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("verb", "body", m.Verb); err != nil {
		return err
	}

	// value enum
	if err := m.validateVerbEnum("verb", "body", *m.Verb); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create story link params based on context it is used
func (m *CreateStoryLinkParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateStoryLinkParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStoryLinkParams) UnmarshalBinary(b []byte) error {
	var res CreateStoryLinkParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
