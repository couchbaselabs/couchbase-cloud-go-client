/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateClusterRequestSupportPackage if support package is not set then default support package is used. 
type CreateClusterRequestSupportPackage struct {
	Type SupportPackageType `json:"type"`
	Timezone SupportTimezones `json:"timezone"`
}

// NewCreateClusterRequestSupportPackage instantiates a new CreateClusterRequestSupportPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequestSupportPackage(type_ SupportPackageType, timezone SupportTimezones) *CreateClusterRequestSupportPackage {
	this := CreateClusterRequestSupportPackage{}
	this.Type = type_
	this.Timezone = timezone
	return &this
}

// NewCreateClusterRequestSupportPackageWithDefaults instantiates a new CreateClusterRequestSupportPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestSupportPackageWithDefaults() *CreateClusterRequestSupportPackage {
	this := CreateClusterRequestSupportPackage{}
	return &this
}

// GetType returns the Type field value
func (o *CreateClusterRequestSupportPackage) GetType() SupportPackageType {
	if o == nil {
		var ret SupportPackageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestSupportPackage) GetTypeOk() (*SupportPackageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateClusterRequestSupportPackage) SetType(v SupportPackageType) {
	o.Type = v
}

// GetTimezone returns the Timezone field value
func (o *CreateClusterRequestSupportPackage) GetTimezone() SupportTimezones {
	if o == nil {
		var ret SupportTimezones
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestSupportPackage) GetTimezoneOk() (*SupportTimezones, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *CreateClusterRequestSupportPackage) SetTimezone(v SupportTimezones) {
	o.Timezone = v
}

func (o CreateClusterRequestSupportPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterRequestSupportPackage struct {
	value *CreateClusterRequestSupportPackage
	isSet bool
}

func (v NullableCreateClusterRequestSupportPackage) Get() *CreateClusterRequestSupportPackage {
	return v.value
}

func (v *NullableCreateClusterRequestSupportPackage) Set(val *CreateClusterRequestSupportPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequestSupportPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequestSupportPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequestSupportPackage(val *CreateClusterRequestSupportPackage) *NullableCreateClusterRequestSupportPackage {
	return &NullableCreateClusterRequestSupportPackage{value: val, isSet: true}
}

func (v NullableCreateClusterRequestSupportPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequestSupportPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


