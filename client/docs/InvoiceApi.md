# \InvoiceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceIdDetails**](InvoiceApi.md#GetInvoiceIdDetails) | **Get** /invoice/{invoiceId}/details | Return detailed information for an invoice
[**GetInvoiceIdResource**](InvoiceApi.md#GetInvoiceIdResource) | **Get** /invoice/{invoiceId} | Return serialized data on a single invoice
[**GetInvoicePdfResource**](InvoiceApi.md#GetInvoicePdfResource) | **Get** /invoice/{invoiceId}/pdf-download | Return an Invoice PDF file in Base64 Encoded Format
[**GetInvoiceResource**](InvoiceApi.md#GetInvoiceResource) | **Get** /invoice/ | Return serialized data on all invoices
[**GetInvoiceSearchResource**](InvoiceApi.md#GetInvoiceSearchResource) | **Get** /invoice/search | Return results of invoice search
[**GetInvoiceUnpaidResource**](InvoiceApi.md#GetInvoiceUnpaidResource) | **Get** /invoice/unpaid | Return total balance of all unpaid invoices



## GetInvoiceIdDetails

> InvoiceDetails GetInvoiceIdDetails(ctx, invoiceId).XFields(xFields).Execute()

Return detailed information for an invoice



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
    invoiceId := int32(56) // int32 | Invoice database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoiceIdDetails(context.Background(), invoiceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceIdDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceIdDetails`: InvoiceDetails
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | Invoice database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceIdDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**InvoiceDetails**](InvoiceDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceIdResource

> Invoice GetInvoiceIdResource(ctx, invoiceId).XFields(xFields).Execute()

Return serialized data on a single invoice

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
    invoiceId := int32(56) // int32 | Invoice database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoiceIdResource(context.Background(), invoiceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceIdResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceIdResource`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceIdResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | Invoice database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceIdResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicePdfResource

> InvoicePDF GetInvoicePdfResource(ctx, invoiceId).XFields(xFields).Execute()

Return an Invoice PDF file in Base64 Encoded Format

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
    invoiceId := int32(56) // int32 | Invoice database ID
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoicePdfResource(context.Background(), invoiceId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoicePdfResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoicePdfResource`: InvoicePDF
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoicePdfResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | Invoice database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePdfResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **string** | An optional fields mask | 

### Return type

[**InvoicePDF**](InvoicePDF.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceResource

> []Invoice GetInvoiceResource(ctx).XFields(xFields).Execute()

Return serialized data on all invoices

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
    resp, r, err := apiClient.InvoiceApi.GetInvoiceResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceResource`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceSearchResource

> InvoiceSeachResult GetInvoiceSearchResource(ctx).EndDate(endDate).StartDate(startDate).PerPage(perPage).Page(page).Q(q).XFields(xFields).Execute()

Return results of invoice search

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
    endDate := int32(56) // int32 | The end date range (optional)
    startDate := int32(56) // int32 | The start date range (optional)
    perPage := int32(56) // int32 | Number of items per page. (optional) (default to 10)
    page := int32(56) // int32 | The page number of search. (optional) (default to 1)
    q := "q_example" // string | Content search. (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoiceSearchResource(context.Background()).EndDate(endDate).StartDate(startDate).PerPage(perPage).Page(page).Q(q).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceSearchResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceSearchResource`: InvoiceSeachResult
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceSearchResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceSearchResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endDate** | **int32** | The end date range | 
 **startDate** | **int32** | The start date range | 
 **perPage** | **int32** | Number of items per page. | [default to 10]
 **page** | **int32** | The page number of search. | [default to 1]
 **q** | **string** | Content search. | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**InvoiceSeachResult**](InvoiceSeachResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceUnpaidResource

> InvoiceUnpaid GetInvoiceUnpaidResource(ctx).XFields(xFields).Execute()

Return total balance of all unpaid invoices

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
    resp, r, err := apiClient.InvoiceApi.GetInvoiceUnpaidResource(context.Background()).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceUnpaidResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceUnpaidResource`: InvoiceUnpaid
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceUnpaidResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceUnpaidResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **string** | An optional fields mask | 

### Return type

[**InvoiceUnpaid**](InvoiceUnpaid.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

