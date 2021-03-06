// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamUpdateOamServiceEndPointInput tapi oam update oam service end point input
// swagger:model tapi.oam.UpdateOamServiceEndPointInput
type TapiOamUpdateOamServiceEndPointInput struct {

	// input
	Input *TapiOamUpdateOamServiceEndPointInputInput `json:"input,omitempty"`
}

// Validate validates this tapi oam update oam service end point input
func (m *TapiOamUpdateOamServiceEndPointInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateOamServiceEndPointInput) validateInput(formats strfmt.Registry) error {

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
func (m *TapiOamUpdateOamServiceEndPointInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamServiceEndPointInput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamServiceEndPointInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamUpdateOamServiceEndPointInputInput tapi oam update oam service end point input input
// swagger:model TapiOamUpdateOamServiceEndPointInputInput
type TapiOamUpdateOamServiceEndPointInputInput struct {

	// none
	OSepID string `json:"o-sep-id,omitempty"`

	// none
	ServiceID string `json:"service-id,omitempty"`

	// none
	State string `json:"state,omitempty"`
}

// Validate validates this tapi oam update oam service end point input input
func (m *TapiOamUpdateOamServiceEndPointInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamUpdateOamServiceEndPointInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamServiceEndPointInputInput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamServiceEndPointInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
