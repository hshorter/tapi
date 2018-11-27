// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmThresholdDataAugmentation4 tapi eth pm threshold data augmentation4
// swagger:model tapi.eth.PmThresholdDataAugmentation4
type TapiEthPmThresholdDataAugmentation4 struct {

	// none
	Eth1DmThresholdData *TapiEthEth1DmThresholdData `json:"eth-1-dm-threshold-data,omitempty"`
}

// Validate validates this tapi eth pm threshold data augmentation4
func (m *TapiEthPmThresholdDataAugmentation4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEth1DmThresholdData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmThresholdDataAugmentation4) validateEth1DmThresholdData(formats strfmt.Registry) error {

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
func (m *TapiEthPmThresholdDataAugmentation4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmThresholdDataAugmentation4) UnmarshalBinary(b []byte) error {
	var res TapiEthPmThresholdDataAugmentation4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
