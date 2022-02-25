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

// Commonv3NodeCondition NodeCondition
//
// Node Condition
//
// swagger:model commonv3NodeCondition
type Commonv3NodeCondition struct {

	// Last time we got an update on a given condition.
	LastHeartbeatTime *Commonv3Time `json:"lastHeartbeatTime,omitempty"`

	// Last time the condition transit from one status to another.
	LastTransitionTime *Commonv3Time `json:"lastTransitionTime,omitempty"`

	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// (brief) reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status,omitempty"`

	// Type
	//
	// Type of node condition
	Type string `json:"type,omitempty"`
}

// Validate validates this commonv3 node condition
func (m *Commonv3NodeCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastHeartbeatTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commonv3NodeCondition) validateLastHeartbeatTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastHeartbeatTime) { // not required
		return nil
	}

	if m.LastHeartbeatTime != nil {
		if err := m.LastHeartbeatTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastHeartbeatTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastHeartbeatTime")
			}
			return err
		}
	}

	return nil
}

func (m *Commonv3NodeCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this commonv3 node condition based on the context it is used
func (m *Commonv3NodeCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastHeartbeatTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commonv3NodeCondition) contextValidateLastHeartbeatTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastHeartbeatTime != nil {
		if err := m.LastHeartbeatTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastHeartbeatTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastHeartbeatTime")
			}
			return err
		}
	}

	return nil
}

func (m *Commonv3NodeCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Commonv3NodeCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commonv3NodeCondition) UnmarshalBinary(b []byte) error {
	var res Commonv3NodeCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}