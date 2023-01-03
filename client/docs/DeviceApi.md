# \DeviceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeviceBondResource**](DeviceApi.md#DeleteDeviceBondResource) | **Delete** /device/{deviceId}/ports/bond | Unbond ports
[**GetAllDeviceTagOrderResource**](DeviceApi.md#GetAllDeviceTagOrderResource) | **Get** /device/tags-order/all | Get all tag orders
[**GetClientDeviceTagOrderResource**](DeviceApi.md#GetClientDeviceTagOrderResource) | **Get** /device/tags-order | Get tag order
[**GetClientDeviceTagResource**](DeviceApi.md#GetClientDeviceTagResource) | **Get** /device/tags | Get all tags
[**GetDeviceIdEventResource**](DeviceApi.md#GetDeviceIdEventResource) | **Get** /device/{deviceId}/events | Get all events
[**GetDeviceIdResource**](DeviceApi.md#GetDeviceIdResource) | **Get** /device/{deviceId} | Get device
[**GetDeviceIdServiceResource**](DeviceApi.md#GetDeviceIdServiceResource) | **Get** /device/{deviceId}/services | Get all device services
[**GetDeviceIpAssignmentsResource**](DeviceApi.md#GetDeviceIpAssignmentsResource) | **Get** /device/{deviceId}/ips | Get IPs
[**GetDeviceIpminatRuleResource**](DeviceApi.md#GetDeviceIpminatRuleResource) | **Get** /device/{deviceId}/ipmi/nat | Clear IPMI whitelist
[**GetDevicePortResource**](DeviceApi.md#GetDevicePortResource) | **Get** /device/{deviceId}/ports | Get ports
[**GetDeviceResource**](DeviceApi.md#GetDeviceResource) | **Get** /device/ | Get all devices
[**GetDeviceTagIdResource**](DeviceApi.md#GetDeviceTagIdResource) | **Get** /device/{deviceId}/tags | Get device tags
[**GetEffectiveIgnitionIdResource**](DeviceApi.md#GetEffectiveIgnitionIdResource) | **Get** /device/{deviceId}/ignition | Get Ignition injections
[**GetInitialCredsIdResource**](DeviceApi.md#GetInitialCredsIdResource) | **Get** /device/{deviceId}/initial-creds | Get initial creds
[**GetInitialPasswordIdResource**](DeviceApi.md#GetInitialPasswordIdResource) | **Get** /device/{deviceId}/initial-password | Get initial password
[**GetIpmiInfoIdResource**](DeviceApi.md#GetIpmiInfoIdResource) | **Get** /device/{deviceId}/ipmi | Get IPMI data
[**GetIpmiInfoLoginDataResource**](DeviceApi.md#GetIpmiInfoLoginDataResource) | **Get** /device/{deviceId}/ipmi/login-data | Get IPMI creds
[**GetIpmiThresholdsIdResource**](DeviceApi.md#GetIpmiThresholdsIdResource) | **Get** /device/{deviceId}/ipmi/thresholds | Get IPMI thresholds
[**GetIpmiValidLoginIdResource**](DeviceApi.md#GetIpmiValidLoginIdResource) | **Get** /device/{deviceId}/ipmi/valid-login | Get IPMI access status
[**GetPowerResource**](DeviceApi.md#GetPowerResource) | **Get** /device/{deviceId}/power | Get power status
[**PostDeviceBondResource**](DeviceApi.md#PostDeviceBondResource) | **Post** /device/{deviceId}/ports/bond | Bond ports
[**PostDeviceIpmiWhitelistResource**](DeviceApi.md#PostDeviceIpmiWhitelistResource) | **Post** /device/{deviceId}/ipmi/whitelist/ | Whitelist IP for IPMI
[**PostDevicePortIdClearResource**](DeviceApi.md#PostDevicePortIdClearResource) | **Post** /device/{deviceId}/port/{portId}/clear | Clear all Port configurations
[**PostDeviceReloadResource**](DeviceApi.md#PostDeviceReloadResource) | **Post** /device/{deviceId}/reload | Reload device
[**PostPowerResource**](DeviceApi.md#PostPowerResource) | **Post** /device/{deviceId}/power | Update power status
[**PostPreviewEffectiveIgnitionResource**](DeviceApi.md#PostPreviewEffectiveIgnitionResource) | **Post** /device/preview-ignition | Preview Ignition injections
[**PutClientDeviceTagOrderResource**](DeviceApi.md#PutClientDeviceTagOrderResource) | **Put** /device/tags-order | Update tag order
[**PutDeviceIdResource**](DeviceApi.md#PutDeviceIdResource) | **Put** /device/{deviceId} | Update device
[**PutDevicePortResource**](DeviceApi.md#PutDevicePortResource) | **Put** /device/{deviceId}/ports | Update port network
[**PutDeviceTagIdResource**](DeviceApi.md#PutDeviceTagIdResource) | **Put** /device/{deviceId}/tags | Update device tags
[**PutIpmiDevicesThresholdsIdResource**](DeviceApi.md#PutIpmiDevicesThresholdsIdResource) | **Put** /device/ipmi/thresholds | Bulk update IPMI thresholds
[**PutIpmiThresholdsIdResource**](DeviceApi.md#PutIpmiThresholdsIdResource) | **Put** /device/{deviceId}/ipmi/thresholds | Update IPMI thresholds



## DeleteDeviceBondResource

> NetworkTaskDump DeleteDeviceBondResource(ctx, deviceId).XFields(xFields).Execute()

Unbond ports

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
    deviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDeviceBondResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDeviceBondResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeviceBondResource`: NetworkTaskDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeleteDeviceBondResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceBondResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTaskDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDeviceTagOrderResource

> DeviceTag GetAllDeviceTagOrderResource(ctx).XFields(xFields).Execute()

Get all tag orders

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
    resp, r, err := apiClient.DeviceApi.GetAllDeviceTagOrderResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetAllDeviceTagOrderResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDeviceTagOrderResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetAllDeviceTagOrderResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeviceTagOrderResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDeviceTagOrderResource

> DeviceTag GetClientDeviceTagOrderResource(ctx).XFields(xFields).Execute()

Get tag order

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
    resp, r, err := apiClient.DeviceApi.GetClientDeviceTagOrderResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetClientDeviceTagOrderResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientDeviceTagOrderResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetClientDeviceTagOrderResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDeviceTagOrderResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDeviceTagResource

> DeviceTag GetClientDeviceTagResource(ctx).XFields(xFields).Execute()

Get all tags

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
    resp, r, err := apiClient.DeviceApi.GetClientDeviceTagResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetClientDeviceTagResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientDeviceTagResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetClientDeviceTagResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDeviceTagResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIdEventResource

> []DeviceEvent GetDeviceIdEventResource(ctx, deviceId).XFields(xFields).Execute()

Get all events

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
    deviceId := "deviceId_example" // string | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceIdEventResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceIdEventResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceIdEventResource`: []DeviceEvent
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceIdEventResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIdEventResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]DeviceEvent**](DeviceEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIdResource

> DeviceDump GetDeviceIdResource(ctx, deviceId).XFields(xFields).Execute()

Get device

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceIdResource`: DeviceDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceDump**](DeviceDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIdServiceResource

> map[string][]Service GetDeviceIdServiceResource(ctx, deviceId).GroupBy(groupBy).TypeCode(typeCode).Execute()

Get all device services

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
    deviceId := int32(56) // int32 | ID of Device to get all services by type
    groupBy := "groupBy_example" // string | Get results grouped by service type code 'type_code', or 'flat' (optional) (default to "flat")
    typeCode := "typeCode_example" // string | Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceIdServiceResource(context.Background(), deviceId).GroupBy(groupBy).TypeCode(typeCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceIdServiceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceIdServiceResource`: map[string][]Service
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceIdServiceResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to get all services by type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIdServiceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **string** | Get results grouped by service type code &#39;type_code&#39;, or &#39;flat&#39; | [default to &quot;flat&quot;]
 **typeCode** | **string** | Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types | [default to &quot;all&quot;]

### Return type

[**map[string][]Service**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIpAssignmentsResource

> []IPAssignment GetDeviceIpAssignmentsResource(ctx, deviceId).XFields(xFields).Execute()

Get IPs

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
    deviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceIpAssignmentsResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceIpAssignmentsResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceIpAssignmentsResource`: []IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceIpAssignmentsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIpAssignmentsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]IPAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIpminatRuleResource

> GetDeviceIpminatRuleResource(ctx, deviceId).Execute()

Clear IPMI whitelist

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
    deviceId := int32(56) // int32 | ID of a client Device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceIpminatRuleResource(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceIpminatRuleResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of a client Device | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIpminatRuleResourceRequest struct via the builder pattern


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


## GetDevicePortResource

> []DevicePort GetDevicePortResource(ctx, deviceId).XFields(xFields).Execute()

Get ports

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
    deviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevicePortResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevicePortResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicePortResource`: []DevicePort
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevicePortResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicePortResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]DevicePort**](DevicePort.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceResource

> []DeviceDump GetDeviceResource(ctx).XFields(xFields).Execute()

Get all devices

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
    resp, r, err := apiClient.DeviceApi.GetDeviceResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceResource`: []DeviceDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]DeviceDump**](DeviceDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTagIdResource

> DeviceTag GetDeviceTagIdResource(ctx, deviceId).XFields(xFields).Execute()

Get device tags

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceTagIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceTagIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTagIdResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceTagIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTagIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEffectiveIgnitionIdResource

> EffectiveIgnitionResponse GetEffectiveIgnitionIdResource(ctx, deviceId).XFields(xFields).Execute()

Get Ignition injections

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
    deviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetEffectiveIgnitionIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetEffectiveIgnitionIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEffectiveIgnitionIdResource`: EffectiveIgnitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetEffectiveIgnitionIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEffectiveIgnitionIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**EffectiveIgnitionResponse**](EffectiveIgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInitialCredsIdResource

> DeviceInitialCreds GetInitialCredsIdResource(ctx, deviceId).XFields(xFields).Execute()

Get initial creds

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
    deviceId := int32(56) // int32 | ID of Device to retrieve initial authentication credentials for
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetInitialCredsIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetInitialCredsIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInitialCredsIdResource`: DeviceInitialCreds
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetInitialCredsIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to retrieve initial authentication credentials for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInitialCredsIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceInitialCreds**](DeviceInitialCreds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInitialPasswordIdResource

> DevicePassword GetInitialPasswordIdResource(ctx, deviceId).XFields(xFields).Execute()

Get initial password

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
    deviceId := int32(56) // int32 | ID of Device to retrieve initial password
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetInitialPasswordIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetInitialPasswordIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInitialPasswordIdResource`: DevicePassword
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetInitialPasswordIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to retrieve initial password | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInitialPasswordIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DevicePassword**](DevicePassword.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiInfoIdResource

> DeviceIPMIInfo GetIpmiInfoIdResource(ctx, deviceId).XFields(xFields).Execute()

Get IPMI data

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
    deviceId := int32(56) // int32 | ID of Device to retrieve IPMI info.
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetIpmiInfoIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetIpmiInfoIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpmiInfoIdResource`: DeviceIPMIInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetIpmiInfoIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to retrieve IPMI info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiInfoIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceIPMIInfo**](DeviceIPMIInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiInfoLoginDataResource

> IPMILoginData GetIpmiInfoLoginDataResource(ctx, deviceId).XFields(xFields).Execute()

Get IPMI creds

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
    deviceId := int32(56) // int32 | ID of Device to retrieve IPMI Login data.
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetIpmiInfoLoginDataResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetIpmiInfoLoginDataResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpmiInfoLoginDataResource`: IPMILoginData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetIpmiInfoLoginDataResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to retrieve IPMI Login data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiInfoLoginDataResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**IPMILoginData**](IPMILoginData.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiThresholdsIdResource

> DeviceIPMIThresholds GetIpmiThresholdsIdResource(ctx, deviceId).XFields(xFields).Execute()

Get IPMI thresholds

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
    deviceId := int32(56) // int32 | ID of Device to View
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetIpmiThresholdsIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetIpmiThresholdsIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpmiThresholdsIdResource`: DeviceIPMIThresholds
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetIpmiThresholdsIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiThresholdsIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceIPMIThresholds**](DeviceIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiValidLoginIdResource

> IPMIValidLogin GetIpmiValidLoginIdResource(ctx, deviceId).XFields(xFields).Execute()

Get IPMI access status

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
    deviceId := int32(56) // int32 | ID of Device to check IPMI credentials
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetIpmiValidLoginIdResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetIpmiValidLoginIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpmiValidLoginIdResource`: IPMIValidLogin
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetIpmiValidLoginIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to check IPMI credentials | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiValidLoginIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**IPMIValidLogin**](IPMIValidLogin.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerResource

> DevicePower GetPowerResource(ctx, deviceId).XFields(xFields).Execute()

Get power status

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetPowerResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetPowerResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPowerResource`: DevicePower
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetPowerResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPowerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**DevicePower**](DevicePower.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeviceBondResource

> NetworkTaskDump PostDeviceBondResource(ctx, deviceId).XFields(xFields).Execute()

Bond ports

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
    deviceId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostDeviceBondResource(context.Background(), deviceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostDeviceBondResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeviceBondResource`: NetworkTaskDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PostDeviceBondResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceBondResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTaskDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeviceIpmiWhitelistResource

> PostDeviceIpmiWhitelistResource(ctx, deviceId).Payload(payload).Execute()

Whitelist IP for IPMI

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
    deviceId := int32(56) // int32 | ID of the Device to put IP in Whitelist
    payload := *openapiclient.NewDeviceIPMIWhitelistIP("CustIp_example") // DeviceIPMIWhitelistIP | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostDeviceIpmiWhitelistResource(context.Background(), deviceId).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostDeviceIpmiWhitelistResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of the Device to put IP in Whitelist | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceIpmiWhitelistResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeviceIPMIWhitelistIP**](DeviceIPMIWhitelistIP.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevicePortIdClearResource

> NetworkTaskDump PostDevicePortIdClearResource(ctx, deviceId, portId).XFields(xFields).Execute()

Clear all Port configurations

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
    deviceId := int32(56) // int32 | 
    portId := int32(56) // int32 | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostDevicePortIdClearResource(context.Background(), deviceId, portId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostDevicePortIdClearResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDevicePortIdClearResource`: NetworkTaskDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PostDevicePortIdClearResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 
**portId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicePortIdClearResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **string** | An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTaskDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeviceReloadResource

> ServiceOption PostDeviceReloadResource(ctx, deviceId).Payload(payload).XFields(xFields).Execute()

Reload device

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
    deviceId := int32(56) // int32 | 
    payload := *openapiclient.NewDeviceReload(int32(123)) // DeviceReload | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostDeviceReloadResource(context.Background(), deviceId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostDeviceReloadResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeviceReloadResource`: ServiceOption
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PostDeviceReloadResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceReloadResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeviceReload**](DeviceReload.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**ServiceOption**](ServiceOption.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPowerResource

> DevicePower PostPowerResource(ctx, deviceId).Action(action).XFields(xFields).Execute()

Update power status

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    action := "action_example" // string | Must be one of boot|reboot|shutdown
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostPowerResource(context.Background(), deviceId).Action(action).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostPowerResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPowerResource`: DevicePower
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PostPowerResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPowerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | Must be one of boot|reboot|shutdown | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DevicePower**](DevicePower.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPreviewEffectiveIgnitionResource

> PreviewEffectiveIgnitionResponse PostPreviewEffectiveIgnitionResource(ctx).Payload(payload).XFields(xFields).Execute()

Preview Ignition injections

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
    payload := *openapiclient.NewPreviewEffectiveIgnition(map[string]interface{}(123)) // PreviewEffectiveIgnition | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PostPreviewEffectiveIgnitionResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PostPreviewEffectiveIgnitionResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPreviewEffectiveIgnitionResource`: PreviewEffectiveIgnitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PostPreviewEffectiveIgnitionResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPreviewEffectiveIgnitionResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PreviewEffectiveIgnition**](PreviewEffectiveIgnition.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**PreviewEffectiveIgnitionResponse**](PreviewEffectiveIgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientDeviceTagOrderResource

> DeviceTag PutClientDeviceTagOrderResource(ctx).Payload(payload).XFields(xFields).Execute()

Update tag order

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
    payload := *openapiclient.NewDeviceTag() // DeviceTag | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutClientDeviceTagOrderResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutClientDeviceTagOrderResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutClientDeviceTagOrderResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutClientDeviceTagOrderResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutClientDeviceTagOrderResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeviceTag**](DeviceTag.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeviceIdResource

> DeviceDump PutDeviceIdResource(ctx, deviceId).Payload(payload).XFields(xFields).Execute()

Update device

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    payload := *openapiclient.NewDeviceUpdate() // DeviceUpdate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutDeviceIdResource(context.Background(), deviceId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutDeviceIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDeviceIdResource`: DeviceDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutDeviceIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeviceIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeviceUpdate**](DeviceUpdate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceDump**](DeviceDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDevicePortResource

> []NetworkTaskDump PutDevicePortResource(ctx, deviceId).Payload(payload).XFields(xFields).Execute()

Update port network

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
    deviceId := int32(56) // int32 | 
    payload := *openapiclient.NewDevicePortsUpdate() // DevicePortsUpdate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutDevicePortResource(context.Background(), deviceId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutDevicePortResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDevicePortResource`: []NetworkTaskDump
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutDevicePortResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDevicePortResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DevicePortsUpdate**](DevicePortsUpdate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]NetworkTaskDump**](NetworkTaskDump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeviceTagIdResource

> DeviceTag PutDeviceTagIdResource(ctx, deviceId).Payload(payload).XFields(xFields).Execute()

Update device tags

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
    deviceId := int32(56) // int32 | ID of Device to View / Update
    payload := *openapiclient.NewDeviceTag() // DeviceTag | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutDeviceTagIdResource(context.Background(), deviceId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutDeviceTagIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDeviceTagIdResource`: DeviceTag
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutDeviceTagIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeviceTagIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeviceTag**](DeviceTag.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIpmiDevicesThresholdsIdResource

> DevicesIPMIThresholds PutIpmiDevicesThresholdsIdResource(ctx).Payload(payload).XFields(xFields).Execute()

Bulk update IPMI thresholds

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
    payload := *openapiclient.NewUpdateDevicesIPMIThresholds(map[string]interface{}(123), []int32{int32(123)}) // UpdateDevicesIPMIThresholds | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutIpmiDevicesThresholdsIdResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutIpmiDevicesThresholdsIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIpmiDevicesThresholdsIdResource`: DevicesIPMIThresholds
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutIpmiDevicesThresholdsIdResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIpmiDevicesThresholdsIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateDevicesIPMIThresholds**](UpdateDevicesIPMIThresholds.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DevicesIPMIThresholds**](DevicesIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIpmiThresholdsIdResource

> DeviceIPMIThresholds PutIpmiThresholdsIdResource(ctx, deviceId).Payload(payload).XFields(xFields).Execute()

Update IPMI thresholds

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
    deviceId := int32(56) // int32 | ID of Device to View
    payload := *openapiclient.NewDeviceIPMIThresholds(map[string]interface{}(123)) // DeviceIPMIThresholds | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutIpmiThresholdsIdResource(context.Background(), deviceId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutIpmiThresholdsIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIpmiThresholdsIdResource`: DeviceIPMIThresholds
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutIpmiThresholdsIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIpmiThresholdsIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeviceIPMIThresholds**](DeviceIPMIThresholds.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**DeviceIPMIThresholds**](DeviceIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

