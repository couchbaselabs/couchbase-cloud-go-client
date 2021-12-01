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
	"time"
)

// V3Cluster struct for V3Cluster
type V3Cluster struct {
	Id                string             `json:"id"`
	Name              string             `json:"name"`
	TenantId          string             `json:"tenantId"`
	ProjectId         string             `json:"projectId"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
	Status            string             `json:"status"`
	Version           V3ClusterVersion   `json:"version"`
	EndpointsSrv      *string            `json:"endpointsSrv,omitempty"`
	Environment       string             `json:"environment"`
	Place             V3ClusterPlace     `json:"place"`
	Servers           []V3ClusterServers `json:"servers"`
	AvailabilityZones []string           `json:"availabilityZones"`
	Support           string             `json:"support"`
}

// NewV3Cluster instantiates a new V3Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3Cluster(id string, name string, tenantId string, projectId string, createdAt time.Time, updatedAt time.Time, status string, version V3ClusterVersion, environment string, place V3ClusterPlace, servers []V3ClusterServers, availabilityZones []string, support string) *V3Cluster {
	this := V3Cluster{}
	this.Id = id
	this.Name = name
	this.TenantId = tenantId
	this.ProjectId = projectId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.Version = version
	this.Environment = environment
	this.Place = place
	this.Servers = servers
	this.AvailabilityZones = availabilityZones
	this.Support = support
	return &this
}

// NewV3ClusterWithDefaults instantiates a new V3Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ClusterWithDefaults() *V3Cluster {
	this := V3Cluster{}
	return &this
}

// GetId returns the Id field value
func (o *V3Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V3Cluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *V3Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V3Cluster) SetName(v string) {
	o.Name = v
}

// GetTenantId returns the TenantId field value
func (o *V3Cluster) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *V3Cluster) SetTenantId(v string) {
	o.TenantId = v
}

// GetProjectId returns the ProjectId field value
func (o *V3Cluster) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *V3Cluster) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *V3Cluster) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *V3Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *V3Cluster) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *V3Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *V3Cluster) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V3Cluster) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *V3Cluster) GetVersion() V3ClusterVersion {
	if o == nil {
		var ret V3ClusterVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetVersionOk() (*V3ClusterVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *V3Cluster) SetVersion(v V3ClusterVersion) {
	o.Version = v
}

// GetEndpointsSrv returns the EndpointsSrv field value if set, zero value otherwise.
func (o *V3Cluster) GetEndpointsSrv() string {
	if o == nil || o.EndpointsSrv == nil {
		var ret string
		return ret
	}
	return *o.EndpointsSrv
}

// GetEndpointsSrvOk returns a tuple with the EndpointsSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetEndpointsSrvOk() (*string, bool) {
	if o == nil || o.EndpointsSrv == nil {
		return nil, false
	}
	return o.EndpointsSrv, true
}

// HasEndpointsSrv returns a boolean if a field has been set.
func (o *V3Cluster) HasEndpointsSrv() bool {
	if o != nil && o.EndpointsSrv != nil {
		return true
	}

	return false
}

// SetEndpointsSrv gets a reference to the given string and assigns it to the EndpointsSrv field.
func (o *V3Cluster) SetEndpointsSrv(v string) {
	o.EndpointsSrv = &v
}

// GetEnvironment returns the Environment field value
func (o *V3Cluster) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *V3Cluster) SetEnvironment(v string) {
	o.Environment = v
}

// GetPlace returns the Place field value
func (o *V3Cluster) GetPlace() V3ClusterPlace {
	if o == nil {
		var ret V3ClusterPlace
		return ret
	}

	return o.Place
}

// GetPlaceOk returns a tuple with the Place field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetPlaceOk() (*V3ClusterPlace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Place, true
}

// SetPlace sets field value
func (o *V3Cluster) SetPlace(v V3ClusterPlace) {
	o.Place = v
}

// GetServers returns the Servers field value
func (o *V3Cluster) GetServers() []V3ClusterServers {
	if o == nil {
		var ret []V3ClusterServers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetServersOk() (*[]V3ClusterServers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Servers, true
}

// SetServers sets field value
func (o *V3Cluster) SetServers(v []V3ClusterServers) {
	o.Servers = v
}

// GetAvailabilityZones returns the AvailabilityZones field value
func (o *V3Cluster) GetAvailabilityZones() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// SetAvailabilityZones sets field value
func (o *V3Cluster) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetSupport returns the Support field value
func (o *V3Cluster) GetSupport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Support
}

// GetSupportOk returns a tuple with the Support field value
// and a boolean to check if the value has been set.
func (o *V3Cluster) GetSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Support, true
}

// SetSupport sets field value
func (o *V3Cluster) SetSupport(v string) {
	o.Support = v
}

func (o V3Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.EndpointsSrv != nil {
		toSerialize["endpointsSrv"] = o.EndpointsSrv
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["place"] = o.Place
	}
	if true {
		toSerialize["servers"] = o.Servers
	}
	if true {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if true {
		toSerialize["support"] = o.Support
	}
	return json.Marshal(toSerialize)
}

type NullableV3Cluster struct {
	value *V3Cluster
	isSet bool
}

func (v NullableV3Cluster) Get() *V3Cluster {
	return v.value
}

func (v *NullableV3Cluster) Set(val *V3Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableV3Cluster) IsSet() bool {
	return v.isSet
}

func (v *NullableV3Cluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3Cluster(val *V3Cluster) *NullableV3Cluster {
	return &NullableV3Cluster{value: val, isSet: true}
}

func (v NullableV3Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3Cluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
