// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation3 tapi eth pm current data augmentation3
// swagger:model tapi.eth.PmCurrentDataAugmentation3
type TapiEthPmCurrentDataAugmentation3 struct {

	// none
	EthOnDemandDmPerformanceData *TapiEthEthOnDemandDmPerformanceData `json:"eth-on-demand-dm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation3
func (m *TapiEthPmCurrentDataAugmentation3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandDmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation3) validateEthOnDemandDmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandDmPerformanceData) { // not required
		return nil
	}

	if m.EthOnDemandDmPerformanceData != nil {
		if err := m.EthOnDemandDmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-dm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation3) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
