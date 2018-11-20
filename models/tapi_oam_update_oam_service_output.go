// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamUpdateOamServiceOutput tapi oam update oam service output
// swagger:model tapi.oam.UpdateOamServiceOutput
type TapiOamUpdateOamServiceOutput struct {

	// output
	Output *TapiOamUpdateOamServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi oam update oam service output
func (m *TapiOamUpdateOamServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateOamServiceOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiOamUpdateOamServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamUpdateOamServiceOutputOutput tapi oam update oam service output output
// swagger:model TapiOamUpdateOamServiceOutputOutput
type TapiOamUpdateOamServiceOutputOutput struct {

	// none
	Service *TapiOamOamService `json:"service,omitempty"`
}

// Validate validates this tapi oam update oam service output output
func (m *TapiOamUpdateOamServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateOamServiceOutputOutput) validateService(formats strfmt.Registry) error {

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
func (m *TapiOamUpdateOamServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
