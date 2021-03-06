// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamCreateOamServiceEndPointInput tapi oam create oam service end point input
// swagger:model tapi.oam.CreateOamServiceEndPointInput
type TapiOamCreateOamServiceEndPointInput struct {

	// input
	Input *TapiOamCreateOamServiceEndPointInputInput `json:"input,omitempty"`
}

// Validate validates this tapi oam create oam service end point input
func (m *TapiOamCreateOamServiceEndPointInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamServiceEndPointInput) validateInput(formats strfmt.Registry) error {

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
func (m *TapiOamCreateOamServiceEndPointInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceEndPointInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamCreateOamServiceEndPointInputInput tapi oam create oam service end point input input
// swagger:model TapiOamCreateOamServiceEndPointInputInput
type TapiOamCreateOamServiceEndPointInputInput struct {

	// none
	CSepID string `json:"c-sep-id,omitempty"`

	// none
	Direction string `json:"direction,omitempty"`

	// none
	Layer string `json:"layer,omitempty"`

	// none
	MepIdentifier string `json:"mep-identifier,omitempty"`

	// none
	PeerMepIdentifier []string `json:"peer-mep-identifier"`

	// none
	ServiceID string `json:"service-id,omitempty"`

	// none
	SipID string `json:"sip-id,omitempty"`

	// none
	State string `json:"state,omitempty"`
}

// Validate validates this tapi oam create oam service end point input input
func (m *TapiOamCreateOamServiceEndPointInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointInputInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceEndPointInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
