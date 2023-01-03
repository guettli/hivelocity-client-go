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

// DeviceIPMIWhitelistIP struct for DeviceIPMIWhitelistIP
type DeviceIPMIWhitelistIP struct {
	CustIp string `json:"custIp"`
}

// NewDeviceIPMIWhitelistIP instantiates a new DeviceIPMIWhitelistIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIPMIWhitelistIP(custIp string) *DeviceIPMIWhitelistIP {
	this := DeviceIPMIWhitelistIP{}
	this.CustIp = custIp
	return &this
}

// NewDeviceIPMIWhitelistIPWithDefaults instantiates a new DeviceIPMIWhitelistIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIPMIWhitelistIPWithDefaults() *DeviceIPMIWhitelistIP {
	this := DeviceIPMIWhitelistIP{}
	return &this
}

// GetCustIp returns the CustIp field value
func (o *DeviceIPMIWhitelistIP) GetCustIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustIp
}

// GetCustIpOk returns a tuple with the CustIp field value
// and a boolean to check if the value has been set.
func (o *DeviceIPMIWhitelistIP) GetCustIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustIp, true
}

// SetCustIp sets field value
func (o *DeviceIPMIWhitelistIP) SetCustIp(v string) {
	o.CustIp = v
}

func (o DeviceIPMIWhitelistIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["custIp"] = o.CustIp
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceIPMIWhitelistIP struct {
	value *DeviceIPMIWhitelistIP
	isSet bool
}

func (v NullableDeviceIPMIWhitelistIP) Get() *DeviceIPMIWhitelistIP {
	return v.value
}

func (v *NullableDeviceIPMIWhitelistIP) Set(val *DeviceIPMIWhitelistIP) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIPMIWhitelistIP) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIPMIWhitelistIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIPMIWhitelistIP(val *DeviceIPMIWhitelistIP) *NullableDeviceIPMIWhitelistIP {
	return &NullableDeviceIPMIWhitelistIP{value: val, isSet: true}
}

func (v NullableDeviceIPMIWhitelistIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIPMIWhitelistIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
