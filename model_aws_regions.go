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

// AwsRegions the model 'AwsRegions'
type AwsRegions string

// List of awsRegions
const (
	US_EAST_1 AwsRegions = "us-east-1"
	US_EAST_2 AwsRegions = "us-east-2"
	US_WEST_2 AwsRegions = "us-west-2"
	EU_WEST_1 AwsRegions = "eu-west-1"
	EU_CENTRAL_1 AwsRegions = "eu-central-1"
	EU_WEST_2 AwsRegions = "eu-west-2"
	EU_WEST_3 AwsRegions = "eu-west-3"
	EU_NORTH_1 AwsRegions = "eu-north-1"
	AP_SOUTHEAST_1 AwsRegions = "ap-southeast-1"
	AP_NORTHEAST_1 AwsRegions = "ap-northeast-1"
	AP_SOUTHEAST_2 AwsRegions = "ap-southeast-2"
	AP_NORTHEAST_2 AwsRegions = "ap-northeast-2"
	AP_SOUTH_1 AwsRegions = "ap-south-1"
	CA_CENTRAL_1 AwsRegions = "ca-central-1"
)

var allowedAwsRegionsEnumValues = []AwsRegions{
	"us-east-1",
	"us-east-2",
	"us-west-2",
	"eu-west-1",
	"eu-central-1",
	"eu-west-2",
	"eu-west-3",
	"eu-north-1",
	"ap-southeast-1",
	"ap-northeast-1",
	"ap-southeast-2",
	"ap-northeast-2",
	"ap-south-1",
	"ca-central-1",
}

func (v *AwsRegions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsRegions(value)
	for _, existing := range allowedAwsRegionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsRegions", value)
}

// NewAwsRegionsFromValue returns a pointer to a valid AwsRegions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsRegionsFromValue(v string) (*AwsRegions, error) {
	ev := AwsRegions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsRegions: valid values are %v", v, allowedAwsRegionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsRegions) IsValid() bool {
	for _, existing := range allowedAwsRegionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to awsRegions value
func (v AwsRegions) Ptr() *AwsRegions {
	return &v
}

type NullableAwsRegions struct {
	value *AwsRegions
	isSet bool
}

func (v NullableAwsRegions) Get() *AwsRegions {
	return v.value
}

func (v *NullableAwsRegions) Set(val *AwsRegions) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsRegions(val *AwsRegions) *NullableAwsRegions {
	return &NullableAwsRegions{value: val, isSet: true}
}

func (v NullableAwsRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

