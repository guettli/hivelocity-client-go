# Stock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCenter** | Pointer to **string** | A facility code. For example &#x60;NYC1&#x60;. | [optional] 
**HourlyLocationPremium** | Pointer to **float32** | Additional hourly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**ProductDrive** | Pointer to **string** | Human readable drive specs. Can include multiple drives. | [optional] 
**Edge** | Pointer to **bool** | true|false if edge site. | [optional] 
**Stock** | Pointer to **string** | available|limited|unavailable. Limited denotes minimal stock. | [optional] 
**ProductBiennialPrice** | Pointer to **float32** | Price per 2 year period (24 months). | [optional] 
**BiennialLocationPremium** | Pointer to **float32** | Additional biennial fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**MonthlyLocationPremium** | Pointer to **float32** | Additional monthly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**ProductHourlyPrice** | Pointer to **float32** | Price per hour. | [optional] 
**TriennialLocationPremium** | Pointer to **float32** | Additional triennial fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**Core** | Pointer to **bool** | true|false if core site. | [optional] 
**AnnuallyLocationPremium** | Pointer to **float32** | Additional annual fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**ProductOriginalPrice** | Pointer to **float32** | Retail price of product. Use with &#x60;product_on_sale&#x60;. | [optional] 
**ProductAnnuallyPrice** | Pointer to **float32** | Price per year (12 months). | [optional] 
**SemiAnnuallyLocationPremium** | Pointer to **float32** | Additional semi-annual fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**ProductSemiAnnuallyPrice** | Pointer to **float32** | Price per half year (6 months). | [optional] 
**ProductBandwidth** | Pointer to **string** | Human readable networking specs in the format: Free Outbound Transfer / NIC Size | [optional] 
**ProductId** | Pointer to **int32** | The unique ID of this product. | [optional] 
**ProductMonthlyPrice** | Pointer to **float32** | Price per month. | [optional] 
**ProductTriennialPrice** | Pointer to **float32** | Price per 3 year period (36 months). | [optional] 
**ProductCpu** | Pointer to **string** | Human readable CPU specs. | [optional] 
**ProductCpuCores** | Pointer to **string** | Human readable CPU core and thread info in HTML format. | [optional] 
**ProductGpu** | Pointer to **string** | Human readable GPU specs. | [optional] 
**QuarterlyLocationPremium** | Pointer to **float32** | Additional quarterly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] 
**ProductName** | Pointer to **string** | The unique name of this product. | [optional] 
**ProductQuarterlyPrice** | Pointer to **float32** | Price per quarter (3 months). | [optional] 
**ProductOnSale** | Pointer to **bool** | true|false. Use with &#x60;product_original_price&#x60;. | [optional] 
**ProductMemory** | Pointer to **string** | Human readable memory specs. | [optional] 
**ProcessorInfo** | Pointer to **map[string]interface{}** | JSON CPU info for cores, threads, sockets, and vCPUs. | [optional] 
**ProductDisabledBillingPeriods** | Pointer to **[]string** | Orders for the periods in the array will fail. Possible values: hourly|monthly|quarterly|semi-annually|biennial|triennial | [optional] 
**ProductDisplayPrice** | Pointer to **float32** | *DEPRECATED*. | [optional] 

## Methods

### NewStock

`func NewStock() *Stock`

NewStock instantiates a new Stock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockWithDefaults

`func NewStockWithDefaults() *Stock`

NewStockWithDefaults instantiates a new Stock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCenter

`func (o *Stock) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *Stock) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *Stock) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *Stock) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetHourlyLocationPremium

`func (o *Stock) GetHourlyLocationPremium() float32`

GetHourlyLocationPremium returns the HourlyLocationPremium field if non-nil, zero value otherwise.

### GetHourlyLocationPremiumOk

`func (o *Stock) GetHourlyLocationPremiumOk() (*float32, bool)`

GetHourlyLocationPremiumOk returns a tuple with the HourlyLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyLocationPremium

`func (o *Stock) SetHourlyLocationPremium(v float32)`

SetHourlyLocationPremium sets HourlyLocationPremium field to given value.

### HasHourlyLocationPremium

`func (o *Stock) HasHourlyLocationPremium() bool`

HasHourlyLocationPremium returns a boolean if a field has been set.

### GetProductDrive

`func (o *Stock) GetProductDrive() string`

GetProductDrive returns the ProductDrive field if non-nil, zero value otherwise.

### GetProductDriveOk

`func (o *Stock) GetProductDriveOk() (*string, bool)`

GetProductDriveOk returns a tuple with the ProductDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDrive

`func (o *Stock) SetProductDrive(v string)`

SetProductDrive sets ProductDrive field to given value.

### HasProductDrive

`func (o *Stock) HasProductDrive() bool`

HasProductDrive returns a boolean if a field has been set.

### GetEdge

`func (o *Stock) GetEdge() bool`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *Stock) GetEdgeOk() (*bool, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *Stock) SetEdge(v bool)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *Stock) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### GetStock

`func (o *Stock) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Stock) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Stock) SetStock(v string)`

SetStock sets Stock field to given value.

### HasStock

`func (o *Stock) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetProductBiennialPrice

