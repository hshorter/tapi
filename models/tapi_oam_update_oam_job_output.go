// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamUpdateOamJobOutput tapi oam update oam job output
// swagger:model tapi.oam.UpdateOamJobOutput
type TapiOamUpdateOamJobOutput struct {

	// output
	Output *TapiOamUpdateOamJobOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi oam update oam job output
func (m *TapiOamUpdateOamJobOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateOamJobOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiOamUpdateOamJobOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamJobOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamJobOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamUpdateOamJobOutputOutput tapi oam update oam job output output
// swagger:model TapiOamUpdateOamJobOutputOutput
type TapiOamUpdateOamJobOutputOutput struct {

	// none
	OamJob *TapiOamOamJob `json:"oam-job,omitempty"`
}

// Validate validates this tapi oam update oam job output output
func (m *TapiOamUpdateOamJobOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateOamJobOutputOutput) validateOamJob(formats strfmt.Registry) error {

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
func (m *TapiOamUpdateOamJobOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateOamJobOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateOamJobOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
