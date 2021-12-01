/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
	"fmt"
)

// CouchbaseServices the model 'CouchbaseServices'
type CouchbaseServices string

// List of couchbaseServices
const (
	DATA CouchbaseServices = "data"
	INDEX CouchbaseServices = "index"
	QUERY CouchbaseServices = "query"
	SEARCH CouchbaseServices = "search"
	EVENTING CouchbaseServices = "eventing"
	ANALYTICS CouchbaseServices = "analytics"
)

var allowedCouchbaseServicesEnumValues = []CouchbaseServices{
	"data",
	"index",
	"query",
	"search",
	"eventing",
	"analytics",
}

func (v *CouchbaseServices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CouchbaseServices(value)
	for _, existing := range allowedCouchbaseServicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CouchbaseServices", value)
}

// NewCouchbaseServicesFromValue returns a pointer to a valid CouchbaseServices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCouchbaseServicesFromValue(v string) (*CouchbaseServices, error) {
	ev := CouchbaseServices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CouchbaseServices: valid values are %v", v, allowedCouchbaseServicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CouchbaseServices) IsValid() bool {
	for _, existing := range allowedCouchbaseServicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to couchbaseServices value
func (v CouchbaseServices) Ptr() *CouchbaseServices {
	return &v
}

type NullableCouchbaseServices struct {
	value *CouchbaseServices
	isSet bool
}

func (v NullableCouchbaseServices) Get() *CouchbaseServices {
	return v.value
}

func (v *NullableCouchbaseServices) Set(val *CouchbaseServices) {
	v.value = val
	v.isSet = true
}

func (v NullableCouchbaseServices) IsSet() bool {
	return v.isSet
}

func (v *NullableCouchbaseServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouchbaseServices(val *CouchbaseServices) *NullableCouchbaseServices {
	return &NullableCouchbaseServices{value: val, isSet: true}
}

func (v NullableCouchbaseServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouchbaseServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

