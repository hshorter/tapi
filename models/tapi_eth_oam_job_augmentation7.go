// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamJobAugmentation7 tapi eth oam job augmentation7
// swagger:model tapi.eth.OamJobAugmentation7
type TapiEthOamJobAugmentation7 struct {

	// none
	EthProActive2wayMeasurementJob *TapiEthEthProActive2wayMeasurementJob `json:"eth-pro-active-2way-measurement-job,omitempty"`
}

// Validate validates this tapi eth oam job augmentation7
func (m *TapiEthOamJobAugmentation7) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActive2wayMeasurementJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamJobAugmentation7) validateEthProActive2wayMeasurementJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActive2wayMeasurementJob) { // not required
		return nil
	}

	if m.EthProActive2wayMeasurementJob != nil {
		if err := m.EthProActive2wayMeasurementJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-2way-measurement-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation7) UnmarshalBinary(b []byte) error {
	var res TapiEthOamJobAugmentation7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}