`func (o *Stock) GetProductBiennialPrice() float32`

GetProductBiennialPrice returns the ProductBiennialPrice field if non-nil, zero value otherwise.

### GetProductBiennialPriceOk

`func (o *Stock) GetProductBiennialPriceOk() (*float32, bool)`

GetProductBiennialPriceOk returns a tuple with the ProductBiennialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBiennialPrice

`func (o *Stock) SetProductBiennialPrice(v float32)`

SetProductBiennialPrice sets ProductBiennialPrice field to given value.

### HasProductBiennialPrice

`func (o *Stock) HasProductBiennialPrice() bool`

HasProductBiennialPrice returns a boolean if a field has been set.

### GetBiennialLocationPremium

`func (o *Stock) GetBiennialLocationPremium() float32`

GetBiennialLocationPremium returns the BiennialLocationPremium field if non-nil, zero value otherwise.

### GetBiennialLocationPremiumOk

`func (o *Stock) GetBiennialLocationPremiumOk() (*float32, bool)`

GetBiennialLocationPremiumOk returns a tuple with the BiennialLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiennialLocationPremium

`func (o *Stock) SetBiennialLocationPremium(v float32)`

SetBiennialLocationPremium sets BiennialLocationPremium field to given value.

### HasBiennialLocationPremium

`func (o *Stock) HasBiennialLocationPremium() bool`

HasBiennialLocationPremium returns a boolean if a field has been set.

### GetMonthlyLocationPremium

`func (o *Stock) GetMonthlyLocationPremium() float32`

GetMonthlyLocationPremium returns the MonthlyLocationPremium field if non-nil, zero value otherwise.

### GetMonthlyLocationPremiumOk

`func (o *Stock) GetMonthlyLocationPremiumOk() (*float32, bool)`

GetMonthlyLocationPremiumOk returns a tuple with the MonthlyLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLocationPremium

`func (o *Stock) SetMonthlyLocationPremium(v float32)`

SetMonthlyLocationPremium sets MonthlyLocationPremium field to given value.

### HasMonthlyLocationPremium

`func (o *Stock) HasMonthlyLocationPremium() bool`

HasMonthlyLocationPremium returns a boolean if a field has been set.

### GetProductHourlyPrice

`func (o *Stock) GetProductHourlyPrice() float32`

GetProductHourlyPrice returns the ProductHourlyPrice field if non-nil, zero value otherwise.

### GetProductHourlyPriceOk

`func (o *Stock) GetProductHourlyPriceOk() (*float32, bool)`

GetProductHourlyPriceOk returns a tuple with the ProductHourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductHourlyPrice

`func (o *Stock) SetProductHourlyPrice(v float32)`

SetProductHourlyPrice sets ProductHourlyPrice field to given value.

### HasProductHourlyPrice

`func (o *Stock) HasProductHourlyPrice() bool`

HasProductHourlyPrice returns a boolean if a field has been set.

### GetTriennialLocationPremium

`func (o *Stock) GetTriennialLocationPremium() float32`

GetTriennialLocationPremium returns the TriennialLocationPremium field if non-nil, zero value otherwise.

### GetTriennialLocationPremiumOk

`func (o *Stock) GetTriennialLocationPremiumOk() (*float32, bool)`

GetTriennialLocationPremiumOk returns a tuple with the TriennialLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriennialLocationPremium

`func (o *Stock) SetTriennialLocationPremium(v float32)`

SetTriennialLocationPremium sets TriennialLocationPremium field to given value.

### HasTriennialLocationPremium

`func (o *Stock) HasTriennialLocationPremium() bool`

HasTriennialLocationPremium returns a boolean if a field has been set.

### GetCore

