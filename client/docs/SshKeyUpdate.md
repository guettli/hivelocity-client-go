# SshKeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewSshKeyUpdate

`func NewSshKeyUpdate() *SshKeyUpdate`

NewSshKeyUpdate instantiates a new SshKeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyUpdateWithDefaults

`func NewSshKeyUpdateWithDefaults() *SshKeyUpdate`

NewSshKeyUpdateWithDefaults instantiates a new SshKeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshKeyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshKeyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *SshKeyUpdate) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshKeyUpdate) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshKeyUpdate) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SshKeyUpdate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


