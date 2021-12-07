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

// NodeStats struct for NodeStats
type NodeStats struct {
	Status string `json:"status"`
	TotalCount int32 `json:"totalCount"`
	FailedOverCount int32 `json:"failedOverCount"`
	UnhealthyCount int32 `json:"unhealthyCount"`
	ServiceStats []ServiceStats `json:"serviceStats"`
}

// NewNodeStats instantiates a new NodeStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeStats(status string, totalCount int32, failedOverCount int32, unhealthyCount int32, serviceStats []ServiceStats) *NodeStats {
	this := NodeStats{}
	this.Status = status
	this.TotalCount = totalCount
	this.FailedOverCount = failedOverCount
	this.UnhealthyCount = unhealthyCount
	this.ServiceStats = serviceStats
	return &this
}

// NewNodeStatsWithDefaults instantiates a new NodeStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeStatsWithDefaults() *NodeStats {
	this := NodeStats{}
	return &this
}

// GetStatus returns the Status field value
func (o *NodeStats) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NodeStats) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NodeStats) SetStatus(v string) {
	o.Status = v
}

// GetTotalCount returns the TotalCount field value
func (o *NodeStats) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *NodeStats) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *NodeStats) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetFailedOverCount returns the FailedOverCount field value
func (o *NodeStats) GetFailedOverCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailedOverCount
}

// GetFailedOverCountOk returns a tuple with the FailedOverCount field value
// and a boolean to check if the value has been set.
func (o *NodeStats) GetFailedOverCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FailedOverCount, true
}

// SetFailedOverCount sets field value
func (o *NodeStats) SetFailedOverCount(v int32) {
	o.FailedOverCount = v
}

// GetUnhealthyCount returns the UnhealthyCount field value
func (o *NodeStats) GetUnhealthyCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnhealthyCount
}

// GetUnhealthyCountOk returns a tuple with the UnhealthyCount field value
// and a boolean to check if the value has been set.
func (o *NodeStats) GetUnhealthyCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnhealthyCount, true
}

// SetUnhealthyCount sets field value
func (o *NodeStats) SetUnhealthyCount(v int32) {
	o.UnhealthyCount = v
}

// GetServiceStats returns the ServiceStats field value
func (o *NodeStats) GetServiceStats() []ServiceStats {
	if o == nil {
		var ret []ServiceStats
		return ret
	}

	return o.ServiceStats
}

// GetServiceStatsOk returns a tuple with the ServiceStats field value
// and a boolean to check if the value has been set.
func (o *NodeStats) GetServiceStatsOk() (*[]ServiceStats, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceStats, true
}

// SetServiceStats sets field value
func (o *NodeStats) SetServiceStats(v []ServiceStats) {
	o.ServiceStats = v
}

func (o NodeStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if true {
		toSerialize["failedOverCount"] = o.FailedOverCount
	}
	if true {
		toSerialize["unhealthyCount"] = o.UnhealthyCount
	}
	if true {
		toSerialize["serviceStats"] = o.ServiceStats
	}
	return json.Marshal(toSerialize)
}

type NullableNodeStats struct {
	value *NodeStats
	isSet bool
}

func (v NullableNodeStats) Get() *NodeStats {
	return v.value
}

func (v *NullableNodeStats) Set(val *NodeStats) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeStats) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeStats(val *NodeStats) *NullableNodeStats {
	return &NullableNodeStats{value: val, isSet: true}
}

func (v NullableNodeStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


