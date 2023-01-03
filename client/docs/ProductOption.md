# ProductOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int32** | Unique product id. | [optional] 
**Options** | Pointer to **map[string]interface{}** | List of available options. | [optional] 
**Pricing** | Pointer to **map[string]interface{}** | *DEPRECATED* | [optional] 

## Methods

### NewProductOption

`func NewProductOption() *ProductOption`

NewProductOption instantiates a new ProductOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOptionWithDefaults

`func NewProductOptionWithDefaults() *ProductOption`

NewProductOptionWithDefaults instantiates a new ProductOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductOption) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductOption) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductOption) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductOption) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetOptions

`func (o *ProductOption) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductOption) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductOption) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProductOption) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPricing

`func (o *ProductOption) GetPricing() map[string]interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ProductOption) GetPricingOk() (*map[string]interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ProductOption) SetPricing(v map[string]interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *ProductOption) HasPricing() bool`

HasPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


