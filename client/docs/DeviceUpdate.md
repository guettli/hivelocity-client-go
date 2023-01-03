# DeviceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** | A FQDN for the device. For example: &#x60;example.hivelocity.net&#x60;. | [optional] 

## Methods

### NewDeviceUpdate

`func NewDeviceUpdate() *DeviceUpdate`

NewDeviceUpdate instantiates a new DeviceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateWithDefaults

`func NewDeviceUpdateWithDefaults() *DeviceUpdate`

NewDeviceUpdateWithDefaults instantiates a new DeviceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *DeviceUpdate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeviceUpdate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeviceUpdate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeviceUpdate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


