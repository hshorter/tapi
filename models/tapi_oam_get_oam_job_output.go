// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamGetOamJobOutput tapi oam get oam job output
// swagger:model tapi.oam.GetOamJobOutput
type TapiOamGetOamJobOutput struct {

	// output
	Output *TapiOamGetOamJobOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi oam get oam job output
func (m *TapiOamGetOamJobOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamGetOamJobOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiOamGetOamJobOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamJobOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamJobOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamGetOamJobOutputOutput tapi oam get oam job output output
// swagger:model TapiOamGetOamJobOutputOutput
type TapiOamGetOamJobOutputOutput struct {

	// none
	OamJob *TapiOamOamJob `json:"oam-job,omitempty"`
}

// Validate validates this tapi oam get oam job output output
func (m *TapiOamGetOamJobOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamGetOamJobOutputOutput) validateOamJob(formats strfmt.Registry) error {

	if swag.IsZero(m.OamJob) { // not required
		return nil
	}

	if m.OamJob != nil {
		if err := m.OamJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "oam-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetOamJobOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamJobOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamJobOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
