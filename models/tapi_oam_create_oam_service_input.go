// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamCreateOamServiceInput tapi oam create oam service input
// swagger:model tapi.oam.CreateOamServiceInput
type TapiOamCreateOamServiceInput struct {

	// input
	Input *TapiOamCreateOamServiceInputInput `json:"input,omitempty"`
}

// Validate validates this tapi oam create oam service input
func (m *TapiOamCreateOamServiceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamServiceInput) validateInput(formats strfmt.Registry) error {

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
func (m *TapiOamCreateOamServiceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamCreateOamServiceInputInput tapi oam create oam service input input
// swagger:model TapiOamCreateOamServiceInputInput
type TapiOamCreateOamServiceInputInput struct {

	// none
	EndPoint []*TapiOamOamServiceEndPoint `json:"end-point"`

	// none
	OamConstraint *TapiOamOamConstraint `json:"oam-constraint,omitempty"`

	// none
	State string `json:"state,omitempty"`
}

// Validate validates this tapi oam create oam service input input
func (m *TapiOamCreateOamServiceInputInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamServiceInputInput) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPoint); i++ {
		if swag.IsZero(m.EndPoint[i]) { // not required
			continue
		}

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + "end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamCreateOamServiceInputInput) validateOamConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.OamConstraint) { // not required
		return nil
	}

	if m.OamConstraint != nil {
		if err := m.OamConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "oam-constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateOamServiceInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceInputInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
