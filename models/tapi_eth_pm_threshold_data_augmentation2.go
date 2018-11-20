// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmThresholdDataAugmentation2 tapi eth pm threshold data augmentation2
// swagger:model tapi.eth.PmThresholdDataAugmentation2
type TapiEthPmThresholdDataAugmentation2 struct {

	// none
	Eth1DmThresholdData *TapiEthEth1DmThresholdData `json:"eth-1-dm-threshold-data,omitempty"`
}

// Validate validates this tapi eth pm threshold data augmentation2
func (m *TapiEthPmThresholdDataAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEth1DmThresholdData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmThresholdDataAugmentation2) validateEth1DmThresholdData(formats strfmt.Registry) error {

	if swag.IsZero(m.Eth1DmThresholdData) { // not required
		return nil
	}

	if m.Eth1DmThresholdData != nil {
		if err := m.Eth1DmThresholdData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-1-dm-threshold-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmThresholdDataAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmThresholdDataAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiEthPmThresholdDataAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
