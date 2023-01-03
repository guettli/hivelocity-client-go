# PublicApiTokenLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to [**PublicApiTokenIp**](PublicApiTokenIp.md) |  | [optional] 

## Methods

### NewPublicApiTokenLoad

`func NewPublicApiTokenLoad() *PublicApiTokenLoad`

NewPublicApiTokenLoad instantiates a new PublicApiTokenLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicApiTokenLoadWithDefaults

`func NewPublicApiTokenLoadWithDefaults() *PublicApiTokenLoad`

NewPublicApiTokenLoadWithDefaults instantiates a new PublicApiTokenLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicApiTokenLoad) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicApiTokenLoad) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicApiTokenLoad) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicApiTokenLoad) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpAddresses

`func (o *PublicApiTokenLoad) GetIpAddresses() PublicApiTokenIp`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *PublicApiTokenLoad) GetIpAddressesOk() (*PublicApiTokenIp, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *PublicApiTokenLoad) SetIpAddresses(v PublicApiTokenIp)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *PublicApiTokenLoad) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


