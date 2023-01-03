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

// UpdateDevicesIPMIThresholds struct for UpdateDevicesIPMIThresholds
type UpdateDevicesIPMIThresholds struct {
	// Alert thresholds for IPMI sensor values.
	Thresholds map[string]interface{} `json:"thresholds"`
	// List of unique device IDs.
	DeviceIds []int32 `json:"device_ids"`
}

// NewUpdateDevicesIPMIThresholds instantiates a new UpdateDevicesIPMIThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDevicesIPMIThresholds(thresholds map[string]interface{}, deviceIds []int32) *UpdateDevicesIPMIThresholds {
	this := UpdateDevicesIPMIThresholds{}
	this.Thresholds = thresholds
	this.DeviceIds = deviceIds
	return &this
}

// NewUpdateDevicesIPMIThresholdsWithDefaults instantiates a new UpdateDevicesIPMIThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDevicesIPMIThresholdsWithDefaults() *UpdateDevicesIPMIThresholds {
	this := UpdateDevicesIPMIThresholds{}
	return &this
}

// GetThresholds returns the Thresholds field value
func (o *UpdateDevicesIPMIThresholds) GetThresholds() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value
// and a boolean to check if the value has been set.
func (o *UpdateDevicesIPMIThresholds) GetThresholdsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Thresholds, true
}

// SetThresholds sets field value
func (o *UpdateDevicesIPMIThresholds) SetThresholds(v map[string]interface{}) {
	o.Thresholds = v
}

// GetDeviceIds returns the DeviceIds field value
func (o *UpdateDevicesIPMIThresholds) GetDeviceIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *UpdateDevicesIPMIThresholds) GetDeviceIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *UpdateDevicesIPMIThresholds) SetDeviceIds(v []int32) {
	o.DeviceIds = v
}

func (o UpdateDevicesIPMIThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["thresholds"] = o.Thresholds
	}
	if true {
		toSerialize["device_ids"] = o.DeviceIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDevicesIPMIThresholds struct {
	value *UpdateDevicesIPMIThresholds
	isSet bool
}

func (v NullableUpdateDevicesIPMIThresholds) Get() *UpdateDevicesIPMIThresholds {
	return v.value
}

func (v *NullableUpdateDevicesIPMIThresholds) Set(val *UpdateDevicesIPMIThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDevicesIPMIThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDevicesIPMIThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDevicesIPMIThresholds(val *UpdateDevicesIPMIThresholds) *NullableUpdateDevicesIPMIThresholds {
	return &NullableUpdateDevicesIPMIThresholds{value: val, isSet: true}
}

func (v NullableUpdateDevicesIPMIThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDevicesIPMIThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
