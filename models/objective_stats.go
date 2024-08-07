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

// ObjectiveStats A group of calculated values for this Objective.
//
// swagger:model ObjectiveStats
type ObjectiveStats struct {

	// The average cycle time (in seconds) of completed stories in this Objective.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`

	// The average lead time (in seconds) of completed stories in this Objective.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`

	// The number of related documents to this Objective.
	// Required: true
	NumRelatedDocuments *int64 `json:"num_related_documents"`
}

// Validate validates this objective stats
func (m *ObjectiveStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectiveStats) validateNumRelatedDocuments(formats strfmt.Registry) error {

	if err := validate.Required("num_related_documents", "body", m.NumRelatedDocuments); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this objective stats based on context it is used
func (m *ObjectiveStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectiveStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectiveStats) UnmarshalBinary(b []byte) error {
	var res ObjectiveStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
