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

// Server struct for Server
type Server struct {
	Size     int32               `json:"size"`
	Services []CouchbaseServices `json:"services"`
	Aws      *ServerAws          `json:"aws,omitempty"`
	Azure    *ServerAzure        `json:"azure,omitempty"`
}

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer(size int32, services []CouchbaseServices) *Server {
	this := Server{}
	this.Size = size
	this.Services = services
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	return &this
}

// GetSize returns the Size field value
func (o *Server) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Server) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Server) SetSize(v int32) {
	o.Size = v
}

// GetServices returns the Services field value
func (o *Server) GetServices() []CouchbaseServices {
	if o == nil {
		var ret []CouchbaseServices
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *Server) GetServicesOk() (*[]CouchbaseServices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *Server) SetServices(v []CouchbaseServices) {
	o.Services = v
}

// GetAws returns the Aws field value if set, zero value otherwise.
func (o *Server) GetAws() ServerAws {
	if o == nil || o.Aws == nil {
		var ret ServerAws
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetAwsOk() (*ServerAws, bool) {
	if o == nil || o.Aws == nil {
		return nil, false
	}
	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *Server) HasAws() bool {
	if o != nil && o.Aws != nil {
		return true
	}

	return false
}

// SetAws gets a reference to the given ServerAws and assigns it to the Aws field.
func (o *Server) SetAws(v ServerAws) {
	o.Aws = &v
}

// GetAzure returns the Azure field value if set, zero value otherwise.
func (o *Server) GetAzure() ServerAzure {
	if o == nil || o.Azure == nil {
		var ret ServerAzure
		return ret
	}
	return *o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetAzureOk() (*ServerAzure, bool) {
	if o == nil || o.Azure == nil {
		return nil, false
	}
	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *Server) HasAzure() bool {
	if o != nil && o.Azure != nil {
		return true
	}

	return false
}

// SetAzure gets a reference to the given ServerAzure and assigns it to the Azure field.
func (o *Server) SetAzure(v ServerAzure) {
	o.Azure = &v
}

func (o Server) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["services"] = o.Services
	}
	if o.Aws != nil {
		toSerialize["aws"] = o.Aws
	}
	if o.Azure != nil {
		toSerialize["azure"] = o.Azure
	}
	return json.Marshal(toSerialize)
}

type NullableServer struct {
	value *Server
	isSet bool
}

func (v NullableServer) Get() *Server {
	return v.value
}

func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

func (v NullableServer) IsSet() bool {
	return v.isSet
}

func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