`func (o *Stock) GetCore() bool`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *Stock) GetCoreOk() (*bool, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *Stock) SetCore(v bool)`

SetCore sets Core field to given value.

### HasCore

`func (o *Stock) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetAnnuallyLocationPremium

`func (o *Stock) GetAnnuallyLocationPremium() float32`

GetAnnuallyLocationPremium returns the AnnuallyLocationPremium field if non-nil, zero value otherwise.

### GetAnnuallyLocationPremiumOk

`func (o *Stock) GetAnnuallyLocationPremiumOk() (*float32, bool)`

GetAnnuallyLocationPremiumOk returns a tuple with the AnnuallyLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuallyLocationPremium

`func (o *Stock) SetAnnuallyLocationPremium(v float32)`

SetAnnuallyLocationPremium sets AnnuallyLocationPremium field to given value.

### HasAnnuallyLocationPremium

`func (o *Stock) HasAnnuallyLocationPremium() bool`

HasAnnuallyLocationPremium returns a boolean if a field has been set.

### GetProductOriginalPrice

`func (o *Stock) GetProductOriginalPrice() float32`

GetProductOriginalPrice returns the ProductOriginalPrice field if non-nil, zero value otherwise.

### GetProductOriginalPriceOk

`func (o *Stock) GetProductOriginalPriceOk() (*float32, bool)`

GetProductOriginalPriceOk returns a tuple with the ProductOriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOriginalPrice

`func (o *Stock) SetProductOriginalPrice(v float32)`

SetProductOriginalPrice sets ProductOriginalPrice field to given value.

### HasProductOriginalPrice

`func (o *Stock) HasProductOriginalPrice() bool`

HasProductOriginalPrice returns a boolean if a field has been set.

### GetProductAnnuallyPrice

`func (o *Stock) GetProductAnnuallyPrice() float32`

GetProductAnnuallyPrice returns the ProductAnnuallyPrice field if non-nil, zero value otherwise.

### GetProductAnnuallyPriceOk

`func (o *Stock) GetProductAnnuallyPriceOk() (*float32, bool)`

GetProductAnnuallyPriceOk returns a tuple with the ProductAnnuallyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAnnuallyPrice

`func (o *Stock) SetProductAnnuallyPrice(v float32)`

SetProductAnnuallyPrice sets ProductAnnuallyPrice field to given value.

### HasProductAnnuallyPrice

`func (o *Stock) HasProductAnnuallyPrice() bool`

HasProductAnnuallyPrice returns a boolean if a field has been set.

### GetSemiAnnuallyLocationPremium

`func (o *Stock) GetSemiAnnuallyLocationPremium() float32`

GetSemiAnnuallyLocationPremium returns the SemiAnnuallyLocationPremium field if non-nil, zero value otherwise.

### GetSemiAnnuallyLocationPremiumOk

`func (o *Stock) GetSemiAnnuallyLocationPremiumOk() (*float32, bool)`

GetSemiAnnuallyLocationPremiumOk returns a tuple with the SemiAnnuallyLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiAnnuallyLocationPremium

`func (o *Stock) SetSemiAnnuallyLocationPremium(v float32)`

SetSemiAnnuallyLocationPremium sets SemiAnnuallyLocationPremium field to given value.

### HasSemiAnnuallyLocationPremium

`func (o *Stock) HasSemiAnnuallyLocationPremium() bool`

HasSemiAnnuallyLocationPremium returns a boolean if a field has been set.

### GetProductSemiAnnuallyPrice

`func (o *Stock) GetProductSemiAnnuallyPrice() float32`

GetProductSemiAnnuallyPrice returns the ProductSemiAnnuallyPrice field if non-nil, zero value otherwise.

### GetProductSemiAnnuallyPriceOk

`func (o *Stock) GetProductSemiAnnuallyPriceOk() (*float32, bool)`

GetProductSemiAnnuallyPriceOk returns a tuple with the ProductSemiAnnuallyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSemiAnnuallyPrice

`func (o *Stock) SetProductSemiAnnuallyPrice(v float32)`

SetProductSemiAnnuallyPrice sets ProductSemiAnnuallyPrice field to given value.

### HasProductSemiAnnuallyPrice

`func (o *Stock) HasProductSemiAnnuallyPrice() bool`

HasProductSemiAnnuallyPrice returns a boolean if a field has been set.

### GetProductBandwidth

`func (o *Stock) GetProductBandwidth() string`

