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

// PasswordResponse struct for PasswordResponse
type PasswordResponse struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

// NewPasswordResponse instantiates a new PasswordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordResponse(password string, token string) *PasswordResponse {
	this := PasswordResponse{}
	this.Password = password
	this.Token = token
	return &this
}

// NewPasswordResponseWithDefaults instantiates a new PasswordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordResponseWithDefaults() *PasswordResponse {
	this := PasswordResponse{}
	return &this
}

// GetPassword returns the Password field value
func (o *PasswordResponse) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PasswordResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PasswordResponse) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *PasswordResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PasswordResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PasswordResponse) SetToken(v string) {
	o.Token = v
}

func (o PasswordResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordResponse struct {
	value *PasswordResponse
	isSet bool
}

func (v NullablePasswordResponse) Get() *PasswordResponse {
	return v.value
}

func (v *NullablePasswordResponse) Set(val *PasswordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordResponse(val *PasswordResponse) *NullablePasswordResponse {
	return &NullablePasswordResponse{value: val, isSet: true}
}

func (v NullablePasswordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
