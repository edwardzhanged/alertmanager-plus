// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SeverityDay severity day
//
// swagger:model severityDay
type SeverityDay struct {

	// p0
	P0 int64 `json:"P0,omitempty"`

	// p1
	P1 int64 `json:"P1,omitempty"`

	// p2
	P2 int64 `json:"P2,omitempty"`

	// p3
	P3 int64 `json:"P3,omitempty"`

	// date
	Date string `json:"date,omitempty"`
}

// Validate validates this severity day
func (m *SeverityDay) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this severity day based on context it is used
func (m *SeverityDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SeverityDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeverityDay) UnmarshalBinary(b []byte) error {
	var res SeverityDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
