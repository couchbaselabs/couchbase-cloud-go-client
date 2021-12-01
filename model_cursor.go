/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// Couchbase, Inc. licenses this to you under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://www.apache.org/licenses/LICENSE-2.0.

// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package couchbasecapella

import (
	"encoding/json"
)

// Cursor struct for Cursor
type Cursor struct {
	Pages CursorPages `json:"pages"`
	Hrefs CursorHrefs `json:"hrefs"`
}

// NewCursor instantiates a new Cursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursor(pages CursorPages, hrefs CursorHrefs) *Cursor {
	this := Cursor{}
	this.Pages = pages
	this.Hrefs = hrefs
	return &this
}

// NewCursorWithDefaults instantiates a new Cursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorWithDefaults() *Cursor {
	this := Cursor{}
	return &this
}

// GetPages returns the Pages field value
func (o *Cursor) GetPages() CursorPages {
	if o == nil {
		var ret CursorPages
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *Cursor) GetPagesOk() (*CursorPages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *Cursor) SetPages(v CursorPages) {
	o.Pages = v
}

// GetHrefs returns the Hrefs field value
func (o *Cursor) GetHrefs() CursorHrefs {
	if o == nil {
		var ret CursorHrefs
		return ret
	}

	return o.Hrefs
}

// GetHrefsOk returns a tuple with the Hrefs field value
// and a boolean to check if the value has been set.
func (o *Cursor) GetHrefsOk() (*CursorHrefs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hrefs, true
}

// SetHrefs sets field value
func (o *Cursor) SetHrefs(v CursorHrefs) {
	o.Hrefs = v
}

func (o Cursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pages"] = o.Pages
	}
	if true {
		toSerialize["hrefs"] = o.Hrefs
	}
	return json.Marshal(toSerialize)
}

type NullableCursor struct {
	value *Cursor
	isSet bool
}

func (v NullableCursor) Get() *Cursor {
	return v.value
}

func (v *NullableCursor) Set(val *Cursor) {
	v.value = val
	v.isSet = true
}

func (v NullableCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursor(val *Cursor) *NullableCursor {
	return &NullableCursor{value: val, isSet: true}
}

func (v NullableCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
