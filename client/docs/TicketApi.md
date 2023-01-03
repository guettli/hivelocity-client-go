# \TicketApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTicketIdResource**](TicketApi.md#GetTicketIdResource) | **Get** /tickets/{ticketId} | Returns details of a specific ticket
[**GetTicketReplyResource**](TicketApi.md#GetTicketReplyResource) | **Get** /tickets/{ticketId}/reply | Returns a reply for a specific ticket
[**GetTicketResource**](TicketApi.md#GetTicketResource) | **Get** /tickets/ | Returns a list with all Tickets
[**GetTicketSearchResource**](TicketApi.md#GetTicketSearchResource) | **Get** /tickets/search | Return results of ticket search
[**PostTicketReplyResource**](TicketApi.md#PostTicketReplyResource) | **Post** /tickets/{ticketId}/reply | Creates reply for a specific Ticket
[**PostTicketResource**](TicketApi.md#PostTicketResource) | **Post** /tickets/ | Creates a new ticket
[**PutTicketIdResource**](TicketApi.md#PutTicketIdResource) | **Put** /tickets/{ticketId} | Updates a specific ticket



## GetTicketIdResource

> Ticket GetTicketIdResource(ctx, ticketId).XFields(xFields).Execute()

Returns details of a specific ticket

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
    ticketId := int32(56) // int32 | ticket database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.GetTicketIdResource(context.Background(), ticketId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.GetTicketIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicketIdResource`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.GetTicketIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | ticket database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketReplyResource

> []TicketPost GetTicketReplyResource(ctx, ticketId).XFields(xFields).Execute()

Returns a reply for a specific ticket

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
    ticketId := int32(56) // int32 | ticket database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.GetTicketReplyResource(context.Background(), ticketId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.GetTicketReplyResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicketReplyResource`: []TicketPost
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.GetTicketReplyResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | ticket database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketReplyResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]TicketPost**](TicketPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketResource

> []Ticket GetTicketResource(ctx).XFields(xFields).Execute()

Returns a list with all Tickets

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
    resp, r, err := apiClient.TicketApi.GetTicketResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.GetTicketResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicketResource`: []Ticket
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.GetTicketResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketSearchResource

> TicketSeachResult GetTicketSearchResource(ctx).PerPage(perPage).Page(page).Q(q).XFields(xFields).Execute()

Return results of ticket search

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
    perPage := int32(56) // int32 | Number of items per page. (optional) (default to 10)
    page := int32(56) // int32 | The page number of search. (optional) (default to 1)
    q := "q_example" // string | Content search. (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.GetTicketSearchResource(context.Background()).PerPage(perPage).Page(page).Q(q).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.GetTicketSearchResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicketSearchResource`: TicketSeachResult
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.GetTicketSearchResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketSearchResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of items per page. | [default to 10]
 **page** | **int32** | The page number of search. | [default to 1]
 **q** | **string** | Content search. | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**TicketSeachResult**](TicketSeachResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketReplyResource

> TicketPost PostTicketReplyResource(ctx, ticketId).Payload(payload).XFields(xFields).Execute()

Creates reply for a specific Ticket

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
    ticketId := int32(56) // int32 | ticket database ID
    payload := *openapiclient.NewTicketCreateReply("Body_example") // TicketCreateReply | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.PostTicketReplyResource(context.Background(), ticketId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.PostTicketReplyResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTicketReplyResource`: TicketPost
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.PostTicketReplyResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | ticket database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTicketReplyResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**TicketCreateReply**](TicketCreateReply.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**TicketPost**](TicketPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketResource

> Ticket PostTicketResource(ctx).Payload(payload).XFields(xFields).Execute()

Creates a new ticket

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
    payload := *openapiclient.NewTicketCreate("Subject_example", "Queue_example") // TicketCreate | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.PostTicketResource(context.Background()).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.PostTicketResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTicketResource`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.PostTicketResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTicketResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**TicketCreate**](TicketCreate.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTicketIdResource

> Ticket PutTicketIdResource(ctx, ticketId).Payload(payload).XFields(xFields).Execute()

Updates a specific ticket

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
    ticketId := int32(56) // int32 | ticket database ID
    payload := *openapiclient.NewTicketPut() // TicketPut | 
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.PutTicketIdResource(context.Background(), ticketId).Payload(payload).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.PutTicketIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTicketIdResource`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.PutTicketIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | ticket database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTicketIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**TicketPut**](TicketPut.md) |  | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

