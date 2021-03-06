// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPathComputationPathComputationContext tapi path computation path computation context
// swagger:model tapi.path.computation.PathComputationContext
type TapiPathComputationPathComputationContext struct {

	// none
	Path []*TapiPathComputationPath `json:"path"`

	// none
	PathCompService []*TapiPathComputationPathComputationService `json:"path-comp-service"`
}

// Validate validates this tapi path computation path computation context
func (m *TapiPathComputationPathComputationContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathCompService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPathComputationPathComputationContext) validatePath(formats strfmt.Registry) error {

	if swag.IsZero(m.Path) { // not required
		return nil
	}

	for i := 0; i < len(m.Path); i++ {
		if swag.IsZero(m.Path[i]) { // not required
			continue
		}

		if m.Path[i] != nil {
			if err := m.Path[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("path" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiPathComputationPathComputationContext) validatePathCompService(formats strfmt.Registry) error {

	if swag.IsZero(m.PathCompService) { // not required
		return nil
	}

	for i := 0; i < len(m.PathCompService); i++ {
		if swag.IsZero(m.PathCompService[i]) { // not required
			continue
		}

		if m.PathCompService[i] != nil {
			if err := m.PathCompService[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("path-comp-service" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationPathComputationContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationPathComputationContext) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationPathComputationContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
