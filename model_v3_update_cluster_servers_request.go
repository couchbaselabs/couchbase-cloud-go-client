/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
)

// V3UpdateClusterServersRequest struct for V3UpdateClusterServersRequest
type V3UpdateClusterServersRequest struct {
	Servers []V3Servers `json:"servers"`
}

// NewV3UpdateClusterServersRequest instantiates a new V3UpdateClusterServersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3UpdateClusterServersRequest(servers []V3Servers) *V3UpdateClusterServersRequest {
	this := V3UpdateClusterServersRequest{}
	this.Servers = servers
	return &this
}

// NewV3UpdateClusterServersRequestWithDefaults instantiates a new V3UpdateClusterServersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3UpdateClusterServersRequestWithDefaults() *V3UpdateClusterServersRequest {
	this := V3UpdateClusterServersRequest{}
	return &this
}

// GetServers returns the Servers field value
func (o *V3UpdateClusterServersRequest) GetServers() []V3Servers {
	if o == nil {
		var ret []V3Servers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *V3UpdateClusterServersRequest) GetServersOk() (*[]V3Servers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Servers, true
}

// SetServers sets field value
func (o *V3UpdateClusterServersRequest) SetServers(v []V3Servers) {
	o.Servers = v
}

func (o V3UpdateClusterServersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableV3UpdateClusterServersRequest struct {
	value *V3UpdateClusterServersRequest
	isSet bool
}

func (v NullableV3UpdateClusterServersRequest) Get() *V3UpdateClusterServersRequest {
	return v.value
}

func (v *NullableV3UpdateClusterServersRequest) Set(val *V3UpdateClusterServersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3UpdateClusterServersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3UpdateClusterServersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3UpdateClusterServersRequest(val *V3UpdateClusterServersRequest) *NullableV3UpdateClusterServersRequest {
	return &NullableV3UpdateClusterServersRequest{value: val, isSet: true}
}

func (v NullableV3UpdateClusterServersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3UpdateClusterServersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


