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

// LabelStats A group of calculated values for this Label. This is not included if the slim? flag is set to true for the List Labels endpoint.
//
// swagger:model LabelStats
type LabelStats struct {

	// The total number of Epics with this Label.
	// Required: true
	NumEpics *int64 `json:"num_epics"`

	// The number of completed Epics associated with this Label.
	// Required: true
	NumEpicsCompleted *int64 `json:"num_epics_completed"`

	// The number of in progress epics associated with this label.
	// Required: true
	NumEpicsInProgress *int64 `json:"num_epics_in_progress"`

	// The total number of Epics associated with this Label.
	// Required: true
	NumEpicsTotal *int64 `json:"num_epics_total"`

	// The number of unstarted epics associated with this label.
	// Required: true
	NumEpicsUnstarted *int64 `json:"num_epics_unstarted"`

	// The total number of backlog points with this Label.
	// Required: true
	NumPointsBacklog *int64 `json:"num_points_backlog"`

	// The total number of completed points with this Label.
	// Required: true
	NumPointsCompleted *int64 `json:"num_points_completed"`

	// The total number of in-progress points with this Label.
	// Required: true
	NumPointsInProgress *int64 `json:"num_points_in_progress"`

	// The total number of points with this Label.
	// Required: true
	NumPointsTotal *int64 `json:"num_points_total"`

	// The total number of unstarted points with this Label.
	// Required: true
	NumPointsUnstarted *int64 `json:"num_points_unstarted"`

	// The total number of Documents associated this Label.
	// Required: true
	NumRelatedDocuments *int64 `json:"num_related_documents"`

	// The total number of stories backlog Stories with this Label.
	// Required: true
	NumStoriesBacklog *int64 `json:"num_stories_backlog"`

	// The total number of completed Stories with this Label.
	// Required: true
	NumStoriesCompleted *int64 `json:"num_stories_completed"`

	// The total number of in-progress Stories with this Label.
	// Required: true
	NumStoriesInProgress *int64 `json:"num_stories_in_progress"`

	// The total number of Stories with this Label.
	// Required: true
	NumStoriesTotal *int64 `json:"num_stories_total"`

	// The total number of Stories with no point estimate with this Label.
	// Required: true
	NumStoriesUnestimated *int64 `json:"num_stories_unestimated"`

	// The total number of stories unstarted Stories with this Label.
	// Required: true
	NumStoriesUnstarted *int64 `json:"num_stories_unstarted"`
}

// Validate validates this label stats
func (m *LabelStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumEpics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumEpicsCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumEpicsInProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumEpicsTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumEpicsUnstarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsBacklog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsInProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPointsUnstarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesBacklog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumStoriesInProgress(formats); err != nil {
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

func (m *LabelStats) validateNumEpics(formats strfmt.Registry) error {

	if err := validate.Required("num_epics", "body", m.NumEpics); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumEpicsCompleted(formats strfmt.Registry) error {

	if err := validate.Required("num_epics_completed", "body", m.NumEpicsCompleted); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumEpicsInProgress(formats strfmt.Registry) error {

	if err := validate.Required("num_epics_in_progress", "body", m.NumEpicsInProgress); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumEpicsTotal(formats strfmt.Registry) error {

	if err := validate.Required("num_epics_total", "body", m.NumEpicsTotal); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumEpicsUnstarted(formats strfmt.Registry) error {

	if err := validate.Required("num_epics_unstarted", "body", m.NumEpicsUnstarted); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumPointsBacklog(formats strfmt.Registry) error {

	if err := validate.Required("num_points_backlog", "body", m.NumPointsBacklog); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumPointsCompleted(formats strfmt.Registry) error {

	if err := validate.Required("num_points_completed", "body", m.NumPointsCompleted); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumPointsInProgress(formats strfmt.Registry) error {

	if err := validate.Required("num_points_in_progress", "body", m.NumPointsInProgress); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumPointsTotal(formats strfmt.Registry) error {

	if err := validate.Required("num_points_total", "body", m.NumPointsTotal); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumPointsUnstarted(formats strfmt.Registry) error {

	if err := validate.Required("num_points_unstarted", "body", m.NumPointsUnstarted); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumRelatedDocuments(formats strfmt.Registry) error {

	if err := validate.Required("num_related_documents", "body", m.NumRelatedDocuments); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesBacklog(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_backlog", "body", m.NumStoriesBacklog); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesCompleted(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_completed", "body", m.NumStoriesCompleted); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesInProgress(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_in_progress", "body", m.NumStoriesInProgress); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesTotal(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_total", "body", m.NumStoriesTotal); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesUnestimated(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_unestimated", "body", m.NumStoriesUnestimated); err != nil {
		return err
	}

	return nil
}

func (m *LabelStats) validateNumStoriesUnstarted(formats strfmt.Registry) error {

	if err := validate.Required("num_stories_unstarted", "body", m.NumStoriesUnstarted); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this label stats based on context it is used
func (m *LabelStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelStats) UnmarshalBinary(b []byte) error {
	var res LabelStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
