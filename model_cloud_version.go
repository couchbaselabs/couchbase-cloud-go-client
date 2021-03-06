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

// CloudVersion struct for CloudVersion
type CloudVersion struct {
	Name string `json:"name"`
	Components map[string]string `json:"components"`
}

// NewCloudVersion instantiates a new CloudVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVersion(name string, components map[string]string) *CloudVersion {
	this := CloudVersion{}
	this.Name = name
	this.Components = components
	return &this
}

// NewCloudVersionWithDefaults instantiates a new CloudVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVersionWithDefaults() *CloudVersion {
	this := CloudVersion{}
	return &this
}

// GetName returns the Name field value
func (o *CloudVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudVersion) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudVersion) SetName(v string) {
	o.Name = v
}

// GetComponents returns the Components field value
func (o *CloudVersion) GetComponents() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *CloudVersion) GetComponentsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *CloudVersion) SetComponents(v map[string]string) {
	o.Components = v
}

func (o CloudVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableCloudVersion struct {
	value *CloudVersion
	isSet bool
}

func (v NullableCloudVersion) Get() *CloudVersion {
	return v.value
}

func (v *NullableCloudVersion) Set(val *CloudVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVersion(val *CloudVersion) *NullableCloudVersion {
	return &NullableCloudVersion{value: val, isSet: true}
}

func (v NullableCloudVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


