# \BillingInfoApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingInfoResource**](BillingInfoApi.md#GetBillingInfoResource) | **Get** /billing-info/ | Return a list with all Billing Info
[**PostBillingInfoResource**](BillingInfoApi.md#PostBillingInfoResource) | **Post** /billing-info/ | Create verification for credit card with all Billing Info
[**PutBillingInfoResource**](BillingInfoApi.md#PutBillingInfoResource) | **Put** /billing-info/ | Verify credit card with all Billing Info



## GetBillingInfoResource

> []BillingInfo GetBillingInfoResource(ctx).XFields(xFields).Execute()

Return a list with all Billing Info

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
    resp, r, err := apiClient.BillingInfoApi.GetBillingInfoResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoApi.GetBillingInfoResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingInfoResource`: []BillingInfo
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoApi.GetBillingInfoResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingInfoResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]BillingInfo**](BillingInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingInfoResource

> BillingInfo PostBillingInfoResource(ctx).Payload(payload).XFields(xFields).Execute()

Create verification for credit card with all Billing Info

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
    payload := *openapiclient.NewBillingInfoVerification(int32(123)) // BillingInfoVerification | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoApi.PostBillingInfoResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoApi.PostBillingInfoResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBillingInfoResource`: BillingInfo
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoApi.PostBillingInfoResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingInfoResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**BillingInfoVerification**](BillingInfoVerification.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**BillingInfo**](BillingInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBillingInfoResource

> Credit PutBillingInfoResource(ctx).Payload(payload).XFields(xFields).Execute()

Verify credit card with all Billing Info

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
    payload := *openapiclient.NewBillingInfoVerification(int32(123)) // BillingInfoVerification | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingInfoApi.PutBillingInfoResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInfoApi.PutBillingInfoResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBillingInfoResource`: Credit
    fmt.Fprintf(os.Stdout, "Response from `BillingInfoApi.PutBillingInfoResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBillingInfoResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**BillingInfoVerification**](BillingInfoVerification.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Credit**](Credit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

