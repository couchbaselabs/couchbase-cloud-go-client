/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecloud

import (
	"encoding/json"
)

// CreateClusterRequest struct for CreateClusterRequest
type CreateClusterRequest struct {
	Name string `json:"name"`
	CloudId string `json:"cloudId"`
	ProjectId string `json:"projectId"`
	Servers *[]Server `json:"servers,omitempty"`
	SupportPackage *CreateClusterRequestSupportPackage `json:"supportPackage,omitempty"`
	Version *ClusterVersions `json:"version,omitempty"`
}

// NewCreateClusterRequest instantiates a new CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest(name string, cloudId string, projectId string) *CreateClusterRequest {
	this := CreateClusterRequest{}
	this.Name = name
	this.CloudId = cloudId
	this.ProjectId = projectId
	return &this
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClusterRequest) SetName(v string) {
	o.Name = v
}

// GetCloudId returns the CloudId field value
func (o *CreateClusterRequest) GetCloudId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetCloudIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudId, true
}

// SetCloudId sets field value
func (o *CreateClusterRequest) SetCloudId(v string) {
	o.CloudId = v
}

// GetProjectId returns the ProjectId field value
func (o *CreateClusterRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateClusterRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetServers() []Server {
	if o == nil || o.Servers == nil {
		var ret []Server
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetServersOk() (*[]Server, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *CreateClusterRequest) SetServers(v []Server) {
	o.Servers = &v
}

// GetSupportPackage returns the SupportPackage field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetSupportPackage() CreateClusterRequestSupportPackage {
	if o == nil || o.SupportPackage == nil {
		var ret CreateClusterRequestSupportPackage
		return ret
	}
	return *o.SupportPackage
}

// GetSupportPackageOk returns a tuple with the SupportPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetSupportPackageOk() (*CreateClusterRequestSupportPackage, bool) {
	if o == nil || o.SupportPackage == nil {
		return nil, false
	}
	return o.SupportPackage, true
}

// HasSupportPackage returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasSupportPackage() bool {
	if o != nil && o.SupportPackage != nil {
		return true
	}

	return false
}

// SetSupportPackage gets a reference to the given CreateClusterRequestSupportPackage and assigns it to the SupportPackage field.
func (o *CreateClusterRequest) SetSupportPackage(v CreateClusterRequestSupportPackage) {
	o.SupportPackage = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetVersion() ClusterVersions {
	if o == nil || o.Version == nil {
		var ret ClusterVersions
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetVersionOk() (*ClusterVersions, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given ClusterVersions and assigns it to the Version field.
func (o *CreateClusterRequest) SetVersion(v ClusterVersions) {
	o.Version = &v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cloudId"] = o.CloudId
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.SupportPackage != nil {
		toSerialize["supportPackage"] = o.SupportPackage
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterRequest struct {
	value *CreateClusterRequest
	isSet bool
}

func (v NullableCreateClusterRequest) Get() *CreateClusterRequest {
	return v.value
}

func (v *NullableCreateClusterRequest) Set(val *CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequest(val *CreateClusterRequest) *NullableCreateClusterRequest {
	return &NullableCreateClusterRequest{value: val, isSet: true}
}

func (v NullableCreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


