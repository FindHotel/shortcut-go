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

// MilestoneStats A group of calculated values for this Milestone.
//
// swagger:model MilestoneStats
type MilestoneStats struct {

	// The average cycle time (in seconds) of completed stories in this Milestone.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`

	// The average lead time (in seconds) of completed stories in this Milestone.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`

	// The number of related documents to this Milestone.
	// Required: true
	NumRelatedDocuments *int64 `json:"num_related_documents"`
}

// Validate validates this milestone stats
func (m *MilestoneStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MilestoneStats) validateNumRelatedDocuments(formats strfmt.Registry) error {

	if err := validate.Required("num_related_documents", "body", m.NumRelatedDocuments); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this milestone stats based on context it is used
func (m *MilestoneStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MilestoneStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MilestoneStats) UnmarshalBinary(b []byte) error {
	var res MilestoneStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
