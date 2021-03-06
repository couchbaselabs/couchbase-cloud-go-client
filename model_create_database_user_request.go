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

// CreateDatabaseUserRequest struct for CreateDatabaseUserRequest
type CreateDatabaseUserRequest struct {
	// Organization user which is assigned to the database user. 
	UserId *string `json:"userId,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Buckets *[]BucketRole `json:"buckets,omitempty"`
	AllBucketsAccess *BucketRoleTypes `json:"allBucketsAccess,omitempty"`
}

// NewCreateDatabaseUserRequest instantiates a new CreateDatabaseUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabaseUserRequest(username string, password string) *CreateDatabaseUserRequest {
	this := CreateDatabaseUserRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewCreateDatabaseUserRequestWithDefaults instantiates a new CreateDatabaseUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseUserRequestWithDefaults() *CreateDatabaseUserRequest {
	this := CreateDatabaseUserRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreateDatabaseUserRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateDatabaseUserRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CreateDatabaseUserRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value
func (o *CreateDatabaseUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateDatabaseUserRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CreateDatabaseUserRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateDatabaseUserRequest) SetPassword(v string) {
	o.Password = v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *CreateDatabaseUserRequest) GetBuckets() []BucketRole {
	if o == nil || o.Buckets == nil {
		var ret []BucketRole
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetBucketsOk() (*[]BucketRole, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *CreateDatabaseUserRequest) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []BucketRole and assigns it to the Buckets field.
func (o *CreateDatabaseUserRequest) SetBuckets(v []BucketRole) {
	o.Buckets = &v
}

// GetAllBucketsAccess returns the AllBucketsAccess field value if set, zero value otherwise.
func (o *CreateDatabaseUserRequest) GetAllBucketsAccess() BucketRoleTypes {
	if o == nil || o.AllBucketsAccess == nil {
		var ret BucketRoleTypes
		return ret
	}
	return *o.AllBucketsAccess
}

// GetAllBucketsAccessOk returns a tuple with the AllBucketsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetAllBucketsAccessOk() (*BucketRoleTypes, bool) {
	if o == nil || o.AllBucketsAccess == nil {
		return nil, false
	}
	return o.AllBucketsAccess, true
}

// HasAllBucketsAccess returns a boolean if a field has been set.
func (o *CreateDatabaseUserRequest) HasAllBucketsAccess() bool {
	if o != nil && o.AllBucketsAccess != nil {
		return true
	}

	return false
}

// SetAllBucketsAccess gets a reference to the given BucketRoleTypes and assigns it to the AllBucketsAccess field.
func (o *CreateDatabaseUserRequest) SetAllBucketsAccess(v BucketRoleTypes) {
	o.AllBucketsAccess = &v
}

func (o CreateDatabaseUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.AllBucketsAccess != nil {
		toSerialize["allBucketsAccess"] = o.AllBucketsAccess
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDatabaseUserRequest struct {
	value *CreateDatabaseUserRequest
	isSet bool
}

func (v NullableCreateDatabaseUserRequest) Get() *CreateDatabaseUserRequest {
	return v.value
}

func (v *NullableCreateDatabaseUserRequest) Set(val *CreateDatabaseUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabaseUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabaseUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabaseUserRequest(val *CreateDatabaseUserRequest) *NullableCreateDatabaseUserRequest {
	return &NullableCreateDatabaseUserRequest{value: val, isSet: true}
}

func (v NullableCreateDatabaseUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabaseUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


