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

// CreateDatabaseUserRequest struct for CreateDatabaseUserRequest
type CreateDatabaseUserRequest struct {
	// Organisation user which is assigned to the database user. 
	UserId *string `json:"userId,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Access []BucketRole `json:"access"`
}

// NewCreateDatabaseUserRequest instantiates a new CreateDatabaseUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabaseUserRequest(username string, password string, access []BucketRole) *CreateDatabaseUserRequest {
	this := CreateDatabaseUserRequest{}
	this.Username = username
	this.Password = password
	this.Access = access
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

// GetAccess returns the Access field value
func (o *CreateDatabaseUserRequest) GetAccess() []BucketRole {
	if o == nil {
		var ret []BucketRole
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateDatabaseUserRequest) GetAccessOk() (*[]BucketRole, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateDatabaseUserRequest) SetAccess(v []BucketRole) {
	o.Access = v
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
	if true {
		toSerialize["access"] = o.Access
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


