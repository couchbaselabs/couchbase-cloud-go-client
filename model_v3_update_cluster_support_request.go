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

// V3UpdateClusterSupportRequest struct for V3UpdateClusterSupportRequest
type V3UpdateClusterSupportRequest struct {
	SupportPackage V3UpdateClusterSupportRequestSupportPackage `json:"supportPackage"`
}

// NewV3UpdateClusterSupportRequest instantiates a new V3UpdateClusterSupportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3UpdateClusterSupportRequest(supportPackage V3UpdateClusterSupportRequestSupportPackage) *V3UpdateClusterSupportRequest {
	this := V3UpdateClusterSupportRequest{}
	this.SupportPackage = supportPackage
	return &this
}

// NewV3UpdateClusterSupportRequestWithDefaults instantiates a new V3UpdateClusterSupportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3UpdateClusterSupportRequestWithDefaults() *V3UpdateClusterSupportRequest {
	this := V3UpdateClusterSupportRequest{}
	return &this
}

// GetSupportPackage returns the SupportPackage field value
func (o *V3UpdateClusterSupportRequest) GetSupportPackage() V3UpdateClusterSupportRequestSupportPackage {
	if o == nil {
		var ret V3UpdateClusterSupportRequestSupportPackage
		return ret
	}

	return o.SupportPackage
}

// GetSupportPackageOk returns a tuple with the SupportPackage field value
// and a boolean to check if the value has been set.
func (o *V3UpdateClusterSupportRequest) GetSupportPackageOk() (*V3UpdateClusterSupportRequestSupportPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPackage, true
}

// SetSupportPackage sets field value
func (o *V3UpdateClusterSupportRequest) SetSupportPackage(v V3UpdateClusterSupportRequestSupportPackage) {
	o.SupportPackage = v
}

func (o V3UpdateClusterSupportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supportPackage"] = o.SupportPackage
	}
	return json.Marshal(toSerialize)
}

type NullableV3UpdateClusterSupportRequest struct {
	value *V3UpdateClusterSupportRequest
	isSet bool
}

func (v NullableV3UpdateClusterSupportRequest) Get() *V3UpdateClusterSupportRequest {
	return v.value
}

func (v *NullableV3UpdateClusterSupportRequest) Set(val *V3UpdateClusterSupportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3UpdateClusterSupportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3UpdateClusterSupportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3UpdateClusterSupportRequest(val *V3UpdateClusterSupportRequest) *NullableV3UpdateClusterSupportRequest {
	return &NullableV3UpdateClusterSupportRequest{value: val, isSet: true}
}

func (v NullableV3UpdateClusterSupportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3UpdateClusterSupportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}