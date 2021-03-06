// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkDeleteVirtualNetworkServiceOutput tapi virtual network delete virtual network service output
// swagger:model tapi.virtual.network.DeleteVirtualNetworkServiceOutput
type TapiVirtualNetworkDeleteVirtualNetworkServiceOutput struct {

	// output
	Output *TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi virtual network delete virtual network service output
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutput) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkDeleteVirtualNetworkServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput tapi virtual network delete virtual network service output output
// swagger:model TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput
type TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput struct {

	// none
	Service *TapiVirtualNetworkVirtualNetworkService `json:"service,omitempty"`
}

// Validate validates this tapi virtual network delete virtual network service output output
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkDeleteVirtualNetworkServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
