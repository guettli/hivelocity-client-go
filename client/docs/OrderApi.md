# \OrderApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderIdResource**](OrderApi.md#GetOrderIdResource) | **Get** /order/{orderId} | Return details of a specific Order
[**GetOrderInprogressResource**](OrderApi.md#GetOrderInprogressResource) | **Get** /order/in-progress | Return a list with all Orders
[**GetOrderResource**](OrderApi.md#GetOrderResource) | **Get** /order/ | Return a list with all Orders



## GetOrderIdResource

> OrderDump GetOrderIdResource(ctx, orderId).XFields(xFields).Execute()

Return details of a specific Order

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
    orderId := int32(56) // int32 | \"ID of Order to View\"
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetOrderIdResource(context.Background(), orderId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrderIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderIdResource`: OrderDump
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrderIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | \&quot;ID of Order to View\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**OrderDump**](OrderDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderInprogressResource

> []OrderDump GetOrderInprogressResource(ctx).XFields(xFields).Execute()

Return a list with all Orders

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
    resp, r, err := apiClient.OrderApi.GetOrderInprogressResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrderInprogressResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderInprogressResource`: []OrderDump
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrderInprogressResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderInprogressResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]OrderDump**](OrderDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderResource

> []OrderDump GetOrderResource(ctx).XFields(xFields).Execute()

Return a list with all Orders

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
    resp, r, err := apiClient.OrderApi.GetOrderResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrderResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderResource`: []OrderDump
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrderResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]OrderDump**](OrderDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

