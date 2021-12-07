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

// SupportPackageType the model 'SupportPackageType'
type SupportPackageType string

// List of supportPackageType
const (
	ENTERPRISE SupportPackageType = "enterprise"
	DEVELOPER_PRO SupportPackageType = "developerPro"
	FORUM_SUPPORT SupportPackageType = "forumSupport"
)

var allowedSupportPackageTypeEnumValues = []SupportPackageType{
	"enterprise",
	"developerPro",
	"forumSupport",
}

func (v *SupportPackageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportPackageType(value)
	for _, existing := range allowedSupportPackageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportPackageType", value)
}

// NewSupportPackageTypeFromValue returns a pointer to a valid SupportPackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportPackageTypeFromValue(v string) (*SupportPackageType, error) {
	ev := SupportPackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportPackageType: valid values are %v", v, allowedSupportPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportPackageType) IsValid() bool {
	for _, existing := range allowedSupportPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to supportPackageType value
func (v SupportPackageType) Ptr() *SupportPackageType {
	return &v
}

type NullableSupportPackageType struct {
	value *SupportPackageType
	isSet bool
}

func (v NullableSupportPackageType) Get() *SupportPackageType {
	return v.value
}

func (v *NullableSupportPackageType) Set(val *SupportPackageType) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportPackageType) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportPackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportPackageType(val *SupportPackageType) *NullableSupportPackageType {
	return &NullableSupportPackageType{value: val, isSet: true}
}

func (v NullableSupportPackageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportPackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

