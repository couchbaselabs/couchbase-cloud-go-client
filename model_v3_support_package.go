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
)

// V3SupportPackage One from \"Basic\", \"DeveloperPro\" or \"Enterprise\"
type V3SupportPackage struct {
	Timezone V3SupportPackageTimezones `json:"timezone"`
	Type     V3SupportPackageType      `json:"type"`
}

// NewV3SupportPackage instantiates a new V3SupportPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3SupportPackage(timezone V3SupportPackageTimezones, type_ V3SupportPackageType) *V3SupportPackage {
	this := V3SupportPackage{}
	this.Timezone = timezone
	this.Type = type_
	return &this
}

// NewV3SupportPackageWithDefaults instantiates a new V3SupportPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3SupportPackageWithDefaults() *V3SupportPackage {
	this := V3SupportPackage{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *V3SupportPackage) GetTimezone() V3SupportPackageTimezones {
	if o == nil {
		var ret V3SupportPackageTimezones
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *V3SupportPackage) GetTimezoneOk() (*V3SupportPackageTimezones, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *V3SupportPackage) SetTimezone(v V3SupportPackageTimezones) {
	o.Timezone = v
}

// GetType returns the Type field value
func (o *V3SupportPackage) GetType() V3SupportPackageType {
	if o == nil {
		var ret V3SupportPackageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V3SupportPackage) GetTypeOk() (*V3SupportPackageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V3SupportPackage) SetType(v V3SupportPackageType) {
	o.Type = v
}

func (o V3SupportPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableV3SupportPackage struct {
	value *V3SupportPackage
	isSet bool
}

func (v NullableV3SupportPackage) Get() *V3SupportPackage {
	return v.value
}

func (v *NullableV3SupportPackage) Set(val *V3SupportPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableV3SupportPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableV3SupportPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3SupportPackage(val *V3SupportPackage) *NullableV3SupportPackage {
	return &NullableV3SupportPackage{value: val, isSet: true}
}

func (v NullableV3SupportPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3SupportPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
