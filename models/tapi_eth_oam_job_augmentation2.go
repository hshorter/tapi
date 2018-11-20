// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamJobAugmentation2 tapi eth oam job augmentation2
// swagger:model tapi.eth.OamJobAugmentation2
type TapiEthOamJobAugmentation2 struct {

	// none
	EthOnDemand1wayMeasurementJob *TapiEthEthOnDemand1wayMeasurementJob `json:"eth-on-demand-1way-measurement-job,omitempty"`
}

// Validate validates this tapi eth oam job augmentation2
func (m *TapiEthOamJobAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemand1wayMeasurementJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamJobAugmentation2) validateEthOnDemand1wayMeasurementJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemand1wayMeasurementJob) { // not required
		return nil
	}

	if m.EthOnDemand1wayMeasurementJob != nil {
		if err := m.EthOnDemand1wayMeasurementJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-1way-measurement-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiEthOamJobAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
