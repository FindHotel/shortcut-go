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

// CreateStoryFromTemplateParams Request parameters for creating a story from a story template. These parameters are merged with the values derived from the template.
//
// swagger:model CreateStoryFromTemplateParams
type CreateStoryFromTemplateParams struct {

	// Controls the story's archived state.
	Archived bool `json:"archived,omitempty"`

	// An array of comments to add to the story.
	Comments []*CreateStoryCommentParams `json:"comments"`

	// A manual override for the time/date the Story was completed.
	// Format: date-time
	CompletedAtOverride strfmt.DateTime `json:"completed_at_override,omitempty"`

	// The time/date the Story was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFields []*CustomFieldValueParams `json:"custom_fields"`

	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. These will be added to any fields provided by the template. Cannot be used in conjunction with `custom_fields`.
	// Unique: true
	CustomFieldsAdd []*CustomFieldValueParams `json:"custom_fields_add"`

	// A map specifying a CustomField ID. These will be removed from any fields provided by the template. Cannot be used in conjunction with `custom_fields`.
	// Unique: true
	CustomFieldsRemove []*RemoveCustomFieldParams `json:"custom_fields_remove"`

	// The due date of the story.
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline,omitempty"`

	// The description of the story.
	// Max Length: 100000
	Description string `json:"description,omitempty"`

	// The ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id,omitempty"`

	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`

	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`

	// An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links"`

	// An array of External Links associated with this story. These will be added to any links provided by the template. Cannot be used in conjunction with `external_links`.
	// Unique: true
	ExternalLinksAdd []string `json:"external_links_add"`

	// An array of External Links associated with this story. These will be removed from any links provided by the template. Cannot be used in conjunction with `external_links`.
	// Unique: true
	ExternalLinksRemove []string `json:"external_links_remove"`

	// An array of IDs of files attached to the story.
	// Unique: true
	FileIds []int64 `json:"file_ids"`

	// An array of IDs of files attached to the story in addition to files from the template. Cannot be used in conjunction with `file_ids`.
	// Unique: true
	FileIdsAdd []int64 `json:"file_ids_add"`

	// An array of IDs of files removed from files from the template. Cannot be used in conjunction with `file_ids`.
	// Unique: true
	FileIdsRemove []int64 `json:"file_ids_remove"`

	// An array of UUIDs of the followers of this story.
	// Unique: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The UUIDs of the new followers to be added in addition to followers from the template. Cannot be used in conjunction with `follower_ids.`
	// Unique: true
	FollowerIdsAdd []strfmt.UUID `json:"follower_ids_add"`

	// The UUIDs of the new followers to be removed from followers from the template. Cannot be used in conjunction with `follower_ids`.
	// Unique: true
	FollowerIdsRemove []strfmt.UUID `json:"follower_ids_remove"`

	// The id of the group to associate with this story.
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id,omitempty"`

	// The ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id,omitempty"`

	// An array of labels attached to the story.
	Labels []*CreateLabelParams `json:"labels"`

	// An array of labels attached to the story in addition to the labels provided by the template. Cannot be used in conjunction with `labels`.
	// Unique: true
	LabelsAdd []*CreateLabelParams `json:"labels_add"`

	// An array of labels to remove from the labels provided by the template. Cannot be used in conjunction with `labels`.
	// Unique: true
	LabelsRemove []*RemoveLabelParams `json:"labels_remove"`

	// An array of IDs of linked files attached to the story.
	// Unique: true
	LinkedFileIds []int64 `json:"linked_file_ids"`

	// An array of IDs of linked files attached to the story in addition to files from the template. Cannot be used in conjunction with `linked_files`.
	// Unique: true
	LinkedFileIdsAdd []int64 `json:"linked_file_ids_add"`

	// An array of IDs of linked files removed from files from the template. Cannot be used in conjunction with `linked_files.`
	// Unique: true
	LinkedFileIdsRemove []int64 `json:"linked_file_ids_remove"`

	// One of "first" or "last". This can be used to move the given story to the first or last position in the workflow state.
	// Enum: ["last","first"]
	MoveTo string `json:"move_to,omitempty"`

	// The name of the story. Must be provided if the template does not provide a name.
	// Max Length: 512
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// An array of UUIDs of the owners of this story.
	// Unique: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The UUIDs of the new owners to be added in addition to owners from the template. Cannot be used in conjunction with `owners`.
	// Unique: true
	OwnerIdsAdd []strfmt.UUID `json:"owner_ids_add"`

	// The UUIDs of the new owners to be removed from owners from the template. Cannot be used in conjunction with `owners`.
	// Unique: true
	OwnerIdsRemove []strfmt.UUID `json:"owner_ids_remove"`

	// The ID of the project the story belongs to.
	ProjectID *int64 `json:"project_id,omitempty"`

	// The ID of the member that requested the story.
	// Format: uuid
	RequestedByID strfmt.UUID `json:"requested_by_id,omitempty"`

	// Given this story was converted from a task in another story, this is the original task ID that was converted to this story.
	SourceTaskID *int64 `json:"source_task_id,omitempty"`

	// A manual override for the time/date the Story was started.
	// Format: date-time
	StartedAtOverride strfmt.DateTime `json:"started_at_override,omitempty"`

	// An array of story links attached to the story.
	StoryLinks []*CreateStoryLinkParams `json:"story_links"`

	// The id of the story template used to create this story.
	// Required: true
	// Format: uuid
	StoryTemplateID *strfmt.UUID `json:"story_template_id"`

	// The type of story (feature, bug, chore).
	// Enum: ["feature","chore","bug"]
	StoryType string `json:"story_type,omitempty"`

	// An array of tasks connected to the story.
	Tasks []*CreateTaskParams `json:"tasks"`

	// The time/date the Story was updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The ID of the workflow state the story will be in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

// Validate validates this create story from template params
func (m *CreateStoryFromTemplateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComments(formats); err != nil {
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

	if err := m.validateCustomFieldsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFieldsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinksAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinksRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIdsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIdsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIdsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIdsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIdsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIdsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIdsAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIdsRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
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

	if err := m.validateTasks(formats); err != nil {
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

func (m *CreateStoryFromTemplateParams) validateComments(formats strfmt.Registry) error {
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

func (m *CreateStoryFromTemplateParams) validateCompletedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateCustomFields(formats strfmt.Registry) error {
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

func (m *CreateStoryFromTemplateParams) validateCustomFieldsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFieldsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("custom_fields_add", "body", m.CustomFieldsAdd); err != nil {
		return err
	}

	for i := 0; i < len(m.CustomFieldsAdd); i++ {
		if swag.IsZero(m.CustomFieldsAdd[i]) { // not required
			continue
		}

		if m.CustomFieldsAdd[i] != nil {
			if err := m.CustomFieldsAdd[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_add" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields_add" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateCustomFieldsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFieldsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("custom_fields_remove", "body", m.CustomFieldsRemove); err != nil {
		return err
	}

	for i := 0; i < len(m.CustomFieldsRemove); i++ {
		if swag.IsZero(m.CustomFieldsRemove[i]) { // not required
			continue
		}

		if m.CustomFieldsRemove[i] != nil {
			if err := m.CustomFieldsRemove[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_remove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields_remove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 100000); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateExternalLinksAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalLinksAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("external_links_add", "body", m.ExternalLinksAdd); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateExternalLinksRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalLinksRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("external_links_remove", "body", m.ExternalLinksRemove); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFileIds(formats strfmt.Registry) error {
	if swag.IsZero(m.FileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("file_ids", "body", m.FileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFileIdsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.FileIdsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("file_ids_add", "body", m.FileIdsAdd); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFileIdsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.FileIdsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("file_ids_remove", "body", m.FileIdsRemove); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFollowerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.FollowerIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("follower_ids", "body", m.FollowerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFollowerIdsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.FollowerIdsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("follower_ids_add", "body", m.FollowerIdsAdd); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIdsAdd); i++ {

		if err := validate.FormatOf("follower_ids_add"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIdsAdd[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateFollowerIdsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.FollowerIdsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("follower_ids_remove", "body", m.FollowerIdsRemove); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIdsRemove); i++ {

		if err := validate.FormatOf("follower_ids_remove"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIdsRemove[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
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

func (m *CreateStoryFromTemplateParams) validateLabelsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("labels_add", "body", m.LabelsAdd); err != nil {
		return err
	}

	for i := 0; i < len(m.LabelsAdd); i++ {
		if swag.IsZero(m.LabelsAdd[i]) { // not required
			continue
		}

		if m.LabelsAdd[i] != nil {
			if err := m.LabelsAdd[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels_add" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels_add" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateLabelsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("labels_remove", "body", m.LabelsRemove); err != nil {
		return err
	}

	for i := 0; i < len(m.LabelsRemove); i++ {
		if swag.IsZero(m.LabelsRemove[i]) { // not required
			continue
		}

		if m.LabelsRemove[i] != nil {
			if err := m.LabelsRemove[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels_remove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels_remove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateLinkedFileIds(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedFileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("linked_file_ids", "body", m.LinkedFileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateLinkedFileIdsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedFileIdsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("linked_file_ids_add", "body", m.LinkedFileIdsAdd); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateLinkedFileIdsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedFileIdsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("linked_file_ids_remove", "body", m.LinkedFileIdsRemove); err != nil {
		return err
	}

	return nil
}

var createStoryFromTemplateParamsTypeMoveToPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["last","first"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createStoryFromTemplateParamsTypeMoveToPropEnum = append(createStoryFromTemplateParamsTypeMoveToPropEnum, v)
	}
}

const (

	// CreateStoryFromTemplateParamsMoveToLast captures enum value "last"
	CreateStoryFromTemplateParamsMoveToLast string = "last"

	// CreateStoryFromTemplateParamsMoveToFirst captures enum value "first"
	CreateStoryFromTemplateParamsMoveToFirst string = "first"
)

// prop value enum
func (m *CreateStoryFromTemplateParams) validateMoveToEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createStoryFromTemplateParamsTypeMoveToPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateStoryFromTemplateParams) validateMoveTo(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveTo) { // not required
		return nil
	}

	// value enum
	if err := m.validateMoveToEnum("move_to", "body", m.MoveTo); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 512); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateOwnerIds(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("owner_ids", "body", m.OwnerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateOwnerIdsAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIdsAdd) { // not required
		return nil
	}

	if err := validate.UniqueItems("owner_ids_add", "body", m.OwnerIdsAdd); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIdsAdd); i++ {

		if err := validate.FormatOf("owner_ids_add"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIdsAdd[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateOwnerIdsRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerIdsRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("owner_ids_remove", "body", m.OwnerIdsRemove); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIdsRemove); i++ {

		if err := validate.FormatOf("owner_ids_remove"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIdsRemove[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateRequestedByID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedByID) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateStartedAtOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateStoryLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.StoryLinks) { // not required
		return nil
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

func (m *CreateStoryFromTemplateParams) validateStoryTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("story_template_id", "body", m.StoryTemplateID); err != nil {
		return err
	}

	if err := validate.FormatOf("story_template_id", "body", "uuid", m.StoryTemplateID.String(), formats); err != nil {
		return err
	}

	return nil
}

var createStoryFromTemplateParamsTypeStoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feature","chore","bug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createStoryFromTemplateParamsTypeStoryTypePropEnum = append(createStoryFromTemplateParamsTypeStoryTypePropEnum, v)
	}
}

const (

	// CreateStoryFromTemplateParamsStoryTypeFeature captures enum value "feature"
	CreateStoryFromTemplateParamsStoryTypeFeature string = "feature"

	// CreateStoryFromTemplateParamsStoryTypeChore captures enum value "chore"
	CreateStoryFromTemplateParamsStoryTypeChore string = "chore"

	// CreateStoryFromTemplateParamsStoryTypeBug captures enum value "bug"
	CreateStoryFromTemplateParamsStoryTypeBug string = "bug"
)

// prop value enum
func (m *CreateStoryFromTemplateParams) validateStoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createStoryFromTemplateParamsTypeStoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateStoryFromTemplateParams) validateStoryType(formats strfmt.Registry) error {
	if swag.IsZero(m.StoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateStoryTypeEnum("story_type", "body", m.StoryType); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create story from template params based on the context it is used
func (m *CreateStoryFromTemplateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFieldsAdd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFieldsRemove(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsAdd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsRemove(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoryLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStoryFromTemplateParams) contextValidateComments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateStoryFromTemplateParams) contextValidateCustomFields(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateStoryFromTemplateParams) contextValidateCustomFieldsAdd(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFieldsAdd); i++ {

		if m.CustomFieldsAdd[i] != nil {

			if swag.IsZero(m.CustomFieldsAdd[i]) { // not required
				return nil
			}

			if err := m.CustomFieldsAdd[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_add" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields_add" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) contextValidateCustomFieldsRemove(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFieldsRemove); i++ {

		if m.CustomFieldsRemove[i] != nil {

			if swag.IsZero(m.CustomFieldsRemove[i]) { // not required
				return nil
			}

			if err := m.CustomFieldsRemove[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields_remove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields_remove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateStoryFromTemplateParams) contextValidateLabelsAdd(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelsAdd); i++ {

		if m.LabelsAdd[i] != nil {

			if swag.IsZero(m.LabelsAdd[i]) { // not required
				return nil
			}

			if err := m.LabelsAdd[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels_add" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels_add" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) contextValidateLabelsRemove(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelsRemove); i++ {

		if m.LabelsRemove[i] != nil {

			if swag.IsZero(m.LabelsRemove[i]) { // not required
				return nil
			}

			if err := m.LabelsRemove[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels_remove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels_remove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryFromTemplateParams) contextValidateStoryLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateStoryFromTemplateParams) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tasks); i++ {

		if m.Tasks[i] != nil {

			if swag.IsZero(m.Tasks[i]) { // not required
				return nil
			}

			if err := m.Tasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateStoryFromTemplateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStoryFromTemplateParams) UnmarshalBinary(b []byte) error {
	var res CreateStoryFromTemplateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
