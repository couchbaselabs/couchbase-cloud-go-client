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

// V3PlaceHosted cluster will be deployed in the managed couchbase capella
type V3PlaceHosted struct {
	Provider V3Provider `json:"provider"`
	// A valid region for the cloudProvider
	Region string `json:"region"`
	// CIDR block.
	CIDR string `json:"CIDR"`
}

// NewV3PlaceHosted instantiates a new V3PlaceHosted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3PlaceHosted(provider V3Provider, region string, cIDR string) *V3PlaceHosted {
	this := V3PlaceHosted{}
	this.Provider = provider
	this.Region = region
	this.CIDR = cIDR
	return &this
}

// NewV3PlaceHostedWithDefaults instantiates a new V3PlaceHosted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3PlaceHostedWithDefaults() *V3PlaceHosted {
	this := V3PlaceHosted{}
	return &this
}

// GetProvider returns the Provider field value
func (o *V3PlaceHosted) GetProvider() V3Provider {
	if o == nil {
		var ret V3Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *V3PlaceHosted) GetProviderOk() (*V3Provider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *V3PlaceHosted) SetProvider(v V3Provider) {
	o.Provider = v
}

// GetRegion returns the Region field value
func (o *V3PlaceHosted) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *V3PlaceHosted) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *V3PlaceHosted) SetRegion(v string) {
	o.Region = v
}

// GetCIDR returns the CIDR field value
func (o *V3PlaceHosted) GetCIDR() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CIDR
}

// GetCIDROk returns a tuple with the CIDR field value
// and a boolean to check if the value has been set.
func (o *V3PlaceHosted) GetCIDROk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CIDR, true
}

// SetCIDR sets field value
func (o *V3PlaceHosted) SetCIDR(v string) {
	o.CIDR = v
}

func (o V3PlaceHosted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["CIDR"] = o.CIDR
	}
	return json.Marshal(toSerialize)
}

type NullableV3PlaceHosted struct {
	value *V3PlaceHosted
	isSet bool
}

func (v NullableV3PlaceHosted) Get() *V3PlaceHosted {
	return v.value
}

func (v *NullableV3PlaceHosted) Set(val *V3PlaceHosted) {
	v.value = val
	v.isSet = true
}

func (v NullableV3PlaceHosted) IsSet() bool {
	return v.isSet
}

func (v *NullableV3PlaceHosted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3PlaceHosted(val *V3PlaceHosted) *NullableV3PlaceHosted {
	return &NullableV3PlaceHosted{value: val, isSet: true}
}

func (v NullableV3PlaceHosted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3PlaceHosted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


