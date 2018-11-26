// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthMipAugmentation2 tapi eth mip augmentation2
// swagger:model tapi.eth.MipAugmentation2
type TapiEthMipAugmentation2 struct {

	// none
	EthMipSpec *TapiEthEthMipSpec `json:"eth-mip-spec,omitempty"`
}

// Validate validates this tapi eth mip augmentation2
func (m *TapiEthMipAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthMipSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthMipAugmentation2) validateEthMipSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.EthMipSpec) { // not required
		return nil
	}

	if m.EthMipSpec != nil {
		if err := m.EthMipSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-mip-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthMipAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthMipAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiEthMipAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
