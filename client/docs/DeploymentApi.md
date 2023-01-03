# \DeploymentApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeploymentIdResource**](DeploymentApi.md#DeleteDeploymentIdResource) | **Delete** /deploy/{deploymentId} | Delete the specified deployment
[**GetDeploymentIdResource**](DeploymentApi.md#GetDeploymentIdResource) | **Get** /deploy/{deploymentId} | Return a dictionary with deployment information
[**GetDeploymentResource**](DeploymentApi.md#GetDeploymentResource) | **Get** /deploy/ | Return a list with all client deployments
[**PostDeploymentIdResource**](DeploymentApi.md#PostDeploymentIdResource) | **Post** /deploy/{deploymentId} | Input a billing info id and script to process and finish a deployment
[**PostDeploymentResource**](DeploymentApi.md#PostDeploymentResource) | **Post** /deploy/ | Start a new deployment
[**PutDeploymentIdResource**](DeploymentApi.md#PutDeploymentIdResource) | **Put** /deploy/{deploymentId} | Receive product, quantity and options to be added on the deployment



## DeleteDeploymentIdResource

> DeleteDeploymentIdResource(ctx, deploymentId).Execute()

Delete the specified deployment

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
    deploymentId := int32(56) // int32 | Id of the deployment to interact with

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.DeleteDeploymentIdResource(context.Background(), deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DeleteDeploymentIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int32** | Id of the deployment to interact with | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentIdResourceRequest struct via the builder pattern


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


## GetDeploymentIdResource

> Deployment GetDeploymentIdResource(ctx, deploymentId).XFields(xFields).Execute()

Return a dictionary with deployment information

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
    deploymentId := int32(56) // int32 | Id of the deployment to interact with
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.GetDeploymentIdResource(context.Background(), deploymentId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetDeploymentIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentIdResource`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetDeploymentIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int32** | Id of the deployment to interact with | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentResource

> []Deployment GetDeploymentResource(ctx).XFields(xFields).Execute()

Return a list with all client deployments

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
    resp, r, err := apiClient.DeploymentApi.GetDeploymentResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetDeploymentResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentResource`: []Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetDeploymentResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeploymentIdResource

> PostDeploymentIdResource(ctx, deploymentId).Payload(payload).Execute()

Input a billing info id and script to process and finish a deployment

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
    deploymentId := int32(56) // int32 | Id of the deployment to interact with
    payload := *openapiclient.NewDeploymentStart(int32(123)) // DeploymentStart | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.PostDeploymentIdResource(context.Background(), deploymentId).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.PostDeploymentIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int32** | Id of the deployment to interact with | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeploymentIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeploymentStart**](DeploymentStart.md) |  | 

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


## PostDeploymentResource

> Deployment PostDeploymentResource(ctx).DeploymentName(deploymentName).XFields(xFields).Execute()

Start a new deployment

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
    deploymentName := "deploymentName_example" // string |  (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.PostDeploymentResource(context.Background()).DeploymentName(deploymentName).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.PostDeploymentResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeploymentResource`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.PostDeploymentResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDeploymentResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentName** | **string** |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeploymentIdResource

> Deployment PutDeploymentIdResource(ctx, deploymentId).Payload(payload).XFields(xFields).Execute()

Receive product, quantity and options to be added on the deployment

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
    deploymentId := int32(56) // int32 | Id of the deployment to interact with
    payload := *openapiclient.NewDeploymentCustomization([]string{"Hostnames_example"}, "OperatingSystem_example", int32(123)) // DeploymentCustomization | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.PutDeploymentIdResource(context.Background(), deploymentId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.PutDeploymentIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDeploymentIdResource`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.PutDeploymentIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int32** | Id of the deployment to interact with | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DeploymentCustomization**](DeploymentCustomization.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

