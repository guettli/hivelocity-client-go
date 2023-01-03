# \InventoryApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocationResource**](InventoryApi.md#GetLocationResource) | **Get** /inventory/locations | Get all facilities
[**GetStockByProductResource**](InventoryApi.md#GetStockByProductResource) | **Get** /inventory/product/{productId} | Get product
[**GetStockResource**](InventoryApi.md#GetStockResource) | **Get** /inventory/product | Get all products



## GetLocationResource

> []Location GetLocationResource(ctx).XFields(xFields).Execute()

Get all facilities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetLocationResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetLocationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationResource`: []Location
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetLocationResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockByProductResource

> Stock GetStockByProductResource(ctx, productId).XFields(xFields).Execute()

Get product

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productId := int32(56) // int32 | Product database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetStockByProductResource(context.Background(), productId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetStockByProductResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockByProductResource`: Stock
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetStockByProductResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | Product database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStockByProductResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**Stock**](Stock.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockResource

> map[string][]Stock GetStockResource(ctx).Location(location).IncludeCustom(includeCustom).BondingSupport(bondingSupport).GroupBy(groupBy).Execute()

Get all products

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    location := "location_example" // string | Retrieve products by their location in the Hivelocity store:   MAIN: The main list of instant (and custom servers). MASSIVESTORAGE: Servers with up to 90 drives. GPU: Servers with GPUs. OUTLET: Discounted older hardware. LANDING: Periodic specials for events. (optional) (default to "MAIN")
    includeCustom := true // bool | Include custom server products. (optional) (default to false)
    bondingSupport := "bondingSupport_example" // string | Filter by bonding support. Values are true/false. None will return a mix of both. *DEPRECATED:* Soon, all servers sold will have bonding support and this will be removed. (optional) (default to "null")
    groupBy := "groupBy_example" // string | Get results grouped by 'city', 'facility', or 'flat' (optional) (default to "facility")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetStockResource(context.Background()).Location(location).IncludeCustom(includeCustom).BondingSupport(bondingSupport).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetStockResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockResource`: map[string][]Stock
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetStockResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStockResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | Retrieve products by their location in the Hivelocity store:   MAIN: The main list of instant (and custom servers). MASSIVESTORAGE: Servers with up to 90 drives. GPU: Servers with GPUs. OUTLET: Discounted older hardware. LANDING: Periodic specials for events. | [default to &quot;MAIN&quot;]
 **includeCustom** | **bool** | Include custom server products. | [default to false]
 **bondingSupport** | **string** | Filter by bonding support. Values are true/false. None will return a mix of both. *DEPRECATED:* Soon, all servers sold will have bonding support and this will be removed. | [default to &quot;null&quot;]
 **groupBy** | **string** | Get results grouped by &#39;city&#39;, &#39;facility&#39;, or &#39;flat&#39; | [default to &quot;facility&quot;]

### Return type

[**map[string][]Stock**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

