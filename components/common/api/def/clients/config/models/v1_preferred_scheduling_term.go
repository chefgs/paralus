// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PreferredSchedulingTerm An empty preferred scheduling term matches all objects with implicit weight 0
// (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
//
// swagger:model v1PreferredSchedulingTerm
type V1PreferredSchedulingTerm struct {

	// A node selector term, associated with the corresponding weight.
	Preference *V1NodeSelectorTerm `json:"preference,omitempty"`

	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight,omitempty"`
}

// Validate validates this v1 preferred scheduling term
func (m *V1PreferredSchedulingTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PreferredSchedulingTerm) validatePreference(formats strfmt.Registry) error {
	if swag.IsZero(m.Preference) { // not required
		return nil
	}

	if m.Preference != nil {
		if err := m.Preference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 preferred scheduling term based on the context it is used
func (m *V1PreferredSchedulingTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PreferredSchedulingTerm) contextValidatePreference(ctx context.Context, formats strfmt.Registry) error {

	if m.Preference != nil {
		if err := m.Preference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PreferredSchedulingTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PreferredSchedulingTerm) UnmarshalBinary(b []byte) error {
	var res V1PreferredSchedulingTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}