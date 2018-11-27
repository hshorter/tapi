// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamJobAugmentation5 tapi eth oam job augmentation5
// swagger:model tapi.eth.OamJobAugmentation5
type TapiEthOamJobAugmentation5 struct {

	// none
	EthLinkTraceJob *TapiEthEthLinkTraceJob `json:"eth-link-trace-job,omitempty"`
}

// Validate validates this tapi eth oam job augmentation5
func (m *TapiEthOamJobAugmentation5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthLinkTraceJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamJobAugmentation5) validateEthLinkTraceJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthLinkTraceJob) { // not required
		return nil
	}

	if m.EthLinkTraceJob != nil {
		if err := m.EthLinkTraceJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-link-trace-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation5) UnmarshalBinary(b []byte) error {
	var res TapiEthOamJobAugmentation5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
