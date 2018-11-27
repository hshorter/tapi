// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamJobAugmentation3 tapi eth oam job augmentation3
// swagger:model tapi.eth.OamJobAugmentation3
type TapiEthOamJobAugmentation3 struct {

	// none
	EthTestJob *TapiEthEthTestJob `json:"eth-test-job,omitempty"`
}

// Validate validates this tapi eth oam job augmentation3
func (m *TapiEthOamJobAugmentation3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthTestJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamJobAugmentation3) validateEthTestJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthTestJob) { // not required
		return nil
	}

	if m.EthTestJob != nil {
		if err := m.EthTestJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-test-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation3) UnmarshalBinary(b []byte) error {
	var res TapiEthOamJobAugmentation3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
