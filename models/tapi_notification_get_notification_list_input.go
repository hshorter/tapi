// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNotificationGetNotificationListInput tapi notification get notification list input
// swagger:model tapi.notification.GetNotificationListInput
type TapiNotificationGetNotificationListInput struct {

	// input
	Input *TapiNotificationGetNotificationListInputInput `json:"input,omitempty"`
}

// Validate validates this tapi notification get notification list input
func (m *TapiNotificationGetNotificationListInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiNotificationGetNotificationListInput) validateInput(formats strfmt.Registry) error {

	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationGetNotificationListInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationGetNotificationListInput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationGetNotificationListInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiNotificationGetNotificationListInputInput tapi notification get notification list input input
// swagger:model TapiNotificationGetNotificationListInputInput
type TapiNotificationGetNotificationListInputInput struct {

	// none
	SubscriptionIDOrName string `json:"subscription-id-or-name,omitempty"`

	// none
	TimePeriod string `json:"time-period,omitempty"`
}

// Validate validates this tapi notification get notification list input input
func (m *TapiNotificationGetNotificationListInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationGetNotificationListInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationGetNotificationListInputInput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationGetNotificationListInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
