# \CreditApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCreditResource**](CreditApi.md#GetCreditResource) | **Get** /credit/ | Return a list with all Credits
[**GetTotalActiveCreditResource**](CreditApi.md#GetTotalActiveCreditResource) | **Get** /credit/total | Return the client&#39;s total active credit amount
[**PostCreditResource**](CreditApi.md#PostCreditResource) | **Post** /credit/ | Receive billing method id and amount and return the created Credit



## GetCreditResource

> []Credit GetCreditResource(ctx).XFields(xFields).Execute()

Return a list with all Credits

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
    resp, r, err := apiClient.CreditApi.GetCreditResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditApi.GetCreditResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreditResource`: []Credit
    fmt.Fprintf(os.Stdout, "Response from `CreditApi.GetCreditResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Credit**](Credit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalActiveCreditResource

> TotalActiveCredit GetTotalActiveCreditResource(ctx).XFields(xFields).Execute()

Return the client's total active credit amount

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
    resp, r, err := apiClient.CreditApi.GetTotalActiveCreditResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditApi.GetTotalActiveCreditResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTotalActiveCreditResource`: TotalActiveCredit
    fmt.Fprintf(os.Stdout, "Response from `CreditApi.GetTotalActiveCreditResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalActiveCreditResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**TotalActiveCredit**](TotalActiveCredit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreditResource

> Credit PostCreditResource(ctx).Payload(payload).XFields(xFields).Execute()

Receive billing method id and amount and return the created Credit

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
    payload := *openapiclient.NewCreateCredit() // CreateCredit | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditApi.PostCreditResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditApi.PostCreditResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreditResource`: Credit
    fmt.Fprintf(os.Stdout, "Response from `CreditApi.PostCreditResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreditResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateCredit**](CreateCredit.md) |  | 
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

