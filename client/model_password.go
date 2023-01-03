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

// Password struct for Password
type Password struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// NewPassword instantiates a new Password object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassword(password string, newPassword string) *Password {
	this := Password{}
	this.Password = password
	this.NewPassword = newPassword
	return &this
}

// NewPasswordWithDefaults instantiates a new Password object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordWithDefaults() *Password {
	this := Password{}
	return &this
}

// GetPassword returns the Password field value
func (o *Password) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Password) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *Password) SetPassword(v string) {
	o.Password = v
}

// GetNewPassword returns the NewPassword field value
func (o *Password) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *Password) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *Password) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o Password) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["newPassword"] = o.NewPassword
	}
	return json.Marshal(toSerialize)
}

type NullablePassword struct {
	value *Password
	isSet bool
}

func (v NullablePassword) Get() *Password {
	return v.value
}

func (v *NullablePassword) Set(val *Password) {
	v.value = val
	v.isSet = true
}

func (v NullablePassword) IsSet() bool {
	return v.isSet
}

func (v *NullablePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassword(val *Password) *NullablePassword {
	return &NullablePassword{value: val, isSet: true}
}

func (v NullablePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
