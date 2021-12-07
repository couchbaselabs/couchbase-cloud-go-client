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

// V3StorageType the model 'V3StorageType'
type V3StorageType string

// List of v3StorageType
const (
	GP3 V3StorageType = "GP3"
	IO2 V3StorageType = "IO2"
)

var allowedV3StorageTypeEnumValues = []V3StorageType{
	"GP3",
	"IO2",
}

func (v *V3StorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V3StorageType(value)
	for _, existing := range allowedV3StorageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V3StorageType", value)
}

// NewV3StorageTypeFromValue returns a pointer to a valid V3StorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV3StorageTypeFromValue(v string) (*V3StorageType, error) {
	ev := V3StorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V3StorageType: valid values are %v", v, allowedV3StorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V3StorageType) IsValid() bool {
	for _, existing := range allowedV3StorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v3StorageType value
func (v V3StorageType) Ptr() *V3StorageType {
	return &v
}

type NullableV3StorageType struct {
	value *V3StorageType
	isSet bool
}

func (v NullableV3StorageType) Get() *V3StorageType {
	return v.value
}

func (v *NullableV3StorageType) Set(val *V3StorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableV3StorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableV3StorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3StorageType(val *V3StorageType) *NullableV3StorageType {
	return &NullableV3StorageType{value: val, isSet: true}
}

func (v NullableV3StorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3StorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

