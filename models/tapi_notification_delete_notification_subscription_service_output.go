// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNotificationDeleteNotificationSubscriptionServiceOutput tapi notification delete notification subscription service output
// swagger:model tapi.notification.DeleteNotificationSubscriptionServiceOutput
type TapiNotificationDeleteNotificationSubscriptionServiceOutput struct {

	// output
	Output *TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi notification delete notification subscription service output
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutput) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationDeleteNotificationSubscriptionServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput tapi notification delete notification subscription service output output
// swagger:model TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput
type TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput struct {

	// none
	SubscriptionService *TapiNotificationNotificationSubscriptionService `json:"subscription-service,omitempty"`
}

// Validate validates this tapi notification delete notification subscription service output output
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput) validateSubscriptionService(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionService) { // not required
		return nil
	}

	if m.SubscriptionService != nil {
		if err := m.SubscriptionService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "subscription-service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationDeleteNotificationSubscriptionServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
