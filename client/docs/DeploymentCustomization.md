# DeploymentCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** |  | [optional] 
**BillingPeriod** | Pointer to **string** | must be one of [&#39;monthly&#39;, &#39;quarterly&#39;, &#39;semi-annually&#39;, &#39;annually&#39;, &#39;biennial&#39;, &#39;triennial&#39;, &#39;hourly&#39;] | [optional] [default to "monthly"]
**LocationCode** | Pointer to **string** |  | [optional] [default to ""]
**CustomIPXEScriptURL** | Pointer to **string** | URL to download custom iPXE script if not supplying script in entirety | [optional] 
**PublicSshKeyId** | Pointer to **int32** | ID of SSH Key to use for deployment | [optional] 
**AdditionalNotes** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to **[]int32** |  | [optional] 
**ForceDeviceIds** | Pointer to **[]int32** | Either deploy these Device IDs or fail | [optional] 
**Hostnames** | **[]string** |  | 
**OperatingSystem** | **string** | Operating System&#39;s Name or ID | 
**IgnitionIds** | Pointer to **[]int32** | Specify Ignition file ID for CoreOS/Flatcar provisions | [optional] 
**ProductId** | **int32** |  | 
**CustomIPXEScriptContents** | Pointer to **string** | Contents of iPXE script if not supplying URL | [optional] 

## Methods

### NewDeploymentCustomization

`func NewDeploymentCustomization(hostnames []string, operatingSystem string, productId int32, ) *DeploymentCustomization`

NewDeploymentCustomization instantiates a new DeploymentCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCustomizationWithDefaults

`func NewDeploymentCustomizationWithDefaults() *DeploymentCustomization`

NewDeploymentCustomizationWithDefaults instantiates a new DeploymentCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *DeploymentCustomization) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DeploymentCustomization) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DeploymentCustomization) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DeploymentCustomization) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DeploymentCustomization) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DeploymentCustomization) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DeploymentCustomization) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DeploymentCustomization) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetLocationCode

`func (o *DeploymentCustomization) GetLocationCode() string`

GetLocationCode returns the LocationCode field if non-nil, zero value otherwise.

### GetLocationCodeOk

`func (o *DeploymentCustomization) GetLocationCodeOk() (*string, bool)`

GetLocationCodeOk returns a tuple with the LocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCode

`func (o *DeploymentCustomization) SetLocationCode(v string)`

SetLocationCode sets LocationCode field to given value.

### HasLocationCode

`func (o *DeploymentCustomization) HasLocationCode() bool`

HasLocationCode returns a boolean if a field has been set.

### GetCustomIPXEScriptURL

`func (o *DeploymentCustomization) GetCustomIPXEScriptURL() string`

GetCustomIPXEScriptURL returns the CustomIPXEScriptURL field if non-nil, zero value otherwise.

### GetCustomIPXEScriptURLOk

`func (o *DeploymentCustomization) GetCustomIPXEScriptURLOk() (*string, bool)`

GetCustomIPXEScriptURLOk returns a tuple with the CustomIPXEScriptURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIPXEScriptURL

`func (o *DeploymentCustomization) SetCustomIPXEScriptURL(v string)`

SetCustomIPXEScriptURL sets CustomIPXEScriptURL field to given value.

### HasCustomIPXEScriptURL

`func (o *DeploymentCustomization) HasCustomIPXEScriptURL() bool`

HasCustomIPXEScriptURL returns a boolean if a field has been set.

### GetPublicSshKeyId

`func (o *DeploymentCustomization) GetPublicSshKeyId() int32`

GetPublicSshKeyId returns the PublicSshKeyId field if non-nil, zero value otherwise.

### GetPublicSshKeyIdOk

`func (o *DeploymentCustomization) GetPublicSshKeyIdOk() (*int32, bool)`

GetPublicSshKeyIdOk returns a tuple with the PublicSshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSshKeyId

`func (o *DeploymentCustomization) SetPublicSshKeyId(v int32)`

SetPublicSshKeyId sets PublicSshKeyId field to given value.

### HasPublicSshKeyId

`func (o *DeploymentCustomization) HasPublicSshKeyId() bool`

HasPublicSshKeyId returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *DeploymentCustomization) GetAdditionalNotes() []string`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *DeploymentCustomization) GetAdditionalNotesOk() (*[]string, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *DeploymentCustomization) SetAdditionalNotes(v []string)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *DeploymentCustomization) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetOptions

`func (o *DeploymentCustomization) GetOptions() []int32`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DeploymentCustomization) GetOptionsOk() (*[]int32, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DeploymentCustomization) SetOptions(v []int32)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DeploymentCustomization) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetForceDeviceIds

`func (o *DeploymentCustomization) GetForceDeviceIds() []int32`

GetForceDeviceIds returns the ForceDeviceIds field if non-nil, zero value otherwise.

### GetForceDeviceIdsOk

`func (o *DeploymentCustomization) GetForceDeviceIdsOk() (*[]int32, bool)`

GetForceDeviceIdsOk returns a tuple with the ForceDeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeviceIds

`func (o *DeploymentCustomization) SetForceDeviceIds(v []int32)`

SetForceDeviceIds sets ForceDeviceIds field to given value.

### HasForceDeviceIds

`func (o *DeploymentCustomization) HasForceDeviceIds() bool`

HasForceDeviceIds returns a boolean if a field has been set.

### GetHostnames

`func (o *DeploymentCustomization) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *DeploymentCustomization) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *DeploymentCustomization) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetOperatingSystem

`func (o *DeploymentCustomization) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *DeploymentCustomization) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *DeploymentCustomization) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetIgnitionIds

`func (o *DeploymentCustomization) GetIgnitionIds() []int32`

GetIgnitionIds returns the IgnitionIds field if non-nil, zero value otherwise.

### GetIgnitionIdsOk

`func (o *DeploymentCustomization) GetIgnitionIdsOk() (*[]int32, bool)`

GetIgnitionIdsOk returns a tuple with the IgnitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnitionIds

`func (o *DeploymentCustomization) SetIgnitionIds(v []int32)`

SetIgnitionIds sets IgnitionIds field to given value.

### HasIgnitionIds

`func (o *DeploymentCustomization) HasIgnitionIds() bool`

HasIgnitionIds returns a boolean if a field has been set.

### GetProductId

`func (o *DeploymentCustomization) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DeploymentCustomization) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DeploymentCustomization) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetCustomIPXEScriptContents

`func (o *DeploymentCustomization) GetCustomIPXEScriptContents() string`

GetCustomIPXEScriptContents returns the CustomIPXEScriptContents field if non-nil, zero value otherwise.

### GetCustomIPXEScriptContentsOk

`func (o *DeploymentCustomization) GetCustomIPXEScriptContentsOk() (*string, bool)`

GetCustomIPXEScriptContentsOk returns a tuple with the CustomIPXEScriptContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIPXEScriptContents

`func (o *DeploymentCustomization) SetCustomIPXEScriptContents(v string)`

SetCustomIPXEScriptContents sets CustomIPXEScriptContents field to given value.

### HasCustomIPXEScriptContents

`func (o *DeploymentCustomization) HasCustomIPXEScriptContents() bool`

HasCustomIPXEScriptContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


