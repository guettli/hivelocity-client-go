# \ServiceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceIdResource**](ServiceApi.md#GetServiceIdResource) | **Get** /service/{serviceId} | Return a dictionary with data from a specific account service
[**GetServiceManageableResource**](ServiceApi.md#GetServiceManageableResource) | **Get** /service/{serviceId}/manageable | Return a dictionary with data verifying managed services eligibility
[**GetServiceManagedReqResource**](ServiceApi.md#GetServiceManagedReqResource) | **Get** /service/managed-requirements | Return a dictionary with managed services operating system and panel requirements
[**GetServiceResource**](ServiceApi.md#GetServiceResource) | **Get** /service/ | Return a list of all account services
[**GetServiceTypeResource**](ServiceApi.md#GetServiceTypeResource) | **Get** /service/types | Return all available service types



## GetServiceIdResource

> Service GetServiceIdResource(ctx, serviceId).IncludeZeroPriceOptions(includeZeroPriceOptions).XFields(xFields).Execute()

Return a dictionary with data from a specific account service

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
    serviceId := int32(56) // int32 | 
    includeZeroPriceOptions := true // bool | Includes on the field serviceOptions the options having price equal to zero (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.GetServiceIdResource(context.Background(), serviceId).IncludeZeroPriceOptions(includeZeroPriceOptions).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceIdResource`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeZeroPriceOptions** | **bool** | Includes on the field serviceOptions the options having price equal to zero | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Service**](Service.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceManageableResource

> ServiceManageable GetServiceManageableResource(ctx, serviceId).XFields(xFields).Execute()

Return a dictionary with data verifying managed services eligibility

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
    serviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.GetServiceManageableResource(context.Background(), serviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceManageableResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceManageableResource`: ServiceManageable
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceManageableResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceManageableResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**ServiceManageable**](ServiceManageable.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceManagedReqResource

> ServiceManagedReq GetServiceManagedReqResource(ctx).XFields(xFields).Execute()

Return a dictionary with managed services operating system and panel requirements

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
    resp, r, err := apiClient.ServiceApi.GetServiceManagedReqResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceManagedReqResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceManagedReqResource`: ServiceManagedReq
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceManagedReqResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceManagedReqResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**ServiceManagedReq**](ServiceManagedReq.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceResource

> []FastServiceDump GetServiceResource(ctx).Status(status).TypeCode(typeCode).OrderId(orderId).XFields(xFields).Execute()

Return a list of all account services

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
    status := "status_example" // string | The current status of the service (default to "all")
    typeCode := "typeCode_example" // string | Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types (optional) (default to "null")
    orderId := int32(56) // int32 | Order id of the service (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.GetServiceResource(context.Background()).Status(status).TypeCode(typeCode).OrderId(orderId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceResource`: []FastServiceDump
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The current status of the service | [default to &quot;all&quot;]
 **typeCode** | **string** | Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types | [default to &quot;null&quot;]
 **orderId** | **int32** | Order id of the service | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]FastServiceDump**](FastServiceDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTypeResource

> []ServiceType GetServiceTypeResource(ctx).XFields(xFields).Execute()

Return all available service types

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
    resp, r, err := apiClient.ServiceApi.GetServiceTypeResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceTypeResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceTypeResource`: []ServiceType
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceTypeResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTypeResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]ServiceType**](ServiceType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

