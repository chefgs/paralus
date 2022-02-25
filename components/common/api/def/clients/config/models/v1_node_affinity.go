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
)

// V1NodeAffinity Node affinity is a group of node affinity scheduling rules.
//
// swagger:model v1NodeAffinity
type V1NodeAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy
	// the affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node matches the corresponding matchExpressions; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	PreferredDuringSchedulingIgnoredDuringExecution []*V1PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to an update), the system
	// may or may not try to eventually evict the pod from its node.
	// +optional
	RequiredDuringSchedulingIgnoredDuringExecution *V1NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// Validate validates this v1 node affinity
func (m *V1NodeAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NodeAffinity) validatePreferredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NodeAffinity) validateRequiredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	if m.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if err := m.RequiredDuringSchedulingIgnoredDuringExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredDuringSchedulingIgnoredDuringExecution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredDuringSchedulingIgnoredDuringExecution")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 node affinity based on the context it is used
func (m *V1NodeAffinity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NodeAffinity) contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NodeAffinity) contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if err := m.RequiredDuringSchedulingIgnoredDuringExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredDuringSchedulingIgnoredDuringExecution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredDuringSchedulingIgnoredDuringExecution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NodeAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NodeAffinity) UnmarshalBinary(b []byte) error {
	var res V1NodeAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}