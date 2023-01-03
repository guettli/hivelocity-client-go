# OperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostInstall** | Pointer to **bool** |  | [optional] 
**SshKey** | Pointer to **bool** |  | [optional] 
**CloudInit** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperatingSystem

`func NewOperatingSystem() *OperatingSystem`

NewOperatingSystem instantiates a new OperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemWithDefaults

`func NewOperatingSystemWithDefaults() *OperatingSystem`

NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostInstall

`func (o *OperatingSystem) GetPostInstall() bool`

GetPostInstall returns the PostInstall field if non-nil, zero value otherwise.

### GetPostInstallOk

`func (o *OperatingSystem) GetPostInstallOk() (*bool, bool)`

GetPostInstallOk returns a tuple with the PostInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInstall

`func (o *OperatingSystem) SetPostInstall(v bool)`

SetPostInstall sets PostInstall field to given value.

### HasPostInstall

`func (o *OperatingSystem) HasPostInstall() bool`

HasPostInstall returns a boolean if a field has been set.

### GetSshKey

`func (o *OperatingSystem) GetSshKey() bool`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *OperatingSystem) GetSshKeyOk() (*bool, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *OperatingSystem) SetSshKey(v bool)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *OperatingSystem) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetCloudInit

`func (o *OperatingSystem) GetCloudInit() bool`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *OperatingSystem) GetCloudInitOk() (*bool, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *OperatingSystem) SetCloudInit(v bool)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *OperatingSystem) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


