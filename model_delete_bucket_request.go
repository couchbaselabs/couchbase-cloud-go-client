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

// DeleteBucketRequest struct for DeleteBucketRequest
type DeleteBucketRequest struct {
	// The name of the bucket to be deleted
	Name string `json:"name"`
}

// NewDeleteBucketRequest instantiates a new DeleteBucketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBucketRequest(name string) *DeleteBucketRequest {
	this := DeleteBucketRequest{}
	this.Name = name
	return &this
}

// NewDeleteBucketRequestWithDefaults instantiates a new DeleteBucketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBucketRequestWithDefaults() *DeleteBucketRequest {
	this := DeleteBucketRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DeleteBucketRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteBucketRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteBucketRequest) SetName(v string) {
	o.Name = v
}

func (o DeleteBucketRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteBucketRequest struct {
	value *DeleteBucketRequest
	isSet bool
}

func (v NullableDeleteBucketRequest) Get() *DeleteBucketRequest {
	return v.value
}

func (v *NullableDeleteBucketRequest) Set(val *DeleteBucketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBucketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBucketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBucketRequest(val *DeleteBucketRequest) *NullableDeleteBucketRequest {
	return &NullableDeleteBucketRequest{value: val, isSet: true}
}

func (v NullableDeleteBucketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBucketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
