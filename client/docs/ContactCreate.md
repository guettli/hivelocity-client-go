# ContactCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int32** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**FullName** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] [default to "c879b075e42bb6b2364a6557f9a5f208"]
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewContactCreate

`func NewContactCreate(active int32, email string, fullName string, ) *ContactCreate`

NewContactCreate instantiates a new ContactCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateWithDefaults

`func NewContactCreateWithDefaults() *ContactCreate`

NewContactCreateWithDefaults instantiates a new ContactCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ContactCreate) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ContactCreate) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ContactCreate) SetActive(v int32)`

SetActive sets Active field to given value.


### GetPhone

`func (o *ContactCreate) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactCreate) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactCreate) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ContactCreate) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *ContactCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *ContactCreate) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ContactCreate) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ContactCreate) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetPassword

`func (o *ContactCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContactCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContactCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ContactCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDescription

`func (o *ContactCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


