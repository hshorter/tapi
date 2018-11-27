// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation9 tapi eth pm current data augmentation9
// swagger:model tapi.eth.PmCurrentDataAugmentation9
type TapiEthPmCurrentDataAugmentation9 struct {

	// none
	EthProActiveLmPerformanceData *TapiEthEthProActiveLmPerformanceData `json:"eth-pro-active-lm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation9
func (m *TapiEthPmCurrentDataAugmentation9) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActiveLmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation9) validateEthProActiveLmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActiveLmPerformanceData) { // not required
		return nil
	}

	if m.EthProActiveLmPerformanceData != nil {
		if err := m.EthProActiveLmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-lm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation9) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation9) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation9
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
