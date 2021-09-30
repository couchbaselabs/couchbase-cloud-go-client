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

// ListCloudsResponse struct for ListCloudsResponse
type ListCloudsResponse struct {
	Cursor Cursor `json:"cursor"`
	Data []CloudSummary `json:"data"`
}

// NewListCloudsResponse instantiates a new ListCloudsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudsResponse(cursor Cursor, data []CloudSummary) *ListCloudsResponse {
	this := ListCloudsResponse{}
	this.Cursor = cursor
	this.Data = data
	return &this
}

// NewListCloudsResponseWithDefaults instantiates a new ListCloudsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudsResponseWithDefaults() *ListCloudsResponse {
	this := ListCloudsResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListCloudsResponse) GetCursor() Cursor {
	if o == nil {
		var ret Cursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListCloudsResponse) GetCursorOk() (*Cursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListCloudsResponse) SetCursor(v Cursor) {
	o.Cursor = v
}

// GetData returns the Data field value
func (o *ListCloudsResponse) GetData() []CloudSummary {
	if o == nil {
		var ret []CloudSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCloudsResponse) GetDataOk() (*[]CloudSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListCloudsResponse) SetData(v []CloudSummary) {
	o.Data = v
}

func (o ListCloudsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListCloudsResponse struct {
	value *ListCloudsResponse
	isSet bool
}

func (v NullableListCloudsResponse) Get() *ListCloudsResponse {
	return v.value
}

func (v *NullableListCloudsResponse) Set(val *ListCloudsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCloudsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCloudsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCloudsResponse(val *ListCloudsResponse) *NullableListCloudsResponse {
	return &NullableListCloudsResponse{value: val, isSet: true}
}

func (v NullableListCloudsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCloudsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


