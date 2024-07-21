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

// EpicSearchResult An Epic in search results. This is typed differently from Epic because the details=slim search argument will omit some fields.
//
// swagger:model EpicSearchResult
type EpicSearchResult struct {

	// The Shortcut application url for the Epic.
	// Required: true
	AppURL *string `json:"app_url"`

	// True/false boolean that indicates whether the Epic is archived or not.
	// Required: true
	Archived *bool `json:"archived"`

	// An array containing Group IDs and Group-owned story counts for the Epic's associated groups.
	// Required: true
	AssociatedGroups []*EpicAssociatedGroup `json:"associated_groups"`

	// A nested array of threaded comments.
	Comments []*ThreadedComment `json:"comments"`

	// A true/false boolean indicating if the Epic has been completed.
	// Required: true
	Completed *bool `json:"completed"`

	// The time/date the Epic was completed.
	// Required: true
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at"`

	// A manual override for the time/date the Epic was completed.
	// Required: true
	// Format: date-time
	CompletedAtOverride *strfmt.DateTime `json:"completed_at_override"`

	// The time/date the Epic was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The Epic's deadline.
	// Required: true
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline"`

	// The Epic's description.
	Description string `json:"description,omitempty"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the Epic State.
	// Required: true
	EpicStateID *int64 `json:"epic_state_id"`

	// This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	// Required: true
	ExternalID *string `json:"external_id"`

	// An array of UUIDs for any Members you want to add as Followers on this Epic.
	// Required: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// global id
	// Required: true
	GlobalID *string `json:"global_id"`

	// `Deprecated` The ID of the group to associate with the epic. Use `group_ids`.
	// Required: true
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id"`

	// An array of UUIDS for Groups to which this Epic is related.
	// Required: true
	GroupIds []strfmt.UUID `json:"group_ids"`

	// An array of Group IDs that have been mentioned in the Epic description.
	// Required: true
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The unique ID of the Epic.
	// Required: true
	ID *int64 `json:"id"`

	// An array of Label ids attached to the Epic.
	// Required: true
	LabelIds []int64 `json:"label_ids"`

	// An array of Labels attached to the Epic.
	// Required: true
	Labels []*LabelSlim `json:"labels"`

	// An array of Member IDs that have been mentioned in the Epic description.
	// Required: true
	MemberMentionIds []strfmt.UUID `json:"member_mention_ids"`

	// `Deprecated:` use `member_mention_ids`.
	// Required: true
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// `Deprecated` The ID of the Objective this Epic is related to. Use `objective_ids`.
	// Required: true
	MilestoneID *int64 `json:"milestone_id"`

	// The name of the Epic.
	// Required: true
	Name *string `json:"name"`

	// An array of IDs for Objectives to which this epic is related.
	// Required: true
	ObjectiveIds []int64 `json:"objective_ids"`

	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	// Required: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The Epic's planned start date.
	// Required: true
	// Format: date-time
	PlannedStartDate *strfmt.DateTime `json:"planned_start_date"`

	// The Epic's relative position in the Epic workflow state.
	// Required: true
	Position *int64 `json:"position"`

	// The ID of the associated productboard feature.
	// Required: true
	// Format: uuid
	ProductboardID *strfmt.UUID `json:"productboard_id"`

	// The name of the associated productboard feature.
	// Required: true
	ProductboardName *string `json:"productboard_name"`

	// The ID of the associated productboard integration.
	// Required: true
	// Format: uuid
	ProductboardPluginID *strfmt.UUID `json:"productboard_plugin_id"`

	// The URL of the associated productboard feature.
	// Required: true
	ProductboardURL *string `json:"productboard_url"`

	// The IDs of Projects related to this Epic.
	// Required: true
	ProjectIds []int64 `json:"project_ids"`

	// The ID of the Member that requested the epic.
	// Required: true
	// Format: uuid
	RequestedByID *strfmt.UUID `json:"requested_by_id"`

	// A true/false boolean indicating if the Epic has been started.
	// Required: true
	Started *bool `json:"started"`

	// The time/date the Epic was started.
	// Required: true
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at"`

	// A manual override for the time/date the Epic was started.
	// Required: true
	// Format: date-time
	StartedAtOverride *strfmt.DateTime `json:"started_at_override"`

	// `Deprecated` The workflow state that the Epic is in.
	// Required: true
	State *string `json:"state"`

	// stats
	// Required: true
	Stats *EpicStats `json:"stats"`

	// The number of stories in this epic which are not associated with a project.
	// Required: true
	StoriesWithoutProjects *int64 `json:"stories_without_projects"`

	// The time/date the Epic was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this epic search result
func (m *EpicSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
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

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpicStateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
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

	if err := m.validateMilestoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectiveIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlannedStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductboardName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductboardPluginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductboardURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
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

	if err := m.validateStoriesWithoutProjects(formats); err != nil {
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

func (m *EpicSearchResult) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateAssociatedGroups(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateComments(formats strfmt.Registry) error {
	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EpicSearchResult) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateCompletedAt(formats strfmt.Registry) error {

	if err := validate.Required("completed_at", "body", m.CompletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateCompletedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("completed_at_override", "body", m.CompletedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateDeadline(formats strfmt.Registry) error {

	if err := validate.Required("deadline", "body", m.Deadline); err != nil {
		return err
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateEpicStateID(formats strfmt.Registry) error {

	if err := validate.Required("epic_state_id", "body", m.EpicStateID); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateFollowerIds(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateGlobalID(formats strfmt.Registry) error {

	if err := validate.Required("global_id", "body", m.GlobalID); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateGroupIds(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateGroupMentionIds(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateLabelIds(formats strfmt.Registry) error {

	if err := validate.Required("label_ids", "body", m.LabelIds); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateLabels(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateMemberMentionIds(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateMentionIds(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateMilestoneID(formats strfmt.Registry) error {

	if err := validate.Required("milestone_id", "body", m.MilestoneID); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateObjectiveIds(formats strfmt.Registry) error {

	if err := validate.Required("objective_ids", "body", m.ObjectiveIds); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateOwnerIds(formats strfmt.Registry) error {

	if err := validate.Required("owner_ids", "body", m.OwnerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *EpicSearchResult) validatePlannedStartDate(formats strfmt.Registry) error {

	if err := validate.Required("planned_start_date", "body", m.PlannedStartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("planned_start_date", "body", "date-time", m.PlannedStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateProductboardID(formats strfmt.Registry) error {

	if err := validate.Required("productboard_id", "body", m.ProductboardID); err != nil {
		return err
	}

	if err := validate.FormatOf("productboard_id", "body", "uuid", m.ProductboardID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateProductboardName(formats strfmt.Registry) error {

	if err := validate.Required("productboard_name", "body", m.ProductboardName); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateProductboardPluginID(formats strfmt.Registry) error {

	if err := validate.Required("productboard_plugin_id", "body", m.ProductboardPluginID); err != nil {
		return err
	}

	if err := validate.FormatOf("productboard_plugin_id", "body", "uuid", m.ProductboardPluginID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateProductboardURL(formats strfmt.Registry) error {

	if err := validate.Required("productboard_url", "body", m.ProductboardURL); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateProjectIds(formats strfmt.Registry) error {

	if err := validate.Required("project_ids", "body", m.ProjectIds); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateRequestedByID(formats strfmt.Registry) error {

	if err := validate.Required("requested_by_id", "body", m.RequestedByID); err != nil {
		return err
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateStarted(formats strfmt.Registry) error {

	if err := validate.Required("started", "body", m.Started); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", m.StartedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateStartedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("started_at_override", "body", m.StartedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateStats(formats strfmt.Registry) error {

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

func (m *EpicSearchResult) validateStoriesWithoutProjects(formats strfmt.Registry) error {

	if err := validate.Required("stories_without_projects", "body", m.StoriesWithoutProjects); err != nil {
		return err
	}

	return nil
}

func (m *EpicSearchResult) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this epic search result based on the context it is used
func (m *EpicSearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssociatedGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComments(ctx, formats); err != nil {
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

func (m *EpicSearchResult) contextValidateAssociatedGroups(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EpicSearchResult) contextValidateComments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Comments); i++ {

		if m.Comments[i] != nil {

			if swag.IsZero(m.Comments[i]) { // not required
				return nil
			}

			if err := m.Comments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EpicSearchResult) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EpicSearchResult) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EpicSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EpicSearchResult) UnmarshalBinary(b []byte) error {
	var res EpicSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
