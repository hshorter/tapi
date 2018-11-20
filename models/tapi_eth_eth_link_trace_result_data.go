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

// TapiEthEthLinkTraceResultData tapi eth eth link trace result data
// swagger:model tapi.eth.EthLinkTraceResultData
type TapiEthEthLinkTraceResultData struct {

	// G.8052: This parameter returns the results of the LT process. It contains a list of the result received from the individual LTR frames.
	//                     The result from the individual LTR frame include the Source Mac Address, the TTL, and TLV.
	ResultList []*TapiEthLinkTraceResult `json:"result-list"`
}

// Validate validates this tapi eth eth link trace result data
func (m *TapiEthEthLinkTraceResultData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthLinkTraceResultData) validateResultList(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultList) { // not required
		return nil
	}

	for i := 0; i < len(m.ResultList); i++ {
		if swag.IsZero(m.ResultList[i]) { // not required
			continue
		}

		if m.ResultList[i] != nil {
			if err := m.ResultList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthLinkTraceResultData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthLinkTraceResultData) UnmarshalBinary(b []byte) error {
	var res TapiEthEthLinkTraceResultData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
