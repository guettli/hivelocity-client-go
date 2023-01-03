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

// DeviceUpdate struct for DeviceUpdate
type DeviceUpdate struct {
	Name *string `json:"name,omitempty"`
	// A FQDN for the device. For example: `example.hivelocity.net`.
	Hostname *string `json:"hostname,omitempty"`
}

// NewDeviceUpdate instantiates a new DeviceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUpdate() *DeviceUpdate {
	this := DeviceUpdate{}
	return &this
}

// NewDeviceUpdateWithDefaults instantiates a new DeviceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUpdateWithDefaults() *DeviceUpdate {
	this := DeviceUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceUpdate) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeviceUpdate) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeviceUpdate) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeviceUpdate) SetHostname(v string) {
	o.Hostname = &v
}

func (o DeviceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceUpdate struct {
	value *DeviceUpdate
	isSet bool
}

func (v NullableDeviceUpdate) Get() *DeviceUpdate {
	return v.value
}

func (v *NullableDeviceUpdate) Set(val *DeviceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpdate(val *DeviceUpdate) *NullableDeviceUpdate {
	return &NullableDeviceUpdate{value: val, isSet: true}
}

func (v NullableDeviceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
