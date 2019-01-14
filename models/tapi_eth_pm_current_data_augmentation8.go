// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation8 tapi eth pm current data augmentation8
// swagger:model tapi.eth.PmCurrentDataAugmentation8
type TapiEthPmCurrentDataAugmentation8 struct {

	// none
	EthProActive1LmPerformanceData *TapiEthEthProActive1LmPerformanceData `json:"eth-pro-active-1-lm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation8
func (m *TapiEthPmCurrentDataAugmentation8) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActive1LmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation8) validateEthProActive1LmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActive1LmPerformanceData) { // not required
		return nil
	}

	if m.EthProActive1LmPerformanceData != nil {
		if err := m.EthProActive1LmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-1-lm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation8) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation8) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation8
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
