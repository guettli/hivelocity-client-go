# \BandwidthApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDeviceIdBandwidthImageResource**](BandwidthApi.md#PostDeviceIdBandwidthImageResource) | **Post** /bandwidth/device/{deviceId}/image | Get PNG by device
[**PostDeviceIdBandwidthResource**](BandwidthApi.md#PostDeviceIdBandwidthResource) | **Post** /bandwidth/device/{deviceId} | Get data by device
[**PostServiceIdBandwidthImageResource**](BandwidthApi.md#PostServiceIdBandwidthImageResource) | **Post** /bandwidth/service/{serviceId}/image | Get PNG by service
[**PostServiceIdBandwidthResource**](BandwidthApi.md#PostServiceIdBandwidthResource) | **Post** /bandwidth/service/{serviceId} | Get data by service



## PostDeviceIdBandwidthImageResource

> []BandwidthImage PostDeviceIdBandwidthImageResource(ctx, deviceId).Period(period).Interface_(interface_).Start(start).End(end).XFields(xFields).Execute()

Get PNG by device

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
    period := "period_example" // string | Return data in the given period. Day, week, month will return the previous day, week, month from now. (default to "day")
    interface_ := "interface__example" // string | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. (default to "eth0")
    start := int32(56) // int32 | Start Time of Custom Time Period. (Unix Epoch Time) (optional) (default to 0)
    end := int32(56) // int32 | End Time of Custom Time Period (Unix Epoch Time) (optional) (default to 1672750303)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandwidthApi.PostDeviceIdBandwidthImageResource(context.Background(), deviceId).Period(period).Interface_(interface_).Start(start).End(end).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandwidthApi.PostDeviceIdBandwidthImageResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeviceIdBandwidthImageResource`: []BandwidthImage
    fmt.Fprintf(os.Stdout, "Response from `BandwidthApi.PostDeviceIdBandwidthImageResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceIdBandwidthImageResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to &quot;day&quot;]
 **interface_** | **string** | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to &quot;eth0&quot;]
 **start** | **int32** | Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **int32** | End Time of Custom Time Period (Unix Epoch Time) | [default to 1672750303]
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]BandwidthImage**](BandwidthImage.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeviceIdBandwidthResource

> []Bandwidth PostDeviceIdBandwidthResource(ctx, deviceId).Period(period).Interface_(interface_).Step(step).Historical(historical).Start(start).End(end).XFields(xFields).Execute()

Get data by device

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
    period := "period_example" // string | Return data in the given period. Day, week, month will return the previous day, week, month from now. (default to "day")
    interface_ := "interface__example" // string | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. (default to "eth0")
    step := int32(56) // int32 | Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. (default to 300)
    historical := true // bool | If you are a reseller, this will include historical interface data for device regardless of the current device owner. (optional) (default to false)
    start := int32(56) // int32 | Start time of custom time period. (Unix Epoch Time) (optional) (default to 0)
    end := int32(56) // int32 | End time of custom time period (Unix Epoch Time) (optional) (default to 1672750303)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandwidthApi.PostDeviceIdBandwidthResource(context.Background(), deviceId).Period(period).Interface_(interface_).Step(step).Historical(historical).Start(start).End(end).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandwidthApi.PostDeviceIdBandwidthResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeviceIdBandwidthResource`: []Bandwidth
    fmt.Fprintf(os.Stdout, "Response from `BandwidthApi.PostDeviceIdBandwidthResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | ID of Device to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceIdBandwidthResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to &quot;day&quot;]
 **interface_** | **string** | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to &quot;eth0&quot;]
 **step** | **int32** | Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. | [default to 300]
 **historical** | **bool** | If you are a reseller, this will include historical interface data for device regardless of the current device owner. | [default to false]
 **start** | **int32** | Start time of custom time period. (Unix Epoch Time) | [default to 0]
 **end** | **int32** | End time of custom time period (Unix Epoch Time) | [default to 1672750303]
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Bandwidth**](Bandwidth.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceIdBandwidthImageResource

> []BandwidthImage PostServiceIdBandwidthImageResource(ctx, serviceId).Period(period).Interface_(interface_).Start(start).End(end).XFields(xFields).Execute()

Get PNG by service

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
    serviceId := int32(56) // int32 | ID of Service to View
    period := "period_example" // string | Return data in the given period. Day, week, month will return the previous day, week, month from now. (default to "day")
    interface_ := "interface__example" // string | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. (default to "eth0")
    start := int32(56) // int32 | Start Time of Custom Time Period. (Unix Epoch Time) (optional) (default to 0)
    end := int32(56) // int32 | End Time of Custom Time Period (Unix Epoch Time) (optional) (default to 1672750303)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandwidthApi.PostServiceIdBandwidthImageResource(context.Background(), serviceId).Period(period).Interface_(interface_).Start(start).End(end).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandwidthApi.PostServiceIdBandwidthImageResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostServiceIdBandwidthImageResource`: []BandwidthImage
    fmt.Fprintf(os.Stdout, "Response from `BandwidthApi.PostServiceIdBandwidthImageResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | ID of Service to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceIdBandwidthImageResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to &quot;day&quot;]
 **interface_** | **string** | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to &quot;eth0&quot;]
 **start** | **int32** | Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **int32** | End Time of Custom Time Period (Unix Epoch Time) | [default to 1672750303]
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]BandwidthImage**](BandwidthImage.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceIdBandwidthResource

> []Bandwidth PostServiceIdBandwidthResource(ctx, serviceId).Period(period).Interface_(interface_).Step(step).Start(start).End(end).XFields(xFields).Execute()

Get data by service

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
    serviceId := int32(56) // int32 | ID of Service to View
    period := "period_example" // string | Return data in the given period. Day, week, month will return the previous day, week, month from now. (default to "day")
    interface_ := "interface__example" // string | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. (default to "eth0")
    step := int32(56) // int32 | Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. (default to 300)
    start := int32(56) // int32 | Start Time of Custom Time Period. (Unix Epoch Time) (optional) (default to 0)
    end := int32(56) // int32 | End Time of Custom Time Period (Unix Epoch Time) (optional) (default to 1672750303)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandwidthApi.PostServiceIdBandwidthResource(context.Background(), serviceId).Period(period).Interface_(interface_).Step(step).Start(start).End(end).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandwidthApi.PostServiceIdBandwidthResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostServiceIdBandwidthResource`: []Bandwidth
    fmt.Fprintf(os.Stdout, "Response from `BandwidthApi.PostServiceIdBandwidthResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | ID of Service to View | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceIdBandwidthResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to &quot;day&quot;]
 **interface_** | **string** | Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to &quot;eth0&quot;]
 **step** | **int32** | Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. | [default to 300]
 **start** | **int32** | Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **int32** | End Time of Custom Time Period (Unix Epoch Time) | [default to 1672750303]
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Bandwidth**](Bandwidth.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

