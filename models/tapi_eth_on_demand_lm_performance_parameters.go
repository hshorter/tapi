// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthOnDemandLmPerformanceParameters tapi eth on demand lm performance parameters
// swagger:model tapi.eth.OnDemandLmPerformanceParameters
type TapiEthOnDemandLmPerformanceParameters struct {

	// This attribute contains the frame loss ratio (number of lost frames divided by the number of total frames (N_LF / N_TF)).
	//                     The accuracy of the value is for further study.
	TotalFrameLossRatio string `json:"total-frame-loss-ratio,omitempty"`

	// This attribute contains the total number of frames lost.
	TotalLostFrames int32 `json:"total-lost-frames,omitempty"`

	// This attribute contains the total number of frames transmitted.
	TotalTransmittedFrames int32 `json:"total-transmitted-frames,omitempty"`
}

// Validate validates this tapi eth on demand lm performance parameters
func (m *TapiEthOnDemandLmPerformanceParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOnDemandLmPerformanceParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOnDemandLmPerformanceParameters) UnmarshalBinary(b []byte) error {
	var res TapiEthOnDemandLmPerformanceParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
