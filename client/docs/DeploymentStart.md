# DeploymentStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**BillingInfo** | **int32** |  | 

## Methods

### NewDeploymentStart

`func NewDeploymentStart(billingInfo int32, ) *DeploymentStart`

NewDeploymentStart instantiates a new DeploymentStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStartWithDefaults

`func NewDeploymentStartWithDefaults() *DeploymentStart`

NewDeploymentStartWithDefaults instantiates a new DeploymentStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *DeploymentStart) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *DeploymentStart) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *DeploymentStart) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *DeploymentStart) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetBillingInfo

`func (o *DeploymentStart) GetBillingInfo() int32`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *DeploymentStart) GetBillingInfoOk() (*int32, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *DeploymentStart) SetBillingInfo(v int32)`

SetBillingInfo sets BillingInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


