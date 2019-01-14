// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation4 tapi eth pm current data augmentation4
// swagger:model tapi.eth.PmCurrentDataAugmentation4
type TapiEthPmCurrentDataAugmentation4 struct {

	// none
	EthOnDemandLmPerformanceData *TapiEthEthOnDemandLmPerformanceData `json:"eth-on-demand-lm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation4
func (m *TapiEthPmCurrentDataAugmentation4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandLmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation4) validateEthOnDemandLmPerformanceData(formats strfmt.Registry) error {

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
func (m *TapiEthPmCurrentDataAugmentation4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation4) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
