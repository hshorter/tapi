// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiTopologyTopologyRef tapi topology topology ref
// swagger:model tapi.topology.TopologyRef
type TapiTopologyTopologyRef struct {

	// none
	TopologyUUID string `json:"topology-uuid,omitempty"`
}

// Validate validates this tapi topology topology ref
func (m *TapiTopologyTopologyRef) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyTopologyRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyTopologyRef) UnmarshalBinary(b []byte) error {
	var res TapiTopologyTopologyRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
