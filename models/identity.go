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

// Identity The Identity of the VCS user that authored the Commit.
//
// swagger:model Identity
type Identity struct {

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// This is your login in VCS.
	// Required: true
	Name *string `json:"name"`

	// The service this Identity is for.
	// Required: true
	// Enum: ["slack","github","gitlab","bitbucket"]
	Type *string `json:"type"`
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
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

func (m *Identity) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Identity) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var identityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["slack","github","gitlab","bitbucket"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		identityTypeTypePropEnum = append(identityTypeTypePropEnum, v)
	}
}

const (

	// IdentityTypeSlack captures enum value "slack"
	IdentityTypeSlack string = "slack"

	// IdentityTypeGithub captures enum value "github"
	IdentityTypeGithub string = "github"

	// IdentityTypeGitlab captures enum value "gitlab"
	IdentityTypeGitlab string = "gitlab"

	// IdentityTypeBitbucket captures enum value "bitbucket"
	IdentityTypeBitbucket string = "bitbucket"
)

// prop value enum
func (m *Identity) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, identityTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Identity) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this identity based on context it is used
func (m *Identity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
