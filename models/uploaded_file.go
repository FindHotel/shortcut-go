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

// UploadedFile An UploadedFile is any document uploaded to your Shortcut Workspace. Files attached from a third-party service are different: see the Linked Files endpoint.
//
// swagger:model UploadedFile
type UploadedFile struct {

	// Free form string corresponding to a text or image file.
	// Required: true
	ContentType *string `json:"content_type"`

	// The time/date that the file was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The description of the file.
	// Required: true
	Description *string `json:"description"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// This field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here.
	// Required: true
	ExternalID *string `json:"external_id"`

	// The name assigned to the file in Shortcut upon upload.
	// Required: true
	Filename *string `json:"filename"`

	// The unique IDs of the Groups who are mentioned in the file description.
	// Required: true
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The unique ID for the file.
	// Required: true
	ID *int64 `json:"id"`

	// The unique IDs of the Members who are mentioned in the file description.
	// Required: true
	MemberMentionIds []strfmt.UUID `json:"member_mention_ids"`

	// Deprecated: use member_mention_ids.
	// Required: true
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// The optional User-specified name of the file.
	// Required: true
	Name *string `json:"name"`

	// The size of the file.
	// Required: true
	Size *int64 `json:"size"`

	// The unique IDs of the Stories associated with this file.
	// Required: true
	StoryIds []int64 `json:"story_ids"`

	// The url where the thumbnail of the file can be found in Shortcut.
	// Required: true
	ThumbnailURL *string `json:"thumbnail_url"`

	// The time/date that the file was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The unique ID of the Member who uploaded the file.
	// Required: true
	// Format: uuid
	UploaderID *strfmt.UUID `json:"uploader_id"`

	// The URL for the file.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this uploaded file
func (m *UploadedFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnailURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploaderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadedFile) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateGroupMentionIds(formats strfmt.Registry) error {

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

func (m *UploadedFile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateMemberMentionIds(formats strfmt.Registry) error {

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

func (m *UploadedFile) validateMentionIds(formats strfmt.Registry) error {

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

func (m *UploadedFile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateStoryIds(formats strfmt.Registry) error {

	if err := validate.Required("story_ids", "body", m.StoryIds); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateThumbnailURL(formats strfmt.Registry) error {

	if err := validate.Required("thumbnail_url", "body", m.ThumbnailURL); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateUploaderID(formats strfmt.Registry) error {

	if err := validate.Required("uploader_id", "body", m.UploaderID); err != nil {
		return err
	}

	if err := validate.FormatOf("uploader_id", "body", "uuid", m.UploaderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UploadedFile) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this uploaded file based on context it is used
func (m *UploadedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadedFile) UnmarshalBinary(b []byte) error {
	var res UploadedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
