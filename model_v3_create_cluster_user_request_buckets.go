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

// V3CreateClusterUserRequestBuckets struct for V3CreateClusterUserRequestBuckets
type V3CreateClusterUserRequestBuckets struct {
	Name string `json:"name"`
	// enter '*' for getting access to all scopes within a bucket or the scope name for getting access to a single scope
	Scope string `json:"scope"`
	Access V3BucketRoles `json:"access"`
}

// NewV3CreateClusterUserRequestBuckets instantiates a new V3CreateClusterUserRequestBuckets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3CreateClusterUserRequestBuckets(name string, scope string, access V3BucketRoles) *V3CreateClusterUserRequestBuckets {
	this := V3CreateClusterUserRequestBuckets{}
	this.Name = name
	this.Scope = scope
	this.Access = access
	return &this
}

// NewV3CreateClusterUserRequestBucketsWithDefaults instantiates a new V3CreateClusterUserRequestBuckets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3CreateClusterUserRequestBucketsWithDefaults() *V3CreateClusterUserRequestBuckets {
	this := V3CreateClusterUserRequestBuckets{}
	return &this
}

// GetName returns the Name field value
func (o *V3CreateClusterUserRequestBuckets) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterUserRequestBuckets) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V3CreateClusterUserRequestBuckets) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *V3CreateClusterUserRequestBuckets) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterUserRequestBuckets) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *V3CreateClusterUserRequestBuckets) SetScope(v string) {
	o.Scope = v
}

// GetAccess returns the Access field value
func (o *V3CreateClusterUserRequestBuckets) GetAccess() V3BucketRoles {
	if o == nil {
		var ret V3BucketRoles
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterUserRequestBuckets) GetAccessOk() (*V3BucketRoles, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *V3CreateClusterUserRequestBuckets) SetAccess(v V3BucketRoles) {
	o.Access = v
}

func (o V3CreateClusterUserRequestBuckets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableV3CreateClusterUserRequestBuckets struct {
	value *V3CreateClusterUserRequestBuckets
	isSet bool
}

func (v NullableV3CreateClusterUserRequestBuckets) Get() *V3CreateClusterUserRequestBuckets {
	return v.value
}

func (v *NullableV3CreateClusterUserRequestBuckets) Set(val *V3CreateClusterUserRequestBuckets) {
	v.value = val
	v.isSet = true
}

func (v NullableV3CreateClusterUserRequestBuckets) IsSet() bool {
	return v.isSet
}

func (v *NullableV3CreateClusterUserRequestBuckets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3CreateClusterUserRequestBuckets(val *V3CreateClusterUserRequestBuckets) *NullableV3CreateClusterUserRequestBuckets {
	return &NullableV3CreateClusterUserRequestBuckets{value: val, isSet: true}
}

func (v NullableV3CreateClusterUserRequestBuckets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3CreateClusterUserRequestBuckets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

