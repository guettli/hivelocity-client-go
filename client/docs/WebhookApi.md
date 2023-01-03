# \WebhookApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhookIdResource**](WebhookApi.md#DeleteWebhookIdResource) | **Delete** /webhooks/{webhookId} | Deletes a single webhook
[**GetWebhookEventResource**](WebhookApi.md#GetWebhookEventResource) | **Get** /webhooks/events | Returns all available Webhook Events
[**GetWebhookIdResource**](WebhookApi.md#GetWebhookIdResource) | **Get** /webhooks/{webhookId} | Returns detailed information for a Single Webhook
[**GetWebhookResource**](WebhookApi.md#GetWebhookResource) | **Get** /webhooks/ | Returns your active Webhooks
[**PostEventScriptActionTriggerResource**](WebhookApi.md#PostEventScriptActionTriggerResource) | **Post** /webhooks/trigger | Queues a webhook for the event script action that was triggered
[**PostWebhookResource**](WebhookApi.md#PostWebhookResource) | **Post** /webhooks/ | Create a new Webhook for a Webhook Event
[**PutWebhookIdResource**](WebhookApi.md#PutWebhookIdResource) | **Put** /webhooks/{webhookId} | Updates a Single Webhook



## DeleteWebhookIdResource

> DeleteWebhookIdResource(ctx, webhookId).Execute()

Deletes a single webhook

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
    webhookId := int32(56) // int32 | ID of Webhook to View / Update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.DeleteWebhookIdResource(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.DeleteWebhookIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID of Webhook to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookIdResourceRequest struct via the builder pattern


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


## GetWebhookEventResource

> []WebhookEvent GetWebhookEventResource(ctx).XFields(xFields).Execute()

Returns all available Webhook Events

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
    resp, r, err := apiClient.WebhookApi.GetWebhookEventResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhookEventResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookEventResource`: []WebhookEvent
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhookEventResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEventResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]WebhookEvent**](WebhookEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookIdResource

> Webhook GetWebhookIdResource(ctx, webhookId).XFields(xFields).Execute()

Returns detailed information for a Single Webhook

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
    webhookId := int32(56) // int32 | ID of Webhook to View / Update
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhookIdResource(context.Background(), webhookId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhookIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookIdResource`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhookIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID of Webhook to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookResource

> []Webhook GetWebhookResource(ctx).XFields(xFields).Execute()

Returns your active Webhooks

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
    resp, r, err := apiClient.WebhookApi.GetWebhookResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhookResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookResource`: []Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhookResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventScriptActionTriggerResource

> PostEventScriptActionTriggerResource(ctx).WebhookName(webhookName).Execute()

Queues a webhook for the event script action that was triggered

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
    webhookName := "webhookName_example" // string | The name of the webhook to trigger. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.PostEventScriptActionTriggerResource(context.Background()).WebhookName(webhookName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PostEventScriptActionTriggerResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventScriptActionTriggerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookName** | **string** | The name of the webhook to trigger. | 

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


## PostWebhookResource

> Webhook PostWebhookResource(ctx).Payload(payload).XFields(xFields).Execute()

Create a new Webhook for a Webhook Event

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
    payload := *openapiclient.NewWebhookCreate("Event_example", "Url_example") // WebhookCreate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.PostWebhookResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PostWebhookResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebhookResource`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.PostWebhookResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**WebhookCreate**](WebhookCreate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhookIdResource

> Webhook PutWebhookIdResource(ctx, webhookId).Payload(payload).XFields(xFields).Execute()

Updates a Single Webhook

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
    webhookId := int32(56) // int32 | ID of Webhook to View / Update
    payload := *openapiclient.NewWebhookUpdate() // WebhookUpdate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.PutWebhookIdResource(context.Background(), webhookId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PutWebhookIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhookIdResource`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.PutWebhookIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID of Webhook to View / Update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhookIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**WebhookUpdate**](WebhookUpdate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

