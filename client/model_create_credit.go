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

// CreateCredit struct for CreateCredit
type CreateCredit struct {
	Amount        *float32 `json:"amount,omitempty"`
	BillingInfoId *int32   `json:"billingInfoId,omitempty"`
}

// NewCreateCredit instantiates a new CreateCredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredit() *CreateCredit {
	this := CreateCredit{}
	return &this
}

// NewCreateCreditWithDefaults instantiates a new CreateCredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCreditWithDefaults() *CreateCredit {
	this := CreateCredit{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateCredit) GetAmount() float32 {
	if o == nil || isNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredit) GetAmountOk() (*float32, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateCredit) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CreateCredit) SetAmount(v float32) {
	o.Amount = &v
}

// GetBillingInfoId returns the BillingInfoId field value if set, zero value otherwise.
func (o *CreateCredit) GetBillingInfoId() int32 {
	if o == nil || isNil(o.BillingInfoId) {
		var ret int32
		return ret
	}
	return *o.BillingInfoId
}

// GetBillingInfoIdOk returns a tuple with the BillingInfoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredit) GetBillingInfoIdOk() (*int32, bool) {
	if o == nil || isNil(o.BillingInfoId) {
		return nil, false
	}
	return o.BillingInfoId, true
}

// HasBillingInfoId returns a boolean if a field has been set.
func (o *CreateCredit) HasBillingInfoId() bool {
	if o != nil && !isNil(o.BillingInfoId) {
		return true
	}

	return false
}

// SetBillingInfoId gets a reference to the given int32 and assigns it to the BillingInfoId field.
func (o *CreateCredit) SetBillingInfoId(v int32) {
	o.BillingInfoId = &v
}

func (o CreateCredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.BillingInfoId) {
		toSerialize["billingInfoId"] = o.BillingInfoId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCredit struct {
	value *CreateCredit
	isSet bool
}

func (v NullableCreateCredit) Get() *CreateCredit {
	return v.value
}

func (v *NullableCreateCredit) Set(val *CreateCredit) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCredit(val *CreateCredit) *NullableCreateCredit {
	return &NullableCreateCredit{value: val, isSet: true}
}

func (v NullableCreateCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
