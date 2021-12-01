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

// ConflictResolution the model 'ConflictResolution'
type ConflictResolution string

// List of conflictResolution
const (
	LWW ConflictResolution = "lww"
	SEQNO ConflictResolution = "seqno"
)

var allowedConflictResolutionEnumValues = []ConflictResolution{
	"lww",
	"seqno",
}

func (v *ConflictResolution) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConflictResolution(value)
	for _, existing := range allowedConflictResolutionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConflictResolution", value)
}

// NewConflictResolutionFromValue returns a pointer to a valid ConflictResolution
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConflictResolutionFromValue(v string) (*ConflictResolution, error) {
	ev := ConflictResolution(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConflictResolution: valid values are %v", v, allowedConflictResolutionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConflictResolution) IsValid() bool {
	for _, existing := range allowedConflictResolutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to conflictResolution value
func (v ConflictResolution) Ptr() *ConflictResolution {
	return &v
}

type NullableConflictResolution struct {
	value *ConflictResolution
	isSet bool
}

func (v NullableConflictResolution) Get() *ConflictResolution {
	return v.value
}

func (v *NullableConflictResolution) Set(val *ConflictResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictResolution(val *ConflictResolution) *NullableConflictResolution {
	return &NullableConflictResolution{value: val, isSet: true}
}

func (v NullableConflictResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

