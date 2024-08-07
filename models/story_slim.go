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

// StorySlim StorySlim represents the same resource as a Story, but is more light-weight. For certain fields it provides ids rather than full resources (e.g., `comment_ids` and `file_ids`) and it also excludes certain aggregate values (e.g., `cycle_time`). The `description` field can be optionally included. Use the [Get Story](#Get-Story) endpoint to fetch the unabridged payload for a Story.
//
// swagger:model StorySlim
type StorySlim struct {

	// The Shortcut application url for the Story.
	// Required: true
	AppURL *string `json:"app_url"`

	// True if the story has been archived or not.
	// Required: true
	Archived *bool `json:"archived"`

	// A true/false boolean indicating if the Story is currently blocked.
	// Required: true
	Blocked *bool `json:"blocked"`

	// A true/false boolean indicating if the Story is currently a blocker of another story.
	// Required: true
	Blocker *bool `json:"blocker"`

	// An array of IDs of Comments attached to the story.
	// Required: true
	CommentIds []int64 `json:"comment_ids"`

	// A true/false boolean indicating if the Story has been completed.
	// Required: true
	Completed *bool `json:"completed"`

	// The time/date the Story was completed.
	// Required: true
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at"`

	// A manual override for the time/date the Story was completed.
	// Required: true
	// Format: date-time
	CompletedAtOverride *strfmt.DateTime `json:"completed_at_override"`

	// The time/date the Story was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// An array of CustomField value assertions for the story.
	CustomFields []*StoryCustomField `json:"custom_fields"`

	// The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time,omitempty"`

	// The due date of the story.
	// Required: true
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline"`

	// The description of the Story.
	Description string `json:"description,omitempty"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the epic the story belongs to.
	// Required: true
	EpicID *int64 `json:"epic_id"`

	// The numeric point estimate of the story. Can also be null, which means unestimated.
	// Required: true
	Estimate *int64 `json:"estimate"`

	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	// Required: true
	ExternalID *string `json:"external_id"`

	// An array of external links (strings) associated with a Story
	// Required: true
	ExternalLinks []string `json:"external_links"`

	// An array of IDs of Files attached to the story.
	// Required: true
	FileIds []int64 `json:"file_ids"`

	// An array of UUIDs for any Members listed as Followers.
	// Required: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// global id
	// Required: true
	GlobalID *string `json:"global_id"`

	// The ID of the group associated with the story.
	// Required: true
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id"`

	// An array of Group IDs that have been mentioned in the Story description.
	// Required: true
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The unique ID of the Story.
	// Required: true
	ID *int64 `json:"id"`

	// The ID of the iteration the story belongs to.
	// Required: true
	IterationID *int64 `json:"iteration_id"`

	// An array of label ids attached to the story.
	// Required: true
	LabelIds []int64 `json:"label_ids"`

	// An array of labels attached to the story.
	// Required: true
	Labels []*LabelSlim `json:"labels"`

	// The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time,omitempty"`

	// An array of IDs of LinkedFiles attached to the story.
	// Required: true
	LinkedFileIds []int64 `json:"linked_file_ids"`

	// An array of Member IDs that have been mentioned in the Story description.
	// Required: true
	MemberMentionIds []strfmt.UUID `json:"member_mention_ids"`

	// `Deprecated:` use `member_mention_ids`.
	// Required: true
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// The time/date the Story was last changed workflow-state.
	// Required: true
	// Format: date-time
	MovedAt *strfmt.DateTime `json:"moved_at"`

	// The name of the story.
	// Required: true
	Name *string `json:"name"`

	// The number of tasks on the story which are complete.
	// Required: true
	NumTasksCompleted *int64 `json:"num_tasks_completed"`

	// An array of UUIDs of the owners of this story.
	// Required: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// A number representing the position of the story in relation to every other story in the current project.
	// Required: true
	Position *int64 `json:"position"`

	// The IDs of the iteration the story belongs to.
	// Required: true
	PreviousIterationIds []int64 `json:"previous_iteration_ids"`

	// The ID of the project the story belongs to.
	// Required: true
	ProjectID *int64 `json:"project_id"`

	// The ID of the Member that requested the story.
	// Required: true
	// Format: uuid
	RequestedByID *strfmt.UUID `json:"requested_by_id"`

	// A true/false boolean indicating if the Story has been started.
	// Required: true
	Started *bool `json:"started"`

	// The time/date the Story was started.
	// Required: true
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at"`

	// A manual override for the time/date the Story was started.
	// Required: true
	// Format: date-time
	StartedAtOverride *strfmt.DateTime `json:"started_at_override"`

	// stats
	// Required: true
	Stats *StoryStats `json:"stats"`

	// An array of story links attached to the Story.
	// Required: true
	StoryLinks []*TypedStoryLink `json:"story_links"`

	// The ID of the story template used to create this story, or null if not created using a template.
	// Required: true
	// Format: uuid
	StoryTemplateID *strfmt.UUID `json:"story_template_id"`

	// The type of story (feature, bug, chore).
	// Required: true
	StoryType *string `json:"story_type"`

	// synced item
	SyncedItem *SyncedItem `json:"synced_item,omitempty"`

	// An array of IDs of Tasks attached to the story.
	// Required: true
	TaskIds []int64 `json:"task_ids"`

	// The IDs of any unresolved blocker comments on the Story.
	UnresolvedBlockerComments []int64 `json:"unresolved_blocker_comments"`

	// The time/date the Story was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The ID of the workflow the story belongs to.
	// Required: true
	WorkflowID *int64 `json:"workflow_id"`

	// The ID of the workflow state the story is currently in.
	// Required: true
	WorkflowStateID *int64 `json:"workflow_state_id"`
}

