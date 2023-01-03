# \ProductApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductOperatingSystemsResource**](ProductApi.md#GetProductOperatingSystemsResource) | **Get** /product/{productId}/operating-systems | Get product operating systems
[**GetProductOptionResource**](ProductApi.md#GetProductOptionResource) | **Get** /product/{productId}/options | Get product options
[**GetProductsAndOptionsResource**](ProductApi.md#GetProductsAndOptionsResource) | **Get** /product/options | Get all options



## GetProductOperatingSystemsResource

> []Option GetProductOperatingSystemsResource(ctx, productId).XFields(xFields).Execute()

Get product operating systems

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
    productId := int32(56) // int32 | ID of the Product
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetProductOperatingSystemsResource(context.Background(), productId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductOperatingSystemsResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductOperatingSystemsResource`: []Option
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductOperatingSystemsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | ID of the Product | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductOperatingSystemsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Option**](Option.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductOptionResource

> map[string][]Option GetProductOptionResource(ctx, productId).GroupBy(groupBy).Execute()

Get product options

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
    productId := int32(56) // int32 | ID of the Product
    groupBy := "groupBy_example" // string | Get results grouped by 'upgrade' or without grouping. (optional) (default to "upgrade")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetProductOptionResource(context.Background(), productId).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductOptionResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductOptionResource`: map[string][]Option
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductOptionResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | ID of the Product | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductOptionResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **string** | Get results grouped by &#39;upgrade&#39; or without grouping. | [default to &quot;upgrade&quot;]

### Return type

[**map[string][]Option**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductsAndOptionsResource

> []ProductOption GetProductsAndOptionsResource(ctx).XFields(xFields).Execute()

Get all options

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
    resp, r, err := apiClient.ProductApi.GetProductsAndOptionsResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductsAndOptionsResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductsAndOptionsResource`: []ProductOption
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductsAndOptionsResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsAndOptionsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]ProductOption**](ProductOption.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

