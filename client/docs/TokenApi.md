# \TokenApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokenIdResource**](TokenApi.md#DeleteTokenIdResource) | **Delete** /token/{token} | Deletes the Public API Token
[**GetTokenIdResource**](TokenApi.md#GetTokenIdResource) | **Get** /token/{token} | Returns Public API Token
[**GetTokenResource**](TokenApi.md#GetTokenResource) | **Get** /token/ | Returns a list of Public API Tokens for the current user
[**PostTokenResource**](TokenApi.md#PostTokenResource) | **Post** /token/ | Create a new Public API Token for the current user
[**PutTokenIdResource**](TokenApi.md#PutTokenIdResource) | **Put** /token/{token} | Updates the Public API Token



## DeleteTokenIdResource

> DeleteTokenIdResource(ctx, token).Execute()

Deletes the Public API Token

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
    token := "token_example" // string | Public API Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.DeleteTokenIdResource(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.DeleteTokenIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Public API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenIdResource

> PublicApiTokenDump GetTokenIdResource(ctx, token).XFields(xFields).Execute()

Returns Public API Token

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
    token := "token_example" // string | Public API Token
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.GetTokenIdResource(context.Background(), token).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetTokenIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenIdResource`: PublicApiTokenDump
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetTokenIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Public API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**PublicApiTokenDump**](PublicApiTokenDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenResource

> []PublicApiTokenDump GetTokenResource(ctx).XFields(xFields).Execute()

Returns a list of Public API Tokens for the current user

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
    resp, r, err := apiClient.TokenApi.GetTokenResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetTokenResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenResource`: []PublicApiTokenDump
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetTokenResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]PublicApiTokenDump**](PublicApiTokenDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTokenResource

> PublicApiTokenDump PostTokenResource(ctx).Payload(payload).XFields(xFields).Execute()

Create a new Public API Token for the current user

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
    payload := *openapiclient.NewPublicApiTokenLoad() // PublicApiTokenLoad | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.PostTokenResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.PostTokenResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTokenResource`: PublicApiTokenDump
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.PostTokenResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTokenResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PublicApiTokenLoad**](PublicApiTokenLoad.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**PublicApiTokenDump**](PublicApiTokenDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTokenIdResource

> PublicApiTokenDump PutTokenIdResource(ctx, token).Payload(payload).XFields(xFields).Execute()

Updates the Public API Token

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
    token := "token_example" // string | Public API Token
    payload := *openapiclient.NewPublicApiTokenLoad() // PublicApiTokenLoad | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.PutTokenIdResource(context.Background(), token).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.PutTokenIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTokenIdResource`: PublicApiTokenDump
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.PutTokenIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Public API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTokenIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**PublicApiTokenLoad**](PublicApiTokenLoad.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**PublicApiTokenDump**](PublicApiTokenDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

