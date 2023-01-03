# \PermissionApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionAllResource**](PermissionApi.md#GetPermissionAllResource) | **Get** /permission/ | Endpoint to get All Permissions
[**GetPermissionContactResource**](PermissionApi.md#GetPermissionContactResource) | **Get** /permission/contact/{contactId} | Endpoint to get Contact Permissions
[**GetPermissionUserResource**](PermissionApi.md#GetPermissionUserResource) | **Get** /permission/user | Endpoint to get User Permissions
[**PostPermissionAssignContactResource**](PermissionApi.md#PostPermissionAssignContactResource) | **Post** /permission/contact | Endpoint to assign a new Permission to a Contact



## GetPermissionAllResource

> GetPermissionAllResource(ctx).Execute()

Endpoint to get All Permissions

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.GetPermissionAllResource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.GetPermissionAllResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionAllResourceRequest struct via the builder pattern


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


## GetPermissionContactResource

> GetPermissionContactResource(ctx, contactId).Execute()

Endpoint to get Contact Permissions

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
    contactId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.GetPermissionContactResource(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.GetPermissionContactResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionContactResourceRequest struct via the builder pattern


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


## GetPermissionUserResource

> GetPermissionUserResource(ctx).Execute()

Endpoint to get User Permissions

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.GetPermissionUserResource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.GetPermissionUserResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionUserResourceRequest struct via the builder pattern


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


## PostPermissionAssignContactResource

> []PermissionReturn PostPermissionAssignContactResource(ctx).Payload(payload).XFields(xFields).Execute()

Endpoint to assign a new Permission to a Contact

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
    payload := *openapiclient.NewPermission(int32(123), []string{"Permissions_example"}) // Permission | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.PostPermissionAssignContactResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.PostPermissionAssignContactResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPermissionAssignContactResource`: []PermissionReturn
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.PostPermissionAssignContactResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPermissionAssignContactResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**Permission**](Permission.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]PermissionReturn**](PermissionReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

