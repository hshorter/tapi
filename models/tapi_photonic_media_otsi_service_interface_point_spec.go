// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaOtsiServiceInterfacePointSpec tapi photonic media otsi service interface point spec
// swagger:model tapi.photonic.media.OtsiServiceInterfacePointSpec
type TapiPhotonicMediaOtsiServiceInterfacePointSpec struct {

	// none
	OtsiCapability *TapiPhotonicMediaOtsiCapabilityPac `json:"otsi-capability,omitempty"`
}

// Validate validates this tapi photonic media otsi service interface point spec
func (m *TapiPhotonicMediaOtsiServiceInterfacePointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtsiCapability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaOtsiServiceInterfacePointSpec) validateOtsiCapability(formats strfmt.Registry) error {

	if swag.IsZero(m.OtsiCapability) { // not required
		return nil
	}

	if m.OtsiCapability != nil {
		if err := m.OtsiCapability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otsi-capability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaOtsiServiceInterfacePointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaOtsiServiceInterfacePointSpec) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaOtsiServiceInterfacePointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
