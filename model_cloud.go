/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
	"time"
)

// Cloud struct for Cloud
type Cloud struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TenantId string `json:"tenantId"`
	Status CloudStatus `json:"status"`
	Provider Provider `json:"provider"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	PreflightStartedAt *time.Time `json:"preflightStartedAt,omitempty"`
	PreflightFailedAt *time.Time `json:"preflightFailedAt,omitempty"`
	PreflightCompletedAt *time.Time `json:"preflightCompletedAt,omitempty"`
	BootstrappedAt *time.Time `json:"bootstrappedAt,omitempty"`
	DeployedAt *time.Time `json:"deployedAt,omitempty"`
	DestroyedAt *time.Time `json:"destroyedAt,omitempty"`
	ProviderSettings CloudProviderSettings `json:"providerSettings"`
	Version CloudVersion `json:"version"`
}

// NewCloud instantiates a new Cloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloud(id string, name string, tenantId string, status CloudStatus, provider Provider, createdAt time.Time, providerSettings CloudProviderSettings, version CloudVersion) *Cloud {
	this := Cloud{}
	this.Id = id
	this.Name = name
	this.TenantId = tenantId
	this.Status = status
	this.Provider = provider
	this.CreatedAt = createdAt
	this.ProviderSettings = providerSettings
	this.Version = version
	return &this
}

// NewCloudWithDefaults instantiates a new Cloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudWithDefaults() *Cloud {
	this := Cloud{}
	return &this
}

// GetId returns the Id field value
func (o *Cloud) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cloud) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Cloud) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cloud) SetName(v string) {
	o.Name = v
}

// GetTenantId returns the TenantId field value
func (o *Cloud) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Cloud) SetTenantId(v string) {
	o.TenantId = v
}

// GetStatus returns the Status field value
func (o *Cloud) GetStatus() CloudStatus {
	if o == nil {
		var ret CloudStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetStatusOk() (*CloudStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Cloud) SetStatus(v CloudStatus) {
	o.Status = v
}

// GetProvider returns the Provider field value
func (o *Cloud) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetProviderOk() (*Provider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Cloud) SetProvider(v Provider) {
	o.Provider = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Cloud) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Cloud) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cloud) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Cloud) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cloud) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetPreflightStartedAt returns the PreflightStartedAt field value if set, zero value otherwise.
func (o *Cloud) GetPreflightStartedAt() time.Time {
	if o == nil || o.PreflightStartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PreflightStartedAt
}

// GetPreflightStartedAtOk returns a tuple with the PreflightStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetPreflightStartedAtOk() (*time.Time, bool) {
	if o == nil || o.PreflightStartedAt == nil {
		return nil, false
	}
	return o.PreflightStartedAt, true
}

// HasPreflightStartedAt returns a boolean if a field has been set.
func (o *Cloud) HasPreflightStartedAt() bool {
	if o != nil && o.PreflightStartedAt != nil {
		return true
	}

	return false
}

// SetPreflightStartedAt gets a reference to the given time.Time and assigns it to the PreflightStartedAt field.
func (o *Cloud) SetPreflightStartedAt(v time.Time) {
	o.PreflightStartedAt = &v
}

// GetPreflightFailedAt returns the PreflightFailedAt field value if set, zero value otherwise.
func (o *Cloud) GetPreflightFailedAt() time.Time {
	if o == nil || o.PreflightFailedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PreflightFailedAt
}

// GetPreflightFailedAtOk returns a tuple with the PreflightFailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetPreflightFailedAtOk() (*time.Time, bool) {
	if o == nil || o.PreflightFailedAt == nil {
		return nil, false
	}
	return o.PreflightFailedAt, true
}

// HasPreflightFailedAt returns a boolean if a field has been set.
func (o *Cloud) HasPreflightFailedAt() bool {
	if o != nil && o.PreflightFailedAt != nil {
		return true
	}

	return false
}

// SetPreflightFailedAt gets a reference to the given time.Time and assigns it to the PreflightFailedAt field.
func (o *Cloud) SetPreflightFailedAt(v time.Time) {
	o.PreflightFailedAt = &v
}

// GetPreflightCompletedAt returns the PreflightCompletedAt field value if set, zero value otherwise.
func (o *Cloud) GetPreflightCompletedAt() time.Time {
	if o == nil || o.PreflightCompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PreflightCompletedAt
}

// GetPreflightCompletedAtOk returns a tuple with the PreflightCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetPreflightCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.PreflightCompletedAt == nil {
		return nil, false
	}
	return o.PreflightCompletedAt, true
}

// HasPreflightCompletedAt returns a boolean if a field has been set.
func (o *Cloud) HasPreflightCompletedAt() bool {
	if o != nil && o.PreflightCompletedAt != nil {
		return true
	}

	return false
}

// SetPreflightCompletedAt gets a reference to the given time.Time and assigns it to the PreflightCompletedAt field.
func (o *Cloud) SetPreflightCompletedAt(v time.Time) {
	o.PreflightCompletedAt = &v
}

// GetBootstrappedAt returns the BootstrappedAt field value if set, zero value otherwise.
func (o *Cloud) GetBootstrappedAt() time.Time {
	if o == nil || o.BootstrappedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.BootstrappedAt
}

// GetBootstrappedAtOk returns a tuple with the BootstrappedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetBootstrappedAtOk() (*time.Time, bool) {
	if o == nil || o.BootstrappedAt == nil {
		return nil, false
	}
	return o.BootstrappedAt, true
}

// HasBootstrappedAt returns a boolean if a field has been set.
func (o *Cloud) HasBootstrappedAt() bool {
	if o != nil && o.BootstrappedAt != nil {
		return true
	}

	return false
}

// SetBootstrappedAt gets a reference to the given time.Time and assigns it to the BootstrappedAt field.
func (o *Cloud) SetBootstrappedAt(v time.Time) {
	o.BootstrappedAt = &v
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise.
func (o *Cloud) GetDeployedAt() time.Time {
	if o == nil || o.DeployedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeployedAt
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetDeployedAtOk() (*time.Time, bool) {
	if o == nil || o.DeployedAt == nil {
		return nil, false
	}
	return o.DeployedAt, true
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *Cloud) HasDeployedAt() bool {
	if o != nil && o.DeployedAt != nil {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given time.Time and assigns it to the DeployedAt field.
func (o *Cloud) SetDeployedAt(v time.Time) {
	o.DeployedAt = &v
}

// GetDestroyedAt returns the DestroyedAt field value if set, zero value otherwise.
func (o *Cloud) GetDestroyedAt() time.Time {
	if o == nil || o.DestroyedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DestroyedAt
}

// GetDestroyedAtOk returns a tuple with the DestroyedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetDestroyedAtOk() (*time.Time, bool) {
	if o == nil || o.DestroyedAt == nil {
		return nil, false
	}
	return o.DestroyedAt, true
}

// HasDestroyedAt returns a boolean if a field has been set.
func (o *Cloud) HasDestroyedAt() bool {
	if o != nil && o.DestroyedAt != nil {
		return true
	}

	return false
}

// SetDestroyedAt gets a reference to the given time.Time and assigns it to the DestroyedAt field.
func (o *Cloud) SetDestroyedAt(v time.Time) {
	o.DestroyedAt = &v
}

// GetProviderSettings returns the ProviderSettings field value
func (o *Cloud) GetProviderSettings() CloudProviderSettings {
	if o == nil {
		var ret CloudProviderSettings
		return ret
	}

	return o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetProviderSettingsOk() (*CloudProviderSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProviderSettings, true
}

// SetProviderSettings sets field value
func (o *Cloud) SetProviderSettings(v CloudProviderSettings) {
	o.ProviderSettings = v
}

// GetVersion returns the Version field value
func (o *Cloud) GetVersion() CloudVersion {
	if o == nil {
		var ret CloudVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetVersionOk() (*CloudVersion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Cloud) SetVersion(v CloudVersion) {
	o.Version = v
}

func (o Cloud) MarshalJSON() ([]byte, error) {
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
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.PreflightStartedAt != nil {
		toSerialize["preflightStartedAt"] = o.PreflightStartedAt
	}
	if o.PreflightFailedAt != nil {
		toSerialize["preflightFailedAt"] = o.PreflightFailedAt
	}
	if o.PreflightCompletedAt != nil {
		toSerialize["preflightCompletedAt"] = o.PreflightCompletedAt
	}
	if o.BootstrappedAt != nil {
		toSerialize["bootstrappedAt"] = o.BootstrappedAt
	}
	if o.DeployedAt != nil {
		toSerialize["deployedAt"] = o.DeployedAt
	}
	if o.DestroyedAt != nil {
		toSerialize["destroyedAt"] = o.DestroyedAt
	}
	if true {
		toSerialize["providerSettings"] = o.ProviderSettings
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCloud struct {
	value *Cloud
	isSet bool
}

func (v NullableCloud) Get() *Cloud {
	return v.value
}

func (v *NullableCloud) Set(val *Cloud) {
	v.value = val
	v.isSet = true
}

func (v NullableCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloud(val *Cloud) *NullableCloud {
	return &NullableCloud{value: val, isSet: true}
}

func (v NullableCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


