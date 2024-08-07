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

// IterationSlim IterationSlim represents the same resource as an Iteration, but is more light-weight. Use the [Get Iteration](#Get-Iteration) endpoint to fetch the unabridged payload for an Iteration.
//
// swagger:model IterationSlim
type IterationSlim struct {

	// The Shortcut application url for the Iteration.
	// Required: true
	AppURL *string `json:"app_url"`

	// An array containing Group IDs and Group-owned story counts for the Iteration's associated groups.
	// Required: true
	AssociatedGroups []*IterationAssociatedGroup `json:"associated_groups"`

	// The instant when this iteration was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The date this iteration ends.
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"end_date"`

	// A string description of this resource
	// Required: true
	EntityType *string `json:"entity_type"`

	// An array of UUIDs for any Members listed as Followers.
	// Required: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// global id
	// Required: true
	GlobalID *string `json:"global_id"`

	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	// Required: true
	GroupIds []strfmt.UUID `json:"group_ids"`

	// An array of Group IDs that have been mentioned in the Story description.
	// Required: true
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The ID of the iteration.
	// Required: true
	ID *int64 `json:"id"`

	// An array of label ids attached to the iteration.
	// Required: true
	LabelIds []int64 `json:"label_ids"`

	// An array of labels attached to the iteration.
	// Required: true
	Labels []*Label `json:"labels"`

	// An array of Member IDs that have been mentioned in the Story description.
	// Required: true
	MemberMentionIds []strfmt.UUID `json:"member_mention_ids"`

	// `Deprecated:` use `member_mention_ids`.
	// Required: true
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// The name of the iteration.
	// Required: true
	Name *string `json:"name"`

	// The date this iteration begins.
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"start_date"`

	// stats
	// Required: true
	Stats *IterationStats `json:"stats"`

	// The status of the iteration. Values are either "unstarted", "started", or "done".
	// Required: true
	Status *string `json:"status"`

	// The instant when this iteration was last updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this iteration slim
func (m *IterationSlim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *IterationSlim) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateAssociatedGroups(formats strfmt.Registry) error {

	if err := validate.Required("associated_groups", "body", m.AssociatedGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.AssociatedGroups); i++ {
		if swag.IsZero(m.AssociatedGroups[i]) { // not required
			continue
		}

		if m.AssociatedGroups[i] != nil {
			if err := m.AssociatedGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associated_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("associated_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IterationSlim) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("end_date", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("end_date", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateFollowerIds(formats strfmt.Registry) error {

	if err := validate.Required("follower_ids", "body", m.FollowerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IterationSlim) validateGlobalID(formats strfmt.Registry) error {

	if err := validate.Required("global_id", "body", m.GlobalID); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateGroupIds(formats strfmt.Registry) error {

	if err := validate.Required("group_ids", "body", m.GroupIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupIds); i++ {

		if err := validate.FormatOf("group_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IterationSlim) validateGroupMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("group_mention_ids", "body", m.GroupMentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupMentionIds); i++ {

		if err := validate.FormatOf("group_mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupMentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IterationSlim) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateLabelIds(formats strfmt.Registry) error {

	if err := validate.Required("label_ids", "body", m.LabelIds); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
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

func (m *IterationSlim) validateMemberMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("member_mention_ids", "body", m.MemberMentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberMentionIds); i++ {

		if err := validate.FormatOf("member_mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MemberMentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IterationSlim) validateMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("mention_ids", "body", m.MentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MentionIds); i++ {

		if err := validate.FormatOf("mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IterationSlim) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("start_date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateStats(formats strfmt.Registry) error {

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

func (m *IterationSlim) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IterationSlim) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this iteration slim based on the context it is used
func (m *IterationSlim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssociatedGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
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

func (m *IterationSlim) contextValidateAssociatedGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AssociatedGroups); i++ {

		if m.AssociatedGroups[i] != nil {

			if swag.IsZero(m.AssociatedGroups[i]) { // not required
				return nil
			}

			if err := m.AssociatedGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associated_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("associated_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IterationSlim) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IterationSlim) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IterationSlim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IterationSlim) UnmarshalBinary(b []byte) error {
	var res IterationSlim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
