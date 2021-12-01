/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
	"fmt"
)

// V3SupportPackageTimezones * `ET` - Eastern Time * `GMT` - Greenwich Mean Time * `IST` - India Standard Time * `PT` - Pacific Time
type V3SupportPackageTimezones string

// List of v3SupportPackageTimezones
const (
	V3_ET  V3SupportPackageTimezones = "ET"
	V3_GMT V3SupportPackageTimezones = "GMT"
	V3_IST V3SupportPackageTimezones = "IST"
	V3_PT  V3SupportPackageTimezones = "PT"
)

var allowedV3SupportPackageTimezonesEnumValues = []V3SupportPackageTimezones{
	"ET",
	"GMT",
	"IST",
	"PT",
}

func (v *V3SupportPackageTimezones) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V3SupportPackageTimezones(value)
	for _, existing := range allowedV3SupportPackageTimezonesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V3SupportPackageTimezones", value)
}

// NewV3SupportPackageTimezonesFromValue returns a pointer to a valid V3SupportPackageTimezones
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV3SupportPackageTimezonesFromValue(v string) (*V3SupportPackageTimezones, error) {
	ev := V3SupportPackageTimezones(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V3SupportPackageTimezones: valid values are %v", v, allowedV3SupportPackageTimezonesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V3SupportPackageTimezones) IsValid() bool {
	for _, existing := range allowedV3SupportPackageTimezonesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v3SupportPackageTimezones value
func (v V3SupportPackageTimezones) Ptr() *V3SupportPackageTimezones {
	return &v
}

type NullableV3SupportPackageTimezones struct {
	value *V3SupportPackageTimezones
	isSet bool
}

func (v NullableV3SupportPackageTimezones) Get() *V3SupportPackageTimezones {
	return v.value
}

func (v *NullableV3SupportPackageTimezones) Set(val *V3SupportPackageTimezones) {
	v.value = val
	v.isSet = true
}

func (v NullableV3SupportPackageTimezones) IsSet() bool {
	return v.isSet
}

func (v *NullableV3SupportPackageTimezones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3SupportPackageTimezones(val *V3SupportPackageTimezones) *NullableV3SupportPackageTimezones {
	return &NullableV3SupportPackageTimezones{value: val, isSet: true}
}

func (v NullableV3SupportPackageTimezones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3SupportPackageTimezones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
