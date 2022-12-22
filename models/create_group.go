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

// CreateGroup create group
//
// swagger:model CreateGroup
type CreateGroup struct {

	// The color you wish to use for the Group in the system.
	// Min Length: 1
	// Pattern: ^#[a-fA-F0-9]{6}$
	Color string `json:"color,omitempty"`

	// The color key you wish to use for the Group in the system.
	// Enum: [blue purple midnight-blue orange yellow-green brass gray fuchsia yellow pink sky-blue green red black slate turquoise]
	ColorKey string `json:"color_key,omitempty"`

	// The description of the Group.
	Description string `json:"description,omitempty"`

	// The Icon id for the avatar of this Group.
	// Format: uuid
	DisplayIconID strfmt.UUID `json:"display_icon_id,omitempty"`

	// The Member ids to add to this Group.
	// Unique: true
	MemberIds []strfmt.UUID `json:"member_ids"`

	// The mention name of this Group.
	// Required: true
	// Max Length: 63
	// Min Length: 1
	MentionName *string `json:"mention_name"`

	// The name of this Group.
	// Required: true
	// Max Length: 63
	// Min Length: 1
	Name *string `json:"name"`

	// The Workflow ids to add to the Group.
	WorkflowIds []int64 `json:"workflow_ids"`
}

// Validate validates this create group
func (m *CreateGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColorKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayIconID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGroup) validateColor(formats strfmt.Registry) error {
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

var createGroupTypeColorKeyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blue","purple","midnight-blue","orange","yellow-green","brass","gray","fuchsia","yellow","pink","sky-blue","green","red","black","slate","turquoise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createGroupTypeColorKeyPropEnum = append(createGroupTypeColorKeyPropEnum, v)
	}
}

const (

	// CreateGroupColorKeyBlue captures enum value "blue"
	CreateGroupColorKeyBlue string = "blue"

	// CreateGroupColorKeyPurple captures enum value "purple"
	CreateGroupColorKeyPurple string = "purple"

	// CreateGroupColorKeyMidnightDashBlue captures enum value "midnight-blue"
	CreateGroupColorKeyMidnightDashBlue string = "midnight-blue"

	// CreateGroupColorKeyOrange captures enum value "orange"
	CreateGroupColorKeyOrange string = "orange"

	// CreateGroupColorKeyYellowDashGreen captures enum value "yellow-green"
	CreateGroupColorKeyYellowDashGreen string = "yellow-green"

	// CreateGroupColorKeyBrass captures enum value "brass"
	CreateGroupColorKeyBrass string = "brass"

	// CreateGroupColorKeyGray captures enum value "gray"
	CreateGroupColorKeyGray string = "gray"

	// CreateGroupColorKeyFuchsia captures enum value "fuchsia"
	CreateGroupColorKeyFuchsia string = "fuchsia"

	// CreateGroupColorKeyYellow captures enum value "yellow"
	CreateGroupColorKeyYellow string = "yellow"

	// CreateGroupColorKeyPink captures enum value "pink"
	CreateGroupColorKeyPink string = "pink"

	// CreateGroupColorKeySkyDashBlue captures enum value "sky-blue"
	CreateGroupColorKeySkyDashBlue string = "sky-blue"

	// CreateGroupColorKeyGreen captures enum value "green"
	CreateGroupColorKeyGreen string = "green"

	// CreateGroupColorKeyRed captures enum value "red"
	CreateGroupColorKeyRed string = "red"

	// CreateGroupColorKeyBlack captures enum value "black"
	CreateGroupColorKeyBlack string = "black"

	// CreateGroupColorKeySlate captures enum value "slate"
	CreateGroupColorKeySlate string = "slate"

	// CreateGroupColorKeyTurquoise captures enum value "turquoise"
	CreateGroupColorKeyTurquoise string = "turquoise"
)

// prop value enum
func (m *CreateGroup) validateColorKeyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createGroupTypeColorKeyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateGroup) validateColorKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ColorKey) { // not required
		return nil
	}

	// value enum
	if err := m.validateColorKeyEnum("color_key", "body", m.ColorKey); err != nil {
		return err
	}

	return nil
}

func (m *CreateGroup) validateDisplayIconID(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayIconID) { // not required
		return nil
	}

	if err := validate.FormatOf("display_icon_id", "body", "uuid", m.DisplayIconID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateGroup) validateMemberIds(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("member_ids", "body", m.MemberIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberIds); i++ {

		if err := validate.FormatOf("member_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MemberIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateGroup) validateMentionName(formats strfmt.Registry) error {

	if err := validate.Required("mention_name", "body", m.MentionName); err != nil {
		return err
	}

	if err := validate.MinLength("mention_name", "body", *m.MentionName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("mention_name", "body", *m.MentionName, 63); err != nil {
		return err
	}

	return nil
}

func (m *CreateGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 63); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create group based on context it is used
func (m *CreateGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGroup) UnmarshalBinary(b []byte) error {
	var res CreateGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}