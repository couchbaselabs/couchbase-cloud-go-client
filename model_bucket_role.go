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

// BucketRole Specifies the roles of the user for the given bucket 
type BucketRole struct {
	// Bucket Name
	BucketName string `json:"bucketName"`
	BucketAccess []BucketRoleTypes `json:"bucketAccess"`
}

// NewBucketRole instantiates a new BucketRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketRole(bucketName string, bucketAccess []BucketRoleTypes) *BucketRole {
	this := BucketRole{}
	this.BucketName = bucketName
	this.BucketAccess = bucketAccess
	return &this
}

// NewBucketRoleWithDefaults instantiates a new BucketRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketRoleWithDefaults() *BucketRole {
	this := BucketRole{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *BucketRole) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *BucketRole) GetBucketNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *BucketRole) SetBucketName(v string) {
	o.BucketName = v
}

// GetBucketAccess returns the BucketAccess field value
func (o *BucketRole) GetBucketAccess() []BucketRoleTypes {
	if o == nil {
		var ret []BucketRoleTypes
		return ret
	}

	return o.BucketAccess
}

// GetBucketAccessOk returns a tuple with the BucketAccess field value
// and a boolean to check if the value has been set.
func (o *BucketRole) GetBucketAccessOk() (*[]BucketRoleTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BucketAccess, true
}

// SetBucketAccess sets field value
func (o *BucketRole) SetBucketAccess(v []BucketRoleTypes) {
	o.BucketAccess = v
}

func (o BucketRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	if true {
		toSerialize["bucketAccess"] = o.BucketAccess
	}
	return json.Marshal(toSerialize)
}

type NullableBucketRole struct {
	value *BucketRole
	isSet bool
}

func (v NullableBucketRole) Get() *BucketRole {
	return v.value
}

func (v *NullableBucketRole) Set(val *BucketRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketRole(val *BucketRole) *NullableBucketRole {
	return &NullableBucketRole{value: val, isSet: true}
}

func (v NullableBucketRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


