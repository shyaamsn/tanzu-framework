// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VSphereVirtualMachine v sphere virtual machine
// swagger:model VSphereVirtualMachine
type VSphereVirtualMachine struct {

	// is template
	// Required: true
	IsTemplate *bool `json:"isTemplate"`

	// k8s version
	K8sVersion string `json:"k8sVersion,omitempty"`

	// moid
	Moid string `json:"moid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// os info
	OsInfo *OSInfo `json:"osInfo,omitempty"`
}

// Validate validates this v sphere virtual machine
func (m *VSphereVirtualMachine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VSphereVirtualMachine) validateIsTemplate(formats strfmt.Registry) error {

	if err := validate.Required("isTemplate", "body", m.IsTemplate); err != nil {
		return err
	}

	return nil
}

func (m *VSphereVirtualMachine) validateOsInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.OsInfo) { // not required
		return nil
	}

	if m.OsInfo != nil {
		if err := m.OsInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VSphereVirtualMachine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSphereVirtualMachine) UnmarshalBinary(b []byte) error {
	var res VSphereVirtualMachine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}