// Validate validates this story slim
func (m *StorySlim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommentIds(formats); err != nil {
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

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpicID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIds(formats); err != nil {
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

	if err := m.validateGroupMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumTasksCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousIterationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
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

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncedItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowStateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorySlim) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateBlocker(formats strfmt.Registry) error {

	if err := validate.Required("blocker", "body", m.Blocker); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCommentIds(formats strfmt.Registry) error {

	if err := validate.Required("comment_ids", "body", m.CommentIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCompletedAt(formats strfmt.Registry) error {

	if err := validate.Required("completed_at", "body", m.CompletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCompletedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("completed_at_override", "body", m.CompletedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateCustomFields(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorySlim) validateDeadline(formats strfmt.Registry) error {

	if err := validate.Required("deadline", "body", m.Deadline); err != nil {
		return err
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateEpicID(formats strfmt.Registry) error {

	if err := validate.Required("epic_id", "body", m.EpicID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateEstimate(formats strfmt.Registry) error {

	if err := validate.Required("estimate", "body", m.Estimate); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateExternalLinks(formats strfmt.Registry) error {

	if err := validate.Required("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateFileIds(formats strfmt.Registry) error {

	if err := validate.Required("file_ids", "body", m.FileIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateFollowerIds(formats strfmt.Registry) error {

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

func (m *StorySlim) validateGlobalID(formats strfmt.Registry) error {

	if err := validate.Required("global_id", "body", m.GlobalID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateGroupMentionIds(formats strfmt.Registry) error {

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

func (m *StorySlim) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateIterationID(formats strfmt.Registry) error {

	if err := validate.Required("iteration_id", "body", m.IterationID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateLabelIds(formats strfmt.Registry) error {

	if err := validate.Required("label_ids", "body", m.LabelIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateLabels(formats strfmt.Registry) error {

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

func (m *StorySlim) validateLinkedFileIds(formats strfmt.Registry) error {

	if err := validate.Required("linked_file_ids", "body", m.LinkedFileIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateMemberMentionIds(formats strfmt.Registry) error {

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

func (m *StorySlim) validateMentionIds(formats strfmt.Registry) error {

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

func (m *StorySlim) validateMovedAt(formats strfmt.Registry) error {

	if err := validate.Required("moved_at", "body", m.MovedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("moved_at", "body", "date-time", m.MovedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateNumTasksCompleted(formats strfmt.Registry) error {

	if err := validate.Required("num_tasks_completed", "body", m.NumTasksCompleted); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateOwnerIds(formats strfmt.Registry) error {

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

func (m *StorySlim) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validatePreviousIterationIds(formats strfmt.Registry) error {

	if err := validate.Required("previous_iteration_ids", "body", m.PreviousIterationIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateRequestedByID(formats strfmt.Registry) error {

	if err := validate.Required("requested_by_id", "body", m.RequestedByID); err != nil {
		return err
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateStarted(formats strfmt.Registry) error {

	if err := validate.Required("started", "body", m.Started); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", m.StartedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateStartedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("started_at_override", "body", m.StartedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateStats(formats strfmt.Registry) error {

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

func (m *StorySlim) validateStoryLinks(formats strfmt.Registry) error {

	if err := validate.Required("story_links", "body", m.StoryLinks); err != nil {
		return err
	}

	for i := 0; i < len(m.StoryLinks); i++ {
		if swag.IsZero(m.StoryLinks[i]) { // not required
			continue
		}

		if m.StoryLinks[i] != nil {
			if err := m.StoryLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("story_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("story_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorySlim) validateStoryTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("story_template_id", "body", m.StoryTemplateID); err != nil {
		return err
	}

	if err := validate.FormatOf("story_template_id", "body", "uuid", m.StoryTemplateID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateStoryType(formats strfmt.Registry) error {

	if err := validate.Required("story_type", "body", m.StoryType); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateSyncedItem(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncedItem) { // not required
		return nil
	}

	if m.SyncedItem != nil {
		if err := m.SyncedItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synced_item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("synced_item")
			}
			return err
		}
	}

	return nil
}

func (m *StorySlim) validateTaskIds(formats strfmt.Registry) error {

	if err := validate.Required("task_ids", "body", m.TaskIds); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateWorkflowID(formats strfmt.Registry) error {

	if err := validate.Required("workflow_id", "body", m.WorkflowID); err != nil {
		return err
	}

	return nil
}

func (m *StorySlim) validateWorkflowStateID(formats strfmt.Registry) error {

	if err := validate.Required("workflow_state_id", "body", m.WorkflowStateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this story slim based on the context it is used
func (m *StorySlim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoryLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncedItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorySlim) contextValidateCustomFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFields); i++ {

		if m.CustomFields[i] != nil {

			if swag.IsZero(m.CustomFields[i]) { // not required
				return nil
			}

			if err := m.CustomFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorySlim) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StorySlim) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StorySlim) contextValidateStoryLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StoryLinks); i++ {

		if m.StoryLinks[i] != nil {

			if swag.IsZero(m.StoryLinks[i]) { // not required
				return nil
			}

			if err := m.StoryLinks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("story_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("story_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorySlim) contextValidateSyncedItem(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncedItem != nil {

		if swag.IsZero(m.SyncedItem) { // not required
			return nil
		}

		if err := m.SyncedItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synced_item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("synced_item")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorySlim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorySlim) UnmarshalBinary(b []byte) error {
	var res StorySlim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