GetProductBandwidth returns the ProductBandwidth field if non-nil, zero value otherwise.

### GetProductBandwidthOk

`func (o *Stock) GetProductBandwidthOk() (*string, bool)`

GetProductBandwidthOk returns a tuple with the ProductBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBandwidth

`func (o *Stock) SetProductBandwidth(v string)`

SetProductBandwidth sets ProductBandwidth field to given value.

### HasProductBandwidth

`func (o *Stock) HasProductBandwidth() bool`

HasProductBandwidth returns a boolean if a field has been set.

### GetProductId

`func (o *Stock) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Stock) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Stock) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Stock) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductMonthlyPrice

`func (o *Stock) GetProductMonthlyPrice() float32`

GetProductMonthlyPrice returns the ProductMonthlyPrice field if non-nil, zero value otherwise.

### GetProductMonthlyPriceOk

`func (o *Stock) GetProductMonthlyPriceOk() (*float32, bool)`

GetProductMonthlyPriceOk returns a tuple with the ProductMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMonthlyPrice

`func (o *Stock) SetProductMonthlyPrice(v float32)`

SetProductMonthlyPrice sets ProductMonthlyPrice field to given value.

### HasProductMonthlyPrice

`func (o *Stock) HasProductMonthlyPrice() bool`

HasProductMonthlyPrice returns a boolean if a field has been set.

### GetProductTriennialPrice

`func (o *Stock) GetProductTriennialPrice() float32`

GetProductTriennialPrice returns the ProductTriennialPrice field if non-nil, zero value otherwise.

### GetProductTriennialPriceOk

`func (o *Stock) GetProductTriennialPriceOk() (*float32, bool)`

GetProductTriennialPriceOk returns a tuple with the ProductTriennialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTriennialPrice

`func (o *Stock) SetProductTriennialPrice(v float32)`

SetProductTriennialPrice sets ProductTriennialPrice field to given value.

### HasProductTriennialPrice

`func (o *Stock) HasProductTriennialPrice() bool`

HasProductTriennialPrice returns a boolean if a field has been set.

### GetProductCpu

`func (o *Stock) GetProductCpu() string`

GetProductCpu returns the ProductCpu field if non-nil, zero value otherwise.

### GetProductCpuOk

`func (o *Stock) GetProductCpuOk() (*string, bool)`

GetProductCpuOk returns a tuple with the ProductCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCpu

`func (o *Stock) SetProductCpu(v string)`

SetProductCpu sets ProductCpu field to given value.

### HasProductCpu

`func (o *Stock) HasProductCpu() bool`

HasProductCpu returns a boolean if a field has been set.

### GetProductCpuCores

`func (o *Stock) GetProductCpuCores() string`

GetProductCpuCores returns the ProductCpuCores field if non-nil, zero value otherwise.

### GetProductCpuCoresOk

`func (o *Stock) GetProductCpuCoresOk() (*string, bool)`

GetProductCpuCoresOk returns a tuple with the ProductCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCpuCores

`func (o *Stock) SetProductCpuCores(v string)`

SetProductCpuCores sets ProductCpuCores field to given value.

### HasProductCpuCores

`func (o *Stock) HasProductCpuCores() bool`

HasProductCpuCores returns a boolean if a field has been set.

### GetProductGpu

`func (o *Stock) GetProductGpu() string`

GetProductGpu returns the ProductGpu field if non-nil, zero value otherwise.

### GetProductGpuOk

`func (o *Stock) GetProductGpuOk() (*string, bool)`

GetProductGpuOk returns a tuple with the ProductGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGpu

`func (o *Stock) SetProductGpu(v string)`

SetProductGpu sets ProductGpu field to given value.

### HasProductGpu

`func (o *Stock) HasProductGpu() bool`

HasProductGpu returns a boolean if a field has been set.

### GetQuarterlyLocationPremium

`func (o *Stock) GetQuarterlyLocationPremium() float32`

GetQuarterlyLocationPremium returns the QuarterlyLocationPremium field if non-nil, zero value otherwise.

### GetQuarterlyLocationPremiumOk

`func (o *Stock) GetQuarterlyLocationPremiumOk() (*float32, bool)`

GetQuarterlyLocationPremiumOk returns a tuple with the QuarterlyLocationPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterlyLocationPremium

`func (o *Stock) SetQuarterlyLocationPremium(v float32)`

SetQuarterlyLocationPremium sets QuarterlyLocationPremium field to given value.

### HasQuarterlyLocationPremium

`func (o *Stock) HasQuarterlyLocationPremium() bool`

HasQuarterlyLocationPremium returns a boolean if a field has been set.

### GetProductName

`func (o *Stock) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Stock) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Stock) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *Stock) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductQuarterlyPrice

