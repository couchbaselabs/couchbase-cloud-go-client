/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecloud

import (
	"encoding/json"
	"fmt"
)

// CloudStatus the model 'CloudStatus'
type CloudStatus string

// List of cloudStatus
const (
	NEEDS_BOOTSTRAP CloudStatus = "needs_bootstrap"
	PREFLIGHT_STARTED CloudStatus = "preflight_started"
	PREFLIGHT_FAILED CloudStatus = "preflight_failed"
	PREFLIGHT_SUCCEEDED CloudStatus = "preflight_succeeded"
	BOOTSTRAP_SUCCEEDED CloudStatus = "bootstrap_succeeded"
	DEPLOYING CloudStatus = "deploying"
	DEPLOY_FAILED CloudStatus = "deploy_failed"
	DEPLOY_SUCCEEDED CloudStatus = "deploy_succeeded"
	READY CloudStatus = "ready"
	DESTROYING CloudStatus = "destroying"
	DESTROY_FAILED CloudStatus = "destroy_failed"
	DESTROY_SUCCEEDED CloudStatus = "destroy_succeeded"
	NEEDS_CLEANUP CloudStatus = "needs_cleanup"
	MANAGEMENT_BLOCKED CloudStatus = "management_blocked"
)

var allowedCloudStatusEnumValues = []CloudStatus{
	"needs_bootstrap",
	"preflight_started",
	"preflight_failed",
	"preflight_succeeded",
	"bootstrap_succeeded",
	"deploying",
	"deploy_failed",
	"deploy_succeeded",
	"ready",
	"destroying",
	"destroy_failed",
	"destroy_succeeded",
	"needs_cleanup",
	"management_blocked",
}

func (v *CloudStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudStatus(value)
	for _, existing := range allowedCloudStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudStatus", value)
}

// NewCloudStatusFromValue returns a pointer to a valid CloudStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudStatusFromValue(v string) (*CloudStatus, error) {
	ev := CloudStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudStatus: valid values are %v", v, allowedCloudStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudStatus) IsValid() bool {
	for _, existing := range allowedCloudStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cloudStatus value
func (v CloudStatus) Ptr() *CloudStatus {
	return &v
}

type NullableCloudStatus struct {
	value *CloudStatus
	isSet bool
}

func (v NullableCloudStatus) Get() *CloudStatus {
	return v.value
}

func (v *NullableCloudStatus) Set(val *CloudStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStatus(val *CloudStatus) *NullableCloudStatus {
	return &NullableCloudStatus{value: val, isSet: true}
}

func (v NullableCloudStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
