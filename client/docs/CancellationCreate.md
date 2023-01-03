# CancellationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **int32** | Unique ID of a device. Must be related to &#x60;service_id&#x60;. | 
**ServiceId** | **int32** | Unique ID of a service. Must be related to &#x60;device_id&#x60;. | 
**Reason** | **string** | My current service is being upgraded|Hardware/software/network or power issues|I think I have found a better deal|I have experienced tech support problems|I am consolidating my Hivelocity accounts|I am a reseller and my customer cancelled|I am moving to a different technology solution: Public Cloud|I am moving to a different technology solution: Managed Hosting Company|A product difference (Example: AWS) | 
**Comments** | Pointer to **string** | Details why the device is being cancelled | [optional] 

## Methods

### NewCancellationCreate

`func NewCancellationCreate(deviceId int32, serviceId int32, reason string, ) *CancellationCreate`

NewCancellationCreate instantiates a new CancellationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationCreateWithDefaults

`func NewCancellationCreateWithDefaults() *CancellationCreate`

NewCancellationCreateWithDefaults instantiates a new CancellationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *CancellationCreate) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CancellationCreate) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CancellationCreate) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.


### GetServiceId

`func (o *CancellationCreate) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CancellationCreate) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CancellationCreate) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetReason

`func (o *CancellationCreate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancellationCreate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancellationCreate) SetReason(v string)`

SetReason sets Reason field to given value.


### GetComments

`func (o *CancellationCreate) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CancellationCreate) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CancellationCreate) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CancellationCreate) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


