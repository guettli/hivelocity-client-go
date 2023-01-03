# \SshKeyApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSshKeyIdResource**](SshKeyApi.md#DeleteSshKeyIdResource) | **Delete** /ssh_key/{sshKeyId} | Removes public ssh key
[**GetSshKeyIdResource**](SshKeyApi.md#GetSshKeyIdResource) | **Get** /ssh_key/{sshKeyId} | Get public ssh key
[**GetSshKeyResource**](SshKeyApi.md#GetSshKeyResource) | **Get** /ssh_key/ | Gets all public ssh key
[**PostSshKeyResource**](SshKeyApi.md#PostSshKeyResource) | **Post** /ssh_key/ | Adds public ssh key
[**PutSshKeyIdResource**](SshKeyApi.md#PutSshKeyIdResource) | **Put** /ssh_key/{sshKeyId} | Updates public ssh key



## DeleteSshKeyIdResource

> DeleteSshKeyIdResource(ctx, sshKeyId).Execute()

Removes public ssh key

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
    sshKeyId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyApi.DeleteSshKeyIdResource(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyApi.DeleteSshKeyIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeyIdResourceRequest struct via the builder pattern


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


## GetSshKeyIdResource

> SshKeyResponse GetSshKeyIdResource(ctx, sshKeyId).XFields(xFields).Execute()

Get public ssh key

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
    sshKeyId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyApi.GetSshKeyIdResource(context.Background(), sshKeyId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyApi.GetSshKeyIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSshKeyIdResource`: SshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyApi.GetSshKeyIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeyIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**SshKeyResponse**](SshKeyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKeyResource

> []SshKeyResponse GetSshKeyResource(ctx).XFields(xFields).Execute()

Gets all public ssh key

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
    resp, r, err := apiClient.SshKeyApi.GetSshKeyResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyApi.GetSshKeyResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSshKeyResource`: []SshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyApi.GetSshKeyResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeyResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]SshKeyResponse**](SshKeyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshKeyResource

> SshKeyResponse PostSshKeyResource(ctx).Payload(payload).XFields(xFields).Execute()

Adds public ssh key

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
    payload := *openapiclient.NewSshKey("Name_example", "PublicKey_example") // SshKey | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyApi.PostSshKeyResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyApi.PostSshKeyResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSshKeyResource`: SshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyApi.PostSshKeyResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshKeyResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**SshKey**](SshKey.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**SshKeyResponse**](SshKeyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSshKeyIdResource

> SshKeyResponse PutSshKeyIdResource(ctx, sshKeyId).Payload(payload).XFields(xFields).Execute()

Updates public ssh key

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
    sshKeyId := int32(56) // int32 | 
    payload := *openapiclient.NewSshKeyUpdate() // SshKeyUpdate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyApi.PutSshKeyIdResource(context.Background(), sshKeyId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyApi.PutSshKeyIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSshKeyIdResource`: SshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyApi.PutSshKeyIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSshKeyIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**SshKeyUpdate**](SshKeyUpdate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**SshKeyResponse**](SshKeyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

