// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation1 tapi eth pm current data augmentation1
// swagger:model tapi.eth.PmCurrentDataAugmentation1
type TapiEthPmCurrentDataAugmentation1 struct {

	// none
	EthOnDemandLmPerformanceData *TapiEthEthOnDemandLmPerformanceData `json:"eth-on-demand-lm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation1
func (m *TapiEthPmCurrentDataAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandLmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation1) validateEthOnDemandLmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandLmPerformanceData) { // not required
		return nil
	}

	if m.EthOnDemandLmPerformanceData != nil {
		if err := m.EthOnDemandLmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-lm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation1) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
