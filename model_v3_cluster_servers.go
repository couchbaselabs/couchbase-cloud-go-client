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

// V3ClusterServers struct for V3ClusterServers
type V3ClusterServers struct {
	Size int32 `json:"size"`
	// Compute instance type
	Compute string `json:"compute"`
	Services []string `json:"services"`
	Storage V3ClusterStorage `json:"storage"`
}

// NewV3ClusterServers instantiates a new V3ClusterServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3ClusterServers(size int32, compute string, services []string, storage V3ClusterStorage) *V3ClusterServers {
	this := V3ClusterServers{}
	this.Size = size
	this.Compute = compute
	this.Services = services
	this.Storage = storage
	return &this
}

// NewV3ClusterServersWithDefaults instantiates a new V3ClusterServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ClusterServersWithDefaults() *V3ClusterServers {
	this := V3ClusterServers{}
	return &this
}

// GetSize returns the Size field value
func (o *V3ClusterServers) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *V3ClusterServers) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *V3ClusterServers) SetSize(v int32) {
	o.Size = v
}

// GetCompute returns the Compute field value
func (o *V3ClusterServers) GetCompute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *V3ClusterServers) GetComputeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value
func (o *V3ClusterServers) SetCompute(v string) {
	o.Compute = v
}

// GetServices returns the Services field value
func (o *V3ClusterServers) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *V3ClusterServers) GetServicesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *V3ClusterServers) SetServices(v []string) {
	o.Services = v
}

// GetStorage returns the Storage field value
func (o *V3ClusterServers) GetStorage() V3ClusterStorage {
	if o == nil {
		var ret V3ClusterStorage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *V3ClusterServers) GetStorageOk() (*V3ClusterStorage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *V3ClusterServers) SetStorage(v V3ClusterStorage) {
	o.Storage = v
}

func (o V3ClusterServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["compute"] = o.Compute
	}
	if true {
		toSerialize["services"] = o.Services
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableV3ClusterServers struct {
	value *V3ClusterServers
	isSet bool
}

func (v NullableV3ClusterServers) Get() *V3ClusterServers {
	return v.value
}

func (v *NullableV3ClusterServers) Set(val *V3ClusterServers) {
	v.value = val
	v.isSet = true
}

func (v NullableV3ClusterServers) IsSet() bool {
	return v.isSet
}

func (v *NullableV3ClusterServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3ClusterServers(val *V3ClusterServers) *NullableV3ClusterServers {
	return &NullableV3ClusterServers{value: val, isSet: true}
}

func (v NullableV3ClusterServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3ClusterServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


