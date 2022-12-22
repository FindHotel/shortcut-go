// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EpicStats A group of calculated values for this Epic.
//
// swagger:model EpicStats
type EpicStats struct {

	// The average cycle time (in seconds) of completed stories in this Epic.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`

	// The average lead time (in seconds) of completed stories in this Epic.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`

	// The date of the last update of a Story in this Epic.
	// Required: true
	// Format: date-time
	LastStoryUpdate *strfmt.DateTime `json:"last_story_update"`

	// The total number of points in this Epic.
	// Required: true
	NumPoints *int64 `json:"num_points"`

	// The total number of completed points in this Epic.
	// Required: true
	NumPointsDone *int64 `json:"num_points_done"`

	// The total number of started points in this Epic.
	// Required: true
	NumPointsStarted *int64 `json:"num_points_started"`

	// The total number of unstarted points in this Epic.
	// Required: true
	NumPointsUnstarted *int64 `json:"num_points_unstarted"`

	// The total number of documents associated with this Epic.
	// Required: true
	NumRelatedDocuments *int64 `json:"num_related_documents"`

	// The total number of done Stories in this Epic.
	// Required: true
	NumStoriesDone *int64 `json:"num_stories_done"`

	// The total number of started Stories in this Epic.
	// Required: true
	NumStoriesStarted *int64 `json:"num_stories_started"`

	// The total number of Stories in this Epic.
	// Required: true
	NumStoriesTotal *int64 `json:"num_stories_total"`

	// The total number of Stories with no point estimate.
	// Required: true
	NumStoriesUnestimated *int64 `json:"num_stories_unestimated"`

	// The total number of unstarted Stories in this Epic.
	// Required: true
	NumStoriesUnstarted *int64 `json:"num_stories_unstarted"`
}

// Validate validates this epic stats
func (m *EpicStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastStoryUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsDone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsUnstarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesDone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesUnestimated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesUnstarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EpicStats) validateLastStoryUpdate(formats strfmt.Registry) error {

	if err := validate.Required("last_story_update", "body", m.LastStoryUpdate); err != nil {
		return err
	}

	if err := validate.FormatOf("last_story_update", "body", "date-time", m.LastStoryUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumPoints(formats strfmt.Registry) error {

	if err := validate.Required("num_points", "body", m.NumPoints); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumPointsDone(formats strfmt.Registry) error {

	if err := validate.Required("num_points_done", "body", m.NumPointsDone); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumPointsStarted(formats strfmt.Registry) error {

	if err := validate.Required("num_points_started", "body", m.NumPointsStarted); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumPointsUnstarted(formats strfmt.Registry) error {

	if err := validate.Required("num_points_unstarted", "body", m.NumPointsUnstarted); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumRelatedDocuments(formats strfmt.Registry) error {

	if err := validate.Required("num_related_documents", "body", m.NumRelatedDocuments); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumStoriesDone(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_done", "body", m.NumStoriesDone); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumStoriesStarted(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_started", "body", m.NumStoriesStarted); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumStoriesTotal(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_total", "body", m.NumStoriesTotal); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumStoriesUnestimated(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_unestimated", "body", m.NumStoriesUnestimated); err != nil {
		return err
	}

	return nil
}

func (m *EpicStats) validateNumStoriesUnstarted(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_unstarted", "body", m.NumStoriesUnstarted); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this epic stats based on context it is used
func (m *EpicStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EpicStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EpicStats) UnmarshalBinary(b []byte) error {
	var res EpicStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}