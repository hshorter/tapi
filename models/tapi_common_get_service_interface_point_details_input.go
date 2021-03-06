// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonGetServiceInterfacePointDetailsInput tapi common get service interface point details input
// swagger:model tapi.common.GetServiceInterfacePointDetailsInput
type TapiCommonGetServiceInterfacePointDetailsInput struct {

	// input
	Input *TapiCommonGetServiceInterfacePointDetailsInputInput `json:"input,omitempty"`
}

// Validate validates this tapi common get service interface point details input
func (m *TapiCommonGetServiceInterfacePointDetailsInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGetServiceInterfacePointDetailsInput) validateInput(formats strfmt.Registry) error {

	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsInput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointDetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiCommonGetServiceInterfacePointDetailsInputInput tapi common get service interface point details input input
// swagger:model TapiCommonGetServiceInterfacePointDetailsInputInput
type TapiCommonGetServiceInterfacePointDetailsInputInput struct {

	// none
	SipIDOrName string `json:"sip-id-or-name,omitempty"`
}

// Validate validates this tapi common get service interface point details input input
func (m *TapiCommonGetServiceInterfacePointDetailsInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsInputInput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointDetailsInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
