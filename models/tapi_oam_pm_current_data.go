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

// TapiOamPmCurrentData tapi oam pm current data
// swagger:model tapi.oam.PmCurrentData
type TapiOamPmCurrentData struct {
	TapiCommonLocalClass

	// none
	ElapsedTime *TapiCommonTimeInterval `json:"elapsed-time,omitempty"`

	// none
	GranularityPeriod *TapiCommonTimePeriod `json:"granularity-period,omitempty"`

	// in case of 24hr Current Data, at least 1 History Data.
	//                     In case of 15min Current Data, at least 16 History Data.
	//                     In case of <15min, the number of History Data shall be able to cover a span of 4 hours.
	PmHistoryData []*TapiOamPmHistoryData `json:"pm-history-data"`

	// This attribute is used to indicate that the performance data for the current period may not be reliable. Some reasons for this to occur are:
	//                     – Suspect data were detected by the actual resource doing data collection.
	//                     – Transition of the administrativeState attribute to/from the 'lock' state.
	//                     – Transition of the operationalState to/from the 'disabled' state.
	//                     – Scheduler setting that inhibits the collection function.
	//                     – The performance counters were reset during the interval.
	//                     – The currentData (or subclass) object instance was created during the monitoring period.
	SuspectIntervalFlag *bool `json:"suspect-interval-flag,omitempty"`

	// This attribute indicates the start of the current monitoring interval.
	//                     The value is bound to the quarter of an hour in case of a 15 minute interval and bound to the hour in case of a 24 hour interval.
	Timestamp string `json:"timestamp,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamPmCurrentData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonLocalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonLocalClass = aO0

	// AO1
	var dataAO1 struct {
		ElapsedTime *TapiCommonTimeInterval `json:"elapsed-time,omitempty"`

		GranularityPeriod *TapiCommonTimePeriod `json:"granularity-period,omitempty"`

		PmHistoryData []*TapiOamPmHistoryData `json:"pm-history-data"`

		SuspectIntervalFlag *bool `json:"suspect-interval-flag,omitempty"`

		Timestamp string `json:"timestamp,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ElapsedTime = dataAO1.ElapsedTime

	m.GranularityPeriod = dataAO1.GranularityPeriod

	m.PmHistoryData = dataAO1.PmHistoryData

	m.SuspectIntervalFlag = dataAO1.SuspectIntervalFlag

	m.Timestamp = dataAO1.Timestamp

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamPmCurrentData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonLocalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ElapsedTime *TapiCommonTimeInterval `json:"elapsed-time,omitempty"`

		GranularityPeriod *TapiCommonTimePeriod `json:"granularity-period,omitempty"`

		PmHistoryData []*TapiOamPmHistoryData `json:"pm-history-data"`

		SuspectIntervalFlag *bool `json:"suspect-interval-flag,omitempty"`

		Timestamp string `json:"timestamp,omitempty"`
	}

	dataAO1.ElapsedTime = m.ElapsedTime

	dataAO1.GranularityPeriod = m.GranularityPeriod

	dataAO1.PmHistoryData = m.PmHistoryData

	dataAO1.SuspectIntervalFlag = m.SuspectIntervalFlag

	dataAO1.Timestamp = m.Timestamp

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam pm current data
func (m *TapiOamPmCurrentData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonLocalClass
	if err := m.TapiCommonLocalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElapsedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGranularityPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmHistoryData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamPmCurrentData) validateElapsedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ElapsedTime) { // not required
		return nil
	}

	if m.ElapsedTime != nil {
		if err := m.ElapsedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elapsed-time")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamPmCurrentData) validateGranularityPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.GranularityPeriod) { // not required
		return nil
	}

	if m.GranularityPeriod != nil {
		if err := m.GranularityPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("granularity-period")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamPmCurrentData) validatePmHistoryData(formats strfmt.Registry) error {

	if swag.IsZero(m.PmHistoryData) { // not required
		return nil
	}

	for i := 0; i < len(m.PmHistoryData); i++ {
		if swag.IsZero(m.PmHistoryData[i]) { // not required
			continue
		}

		if m.PmHistoryData[i] != nil {
			if err := m.PmHistoryData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pm-history-data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamPmCurrentData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamPmCurrentData) UnmarshalBinary(b []byte) error {
	var res TapiOamPmCurrentData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