`func (o *Stock) GetProductQuarterlyPrice() float32`

GetProductQuarterlyPrice returns the ProductQuarterlyPrice field if non-nil, zero value otherwise.

### GetProductQuarterlyPriceOk

`func (o *Stock) GetProductQuarterlyPriceOk() (*float32, bool)`

GetProductQuarterlyPriceOk returns a tuple with the ProductQuarterlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductQuarterlyPrice

`func (o *Stock) SetProductQuarterlyPrice(v float32)`

SetProductQuarterlyPrice sets ProductQuarterlyPrice field to given value.

### HasProductQuarterlyPrice

`func (o *Stock) HasProductQuarterlyPrice() bool`

HasProductQuarterlyPrice returns a boolean if a field has been set.

### GetProductOnSale

`func (o *Stock) GetProductOnSale() bool`

GetProductOnSale returns the ProductOnSale field if non-nil, zero value otherwise.

### GetProductOnSaleOk

`func (o *Stock) GetProductOnSaleOk() (*bool, bool)`

GetProductOnSaleOk returns a tuple with the ProductOnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOnSale

`func (o *Stock) SetProductOnSale(v bool)`

SetProductOnSale sets ProductOnSale field to given value.

### HasProductOnSale

`func (o *Stock) HasProductOnSale() bool`

HasProductOnSale returns a boolean if a field has been set.

### GetProductMemory

`func (o *Stock) GetProductMemory() string`

GetProductMemory returns the ProductMemory field if non-nil, zero value otherwise.

### GetProductMemoryOk

`func (o *Stock) GetProductMemoryOk() (*string, bool)`

GetProductMemoryOk returns a tuple with the ProductMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMemory

`func (o *Stock) SetProductMemory(v string)`

SetProductMemory sets ProductMemory field to given value.

### HasProductMemory

`func (o *Stock) HasProductMemory() bool`

HasProductMemory returns a boolean if a field has been set.

### GetProcessorInfo

`func (o *Stock) GetProcessorInfo() map[string]interface{}`

GetProcessorInfo returns the ProcessorInfo field if non-nil, zero value otherwise.

### GetProcessorInfoOk

`func (o *Stock) GetProcessorInfoOk() (*map[string]interface{}, bool)`

GetProcessorInfoOk returns a tuple with the ProcessorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorInfo

`func (o *Stock) SetProcessorInfo(v map[string]interface{})`

SetProcessorInfo sets ProcessorInfo field to given value.

### HasProcessorInfo

`func (o *Stock) HasProcessorInfo() bool`

HasProcessorInfo returns a boolean if a field has been set.

### GetProductDisabledBillingPeriods

`func (o *Stock) GetProductDisabledBillingPeriods() []string`

GetProductDisabledBillingPeriods returns the ProductDisabledBillingPeriods field if non-nil, zero value otherwise.

### GetProductDisabledBillingPeriodsOk

`func (o *Stock) GetProductDisabledBillingPeriodsOk() (*[]string, bool)`

GetProductDisabledBillingPeriodsOk returns a tuple with the ProductDisabledBillingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDisabledBillingPeriods

`func (o *Stock) SetProductDisabledBillingPeriods(v []string)`

SetProductDisabledBillingPeriods sets ProductDisabledBillingPeriods field to given value.

### HasProductDisabledBillingPeriods

`func (o *Stock) HasProductDisabledBillingPeriods() bool`

HasProductDisabledBillingPeriods returns a boolean if a field has been set.

### GetProductDisplayPrice

`func (o *Stock) GetProductDisplayPrice() float32`

GetProductDisplayPrice returns the ProductDisplayPrice field if non-nil, zero value otherwise.

### GetProductDisplayPriceOk

`func (o *Stock) GetProductDisplayPriceOk() (*float32, bool)`

GetProductDisplayPriceOk returns a tuple with the ProductDisplayPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDisplayPrice

`func (o *Stock) SetProductDisplayPrice(v float32)`

SetProductDisplayPrice sets ProductDisplayPrice field to given value.

### HasProductDisplayPrice

`func (o *Stock) HasProductDisplayPrice() bool`

HasProductDisplayPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


