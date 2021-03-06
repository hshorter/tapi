// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamGetOamJobInput tapi oam get oam job input
// swagger:model tapi.oam.GetOamJobInput
type TapiOamGetOamJobInput struct {

	// input
	Input *TapiOamGetOamJobInputInput `json:"input,omitempty"`
}

// Validate validates this tapi oam get oam job input
func (m *TapiOamGetOamJobInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamGetOamJobInput) validateInput(formats strfmt.Registry) error {

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
func (m *TapiOamGetOamJobInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamJobInput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamJobInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamGetOamJobInputInput tapi oam get oam job input input
// swagger:model TapiOamGetOamJobInputInput
type TapiOamGetOamJobInputInput struct {

	// none
	JobID string `json:"job-id,omitempty"`
}

// Validate validates this tapi oam get oam job input input
func (m *TapiOamGetOamJobInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetOamJobInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamJobInputInput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamJobInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
