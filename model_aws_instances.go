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

// AwsInstances the model 'AwsInstances'
type AwsInstances string

// List of awsInstances
const (
	M5_LARGE    AwsInstances = "m5.large"
	M5_XLARGE   AwsInstances = "m5.xlarge"
	M5_2XLARGE  AwsInstances = "m5.2xlarge"
	M5_4XLARGE  AwsInstances = "m5.4xlarge"
	M5_8XLARGE  AwsInstances = "m5.8xlarge"
	M5_12XLARGE AwsInstances = "m5.12xlarge"
	M5_16XLARGE AwsInstances = "m5.16xlarge"
	M5_24XLARGE AwsInstances = "m5.24xlarge"
	R5_LARGE    AwsInstances = "r5.large"
	R5_XLARGE   AwsInstances = "r5.xlarge"
	R5_2XLARGE  AwsInstances = "r5.2xlarge"
	R5_4XLARGE  AwsInstances = "r5.4xlarge"
	R5_8XLARGE  AwsInstances = "r5.8xlarge"
	R5_12XLARGE AwsInstances = "r5.12xlarge"
	R5_24XLARGE AwsInstances = "r5.24xlarge"
	C5_LARGE    AwsInstances = "c5.large"
	C5_XLARGE   AwsInstances = "c5.xlarge"
	C5_2XLARGE  AwsInstances = "c5.2xlarge"
	C5_4XLARGE  AwsInstances = "c5.4xlarge"
	C5_9XLARGE  AwsInstances = "c5.9xlarge"
	C5_12XLARGE AwsInstances = "c5.12xlarge"
	C5_18XLARGE AwsInstances = "c5.18xlarge"
	X1_16XLARGE AwsInstances = "x1.16xlarge"
	X1_32XLARGE AwsInstances = "x1.32xlarge"
)

var allowedAwsInstancesEnumValues = []AwsInstances{
	"m5.large",
	"m5.xlarge",
	"m5.2xlarge",
	"m5.4xlarge",
	"m5.8xlarge",
	"m5.12xlarge",
	"m5.16xlarge",
	"m5.24xlarge",
	"r5.large",
	"r5.xlarge",
	"r5.2xlarge",
	"r5.4xlarge",
	"r5.8xlarge",
	"r5.12xlarge",
	"r5.24xlarge",
	"c5.large",
	"c5.xlarge",
	"c5.2xlarge",
	"c5.4xlarge",
	"c5.9xlarge",
	"c5.12xlarge",
	"c5.18xlarge",
	"x1.16xlarge",
	"x1.32xlarge",
}

func (v *AwsInstances) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsInstances(value)
	for _, existing := range allowedAwsInstancesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsInstances", value)
}

// NewAwsInstancesFromValue returns a pointer to a valid AwsInstances
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsInstancesFromValue(v string) (*AwsInstances, error) {
	ev := AwsInstances(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsInstances: valid values are %v", v, allowedAwsInstancesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsInstances) IsValid() bool {
	for _, existing := range allowedAwsInstancesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to awsInstances value
func (v AwsInstances) Ptr() *AwsInstances {
	return &v
}

type NullableAwsInstances struct {
	value *AwsInstances
	isSet bool
}

func (v NullableAwsInstances) Get() *AwsInstances {
	return v.value
}

func (v *NullableAwsInstances) Set(val *AwsInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsInstances(val *AwsInstances) *NullableAwsInstances {
	return &NullableAwsInstances{value: val, isSet: true}
}

func (v NullableAwsInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
