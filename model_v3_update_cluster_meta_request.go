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

// V3UpdateClusterMetaRequest struct for V3UpdateClusterMetaRequest
type V3UpdateClusterMetaRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewV3UpdateClusterMetaRequest instantiates a new V3UpdateClusterMetaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3UpdateClusterMetaRequest() *V3UpdateClusterMetaRequest {
	this := V3UpdateClusterMetaRequest{}
	return &this
}

// NewV3UpdateClusterMetaRequestWithDefaults instantiates a new V3UpdateClusterMetaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3UpdateClusterMetaRequestWithDefaults() *V3UpdateClusterMetaRequest {
	this := V3UpdateClusterMetaRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V3UpdateClusterMetaRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3UpdateClusterMetaRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V3UpdateClusterMetaRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V3UpdateClusterMetaRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V3UpdateClusterMetaRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3UpdateClusterMetaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V3UpdateClusterMetaRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V3UpdateClusterMetaRequest) SetDescription(v string) {
	o.Description = &v
}

func (o V3UpdateClusterMetaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableV3UpdateClusterMetaRequest struct {
	value *V3UpdateClusterMetaRequest
	isSet bool
}

func (v NullableV3UpdateClusterMetaRequest) Get() *V3UpdateClusterMetaRequest {
	return v.value
}

func (v *NullableV3UpdateClusterMetaRequest) Set(val *V3UpdateClusterMetaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3UpdateClusterMetaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3UpdateClusterMetaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3UpdateClusterMetaRequest(val *V3UpdateClusterMetaRequest) *NullableV3UpdateClusterMetaRequest {
	return &NullableV3UpdateClusterMetaRequest{value: val, isSet: true}
}

func (v NullableV3UpdateClusterMetaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3UpdateClusterMetaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}