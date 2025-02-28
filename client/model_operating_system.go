/*
Hivelocity API

Interact with Hivelocity

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	PostInstall *bool `json:"postInstall,omitempty"`
	SshKey      *bool `json:"sshKey,omitempty"`
	CloudInit   *bool `json:"cloudInit,omitempty"`
}

// NewOperatingSystem instantiates a new OperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystem() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemWithDefaults() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// GetPostInstall returns the PostInstall field value if set, zero value otherwise.
func (o *OperatingSystem) GetPostInstall() bool {
	if o == nil || isNil(o.PostInstall) {
		var ret bool
		return ret
	}
	return *o.PostInstall
}

// GetPostInstallOk returns a tuple with the PostInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetPostInstallOk() (*bool, bool) {
	if o == nil || isNil(o.PostInstall) {
		return nil, false
	}
	return o.PostInstall, true
}

// HasPostInstall returns a boolean if a field has been set.
func (o *OperatingSystem) HasPostInstall() bool {
	if o != nil && !isNil(o.PostInstall) {
		return true
	}

	return false
}

// SetPostInstall gets a reference to the given bool and assigns it to the PostInstall field.
func (o *OperatingSystem) SetPostInstall(v bool) {
	o.PostInstall = &v
}

// GetSshKey returns the SshKey field value if set, zero value otherwise.
func (o *OperatingSystem) GetSshKey() bool {
	if o == nil || isNil(o.SshKey) {
		var ret bool
		return ret
	}
	return *o.SshKey
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetSshKeyOk() (*bool, bool) {
	if o == nil || isNil(o.SshKey) {
		return nil, false
	}
	return o.SshKey, true
}

// HasSshKey returns a boolean if a field has been set.
func (o *OperatingSystem) HasSshKey() bool {
	if o != nil && !isNil(o.SshKey) {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given bool and assigns it to the SshKey field.
func (o *OperatingSystem) SetSshKey(v bool) {
	o.SshKey = &v
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise.
func (o *OperatingSystem) GetCloudInit() bool {
	if o == nil || isNil(o.CloudInit) {
		var ret bool
		return ret
	}
	return *o.CloudInit
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetCloudInitOk() (*bool, bool) {
	if o == nil || isNil(o.CloudInit) {
		return nil, false
	}
	return o.CloudInit, true
}

// HasCloudInit returns a boolean if a field has been set.
func (o *OperatingSystem) HasCloudInit() bool {
	if o != nil && !isNil(o.CloudInit) {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given bool and assigns it to the CloudInit field.
func (o *OperatingSystem) SetCloudInit(v bool) {
	o.CloudInit = &v
}

func (o OperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PostInstall) {
		toSerialize["postInstall"] = o.PostInstall
	}
	if !isNil(o.SshKey) {
		toSerialize["sshKey"] = o.SshKey
	}
	if !isNil(o.CloudInit) {
		toSerialize["cloudInit"] = o.CloudInit
	}
	return json.Marshal(toSerialize)
}

type NullableOperatingSystem struct {
	value *OperatingSystem
	isSet bool
}

func (v NullableOperatingSystem) Get() *OperatingSystem {
	return v.value
}

func (v *NullableOperatingSystem) Set(val *OperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystem(val *OperatingSystem) *NullableOperatingSystem {
	return &NullableOperatingSystem{value: val, isSet: true}
}

func (v NullableOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
