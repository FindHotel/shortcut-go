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

// Group A Group.
//
// swagger:model Group
type Group struct {

	// The Shortcut application url for the Group.
	// Required: true
	AppURL *string `json:"app_url"`

	// Whether or not the Group is archived.
	// Required: true
	Archived *bool `json:"archived"`

	// The hex color to be displayed with the Group (for example, "#ff0000").
	// Required: true
	// Min Length: 1
	// Pattern: ^#[a-fA-F0-9]{6}$
	Color *string `json:"color"`

	// The color key to be displayed with the Group.
	// Required: true
	// Enum: [blue purple midnight-blue orange yellow-green brass gray fuchsia yellow pink sky-blue green red black slate turquoise]
	ColorKey *string `json:"color_key"`

	// The description of the Group.
	// Required: true
	Description *string `json:"description"`

	// display icon
	// Required: true
	DisplayIcon *Icon `json:"display_icon"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The id of the Group.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// The Member IDs contain within the Group.
	// Required: true
	MemberIds []strfmt.UUID `json:"member_ids"`

	// The mention name of the Group.
	// Required: true
	// Min Length: 1
	// Pattern: ^[a-z0-9\-\_\.]+$
	MentionName *string `json:"mention_name"`

	// The name of the Group.
	// Required: true
	Name *string `json:"name"`

	// The number of epics assigned to the group which are in the started workflow state.
	// Required: true
	NumEpicsStarted *int64 `json:"num_epics_started"`

	// The total number of stories assigned ot the group.
	// Required: true
	NumStories *int64 `json:"num_stories"`

	// The number of stories assigned to the group which are in a started workflow state.
	// Required: true
	NumStoriesStarted *int64 `json:"num_stories_started"`

	// The Workflow IDs contained within the Group.
	// Required: true
	WorkflowIds []int64 `json:"workflow_ids"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColorKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateNumEpicsStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	if err := validate.MinLength("color", "body", *m.Color, 1); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", *m.Color, `^#[a-fA-F0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

var groupTypeColorKeyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blue","purple","midnight-blue","orange","yellow-green","brass","gray","fuchsia","yellow","pink","sky-blue","green","red","black","slate","turquoise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupTypeColorKeyPropEnum = append(groupTypeColorKeyPropEnum, v)
	}
}

const (

	// GroupColorKeyBlue captures enum value "blue"
	GroupColorKeyBlue string = "blue"

	// GroupColorKeyPurple captures enum value "purple"
	GroupColorKeyPurple string = "purple"

	// GroupColorKeyMidnightDashBlue captures enum value "midnight-blue"
	GroupColorKeyMidnightDashBlue string = "midnight-blue"

	// GroupColorKeyOrange captures enum value "orange"
	GroupColorKeyOrange string = "orange"

	// GroupColorKeyYellowDashGreen captures enum value "yellow-green"
	GroupColorKeyYellowDashGreen string = "yellow-green"

	// GroupColorKeyBrass captures enum value "brass"
	GroupColorKeyBrass string = "brass"

	// GroupColorKeyGray captures enum value "gray"
	GroupColorKeyGray string = "gray"

	// GroupColorKeyFuchsia captures enum value "fuchsia"
	GroupColorKeyFuchsia string = "fuchsia"

	// GroupColorKeyYellow captures enum value "yellow"
	GroupColorKeyYellow string = "yellow"

	// GroupColorKeyPink captures enum value "pink"
	GroupColorKeyPink string = "pink"

	// GroupColorKeySkyDashBlue captures enum value "sky-blue"
	GroupColorKeySkyDashBlue string = "sky-blue"

	// GroupColorKeyGreen captures enum value "green"
	GroupColorKeyGreen string = "green"

	// GroupColorKeyRed captures enum value "red"
	GroupColorKeyRed string = "red"

	// GroupColorKeyBlack captures enum value "black"
	GroupColorKeyBlack string = "black"

	// GroupColorKeySlate captures enum value "slate"
	GroupColorKeySlate string = "slate"

	// GroupColorKeyTurquoise captures enum value "turquoise"
	GroupColorKeyTurquoise string = "turquoise"
)

// prop value enum
func (m *Group) validateColorKeyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupTypeColorKeyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Group) validateColorKey(formats strfmt.Registry) error {

	if err := validate.Required("color_key", "body", m.ColorKey); err != nil {
		return err
	}

	// value enum
	if err := m.validateColorKeyEnum("color_key", "body", *m.ColorKey); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateDisplayIcon(formats strfmt.Registry) error {

	if err := validate.Required("display_icon", "body", m.DisplayIcon); err != nil {
		return err
	}

	if m.DisplayIcon != nil {
		if err := m.DisplayIcon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_icon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_icon")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateMemberIds(formats strfmt.Registry) error {

	if err := validate.Required("member_ids", "body", m.MemberIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberIds); i++ {

		if err := validate.FormatOf("member_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MemberIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Group) validateMentionName(formats strfmt.Registry) error {

	if err := validate.Required("mention_name", "body", m.MentionName); err != nil {
		return err
	}

	if err := validate.MinLength("mention_name", "body", *m.MentionName, 1); err != nil {
		return err
	}

	if err := validate.Pattern("mention_name", "body", *m.MentionName, `^[a-z0-9\-\_\.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateNumEpicsStarted(formats strfmt.Registry) error {

	if err := validate.Required("num_epics_started", "body", m.NumEpicsStarted); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateNumStories(formats strfmt.Registry) error {

	if err := validate.Required("num_stories", "body", m.NumStories); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateNumStoriesStarted(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_started", "body", m.NumStoriesStarted); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateWorkflowIds(formats strfmt.Registry) error {

	if err := validate.Required("workflow_ids", "body", m.WorkflowIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this group based on the context it is used
func (m *Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayIcon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) contextValidateDisplayIcon(ctx context.Context, formats strfmt.Registry) error {

	if m.DisplayIcon != nil {
		if err := m.DisplayIcon.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_icon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_icon")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Group) UnmarshalBinary(b []byte) error {
	var res Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
