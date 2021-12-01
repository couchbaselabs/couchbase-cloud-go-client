/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// Couchbase, Inc. licenses this to you under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://www.apache.org/licenses/LICENSE-2.0.

// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package couchbasecapella

import (
	"encoding/json"
	"fmt"
)

// AzureInstances the model 'AzureInstances'
type AzureInstances string

// List of azureInstances
const (
	F4S_V2  AzureInstances = "Standard_F4s_v2"
	F8S_V2  AzureInstances = "Standard_F8s_v2"
	F16S_V2 AzureInstances = "Standard_F16s_v2"
	F32S_V2 AzureInstances = "Standard_F32s_v2"
	F48S_V2 AzureInstances = "Standard_F48s_v2"
	F64S_V2 AzureInstances = "Standard_F64s_v2"
	F72S_V2 AzureInstances = "Standard_F72s_v2"
	D4S_V3  AzureInstances = "Standard_D4s_v3"
	D8S_V3  AzureInstances = "Standard_D8s_v3"
	D16S_V3 AzureInstances = "Standard_D16s_v3"
	D32S_V3 AzureInstances = "Standard_D32s_v3"
	D48S_V3 AzureInstances = "Standard_D48s_v3"
	D64S_V3 AzureInstances = "Standard_D64s_v3"
	E4S_V3  AzureInstances = "Standard_E4s_v3"
	E8S_V3  AzureInstances = "Standard_E8s_v3"
	E16S_V3 AzureInstances = "Standard_E16s_v3"
	E20S_V3 AzureInstances = "Standard_E20s_v3"
	E32S_V3 AzureInstances = "Standard_E32s_v3"
	E48S_V3 AzureInstances = "Standard_E48s_v3"
	E64S_V3 AzureInstances = "Standard_E64s_v3"
	GS2     AzureInstances = "Standard_GS2"
	GS3     AzureInstances = "Standard_GS3"
	GS4     AzureInstances = "Standard_GS4"
	GS5     AzureInstances = "Standard_GS5"
)

var allowedAzureInstancesEnumValues = []AzureInstances{
	"Standard_F4s_v2",
	"Standard_F8s_v2",
	"Standard_F16s_v2",
	"Standard_F32s_v2",
	"Standard_F48s_v2",
	"Standard_F64s_v2",
	"Standard_F72s_v2",
	"Standard_D4s_v3",
	"Standard_D8s_v3",
	"Standard_D16s_v3",
	"Standard_D32s_v3",
	"Standard_D48s_v3",
	"Standard_D64s_v3",
	"Standard_E4s_v3",
	"Standard_E8s_v3",
	"Standard_E16s_v3",
	"Standard_E20s_v3",
	"Standard_E32s_v3",
	"Standard_E48s_v3",
	"Standard_E64s_v3",
	"Standard_GS2",
	"Standard_GS3",
	"Standard_GS4",
	"Standard_GS5",
}

func (v *AzureInstances) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureInstances(value)
	for _, existing := range allowedAzureInstancesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureInstances", value)
}

// NewAzureInstancesFromValue returns a pointer to a valid AzureInstances
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureInstancesFromValue(v string) (*AzureInstances, error) {
	ev := AzureInstances(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureInstances: valid values are %v", v, allowedAzureInstancesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureInstances) IsValid() bool {
	for _, existing := range allowedAzureInstancesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to azureInstances value
func (v AzureInstances) Ptr() *AzureInstances {
	return &v
}

type NullableAzureInstances struct {
	value *AzureInstances
	isSet bool
}

func (v NullableAzureInstances) Get() *AzureInstances {
	return v.value
}

func (v *NullableAzureInstances) Set(val *AzureInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureInstances(val *AzureInstances) *NullableAzureInstances {
	return &NullableAzureInstances{value: val, isSet: true}
}

func (v NullableAzureInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
