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

// CloudProviderSettings - struct for CloudProviderSettings
type CloudProviderSettings struct {
	AWSProviderSettings   *AWSProviderSettings
	AzureProviderSettings *AzureProviderSettings
}

// AWSProviderSettingsAsCloudProviderSettings is a convenience function that returns AWSProviderSettings wrapped in CloudProviderSettings
func AWSProviderSettingsAsCloudProviderSettings(v *AWSProviderSettings) CloudProviderSettings {
	return CloudProviderSettings{AWSProviderSettings: v}
}

// AzureProviderSettingsAsCloudProviderSettings is a convenience function that returns AzureProviderSettings wrapped in CloudProviderSettings
func AzureProviderSettingsAsCloudProviderSettings(v *AzureProviderSettings) CloudProviderSettings {
	return CloudProviderSettings{AzureProviderSettings: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CloudProviderSettings) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSProviderSettings
	err = json.Unmarshal(data, &dst.AWSProviderSettings)
	if err == nil {
		jsonAWSProviderSettings, _ := json.Marshal(dst.AWSProviderSettings)
		if string(jsonAWSProviderSettings) == "{}" { // empty struct
			dst.AWSProviderSettings = nil
		} else {
			match++
		}
	} else {
		dst.AWSProviderSettings = nil
	}

	// try to unmarshal data into AzureProviderSettings
	err = json.Unmarshal(data, &dst.AzureProviderSettings)
	if err == nil {
		jsonAzureProviderSettings, _ := json.Marshal(dst.AzureProviderSettings)
		if string(jsonAzureProviderSettings) == "{}" { // empty struct
			dst.AzureProviderSettings = nil
		} else {
			match++
		}
	} else {
		dst.AzureProviderSettings = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSProviderSettings = nil
		dst.AzureProviderSettings = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CloudProviderSettings)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CloudProviderSettings)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CloudProviderSettings) MarshalJSON() ([]byte, error) {
	if src.AWSProviderSettings != nil {
		return json.Marshal(&src.AWSProviderSettings)
	}

	if src.AzureProviderSettings != nil {
		return json.Marshal(&src.AzureProviderSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CloudProviderSettings) GetActualInstance() interface{} {
	if obj.AWSProviderSettings != nil {
		return obj.AWSProviderSettings
	}

	if obj.AzureProviderSettings != nil {
		return obj.AzureProviderSettings
	}

	// all schemas are nil
	return nil
}

type NullableCloudProviderSettings struct {
	value *CloudProviderSettings
	isSet bool
}

func (v NullableCloudProviderSettings) Get() *CloudProviderSettings {
	return v.value
}

func (v *NullableCloudProviderSettings) Set(val *CloudProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderSettings(val *CloudProviderSettings) *NullableCloudProviderSettings {
	return &NullableCloudProviderSettings{value: val, isSet: true}
}

func (v NullableCloudProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
