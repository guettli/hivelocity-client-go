# WebhookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWebhookUpdate

`func NewWebhookUpdate() *WebhookUpdate`

NewWebhookUpdate instantiates a new WebhookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateWithDefaults

`func NewWebhookUpdateWithDefaults() *WebhookUpdate`

NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvent

`func (o *WebhookUpdate) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookUpdate) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookUpdate) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookUpdate) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *WebhookUpdate) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookUpdate) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookUpdate) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookUpdate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


