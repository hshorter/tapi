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

// TapiOamCreateOamJobInput tapi oam create oam job input
// swagger:model tapi.oam.CreateOamJobInput
type TapiOamCreateOamJobInput struct {

	// input
	Input *TapiOamCreateOamJobInputInput `json:"input,omitempty"`
}

// Validate validates this tapi oam create oam job input
func (m *TapiOamCreateOamJobInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamJobInput) validateInput(formats strfmt.Registry) error {

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
func (m *TapiOamCreateOamJobInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamJobInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamJobInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamCreateOamJobInputInput tapi oam create oam job input input
// swagger:model TapiOamCreateOamJobInputInput
type TapiOamCreateOamJobInputInput struct {

	// none
	OamJobType string `json:"oam-job-type,omitempty"`

	// none
	OamProfile *TapiOamOamProfile `json:"oam-profile,omitempty"`

	// none
	OamServiceEndPoint []*TapiOamOamServiceEndPoint `json:"oam-service-end-point"`

	// none
	Schedule string `json:"schedule,omitempty"`

	// none
	State string `json:"state,omitempty"`
}

// Validate validates this tapi oam create oam job input input
func (m *TapiOamCreateOamJobInputInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamServiceEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamJobInputInput) validateOamProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.OamProfile) { // not required
		return nil
	}

	if m.OamProfile != nil {
		if err := m.OamProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "oam-profile")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamCreateOamJobInputInput) validateOamServiceEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.OamServiceEndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.OamServiceEndPoint); i++ {
		if swag.IsZero(m.OamServiceEndPoint[i]) { // not required
			continue
		}

		if m.OamServiceEndPoint[i] != nil {
			if err := m.OamServiceEndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + "oam-service-end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateOamJobInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamJobInputInput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamJobInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}