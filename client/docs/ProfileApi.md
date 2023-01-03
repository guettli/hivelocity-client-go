# \ProfileApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicProfileResource**](ProfileApi.md#GetBasicProfileResource) | **Get** /profile/basic | Get Basic Profile of current user or a contact with id
[**GetProfileResource**](ProfileApi.md#GetProfileResource) | **Get** /profile/ | Get Profile of current user or a contact with id
[**PutProfileResource**](ProfileApi.md#PutProfileResource) | **Put** /profile/ | Update Profile of current user or a contact with id



## GetBasicProfileResource

> BasicProfile GetBasicProfileResource(ctx).ContactId(contactId).XFields(xFields).Execute()

Get Basic Profile of current user or a contact with id



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
    contactId := "contactId_example" // string | \"ID of Contact to manage Basic Profile\" (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.GetBasicProfileResource(context.Background()).ContactId(contactId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetBasicProfileResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicProfileResource`: BasicProfile
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetBasicProfileResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicProfileResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** | \&quot;ID of Contact to manage Basic Profile\&quot; | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**BasicProfile**](BasicProfile.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileResource

> Profile GetProfileResource(ctx).ContactId(contactId).XFields(xFields).Execute()

Get Profile of current user or a contact with id



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
    contactId := "contactId_example" // string | \"ID of Contact to manage Profile\" (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.GetProfileResource(context.Background()).ContactId(contactId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileResource`: Profile
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** | \&quot;ID of Contact to manage Profile\&quot; | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Profile**](Profile.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProfileResource

> Profile PutProfileResource(ctx).Payload(payload).ContactId(contactId).XFields(xFields).Execute()

Update Profile of current user or a contact with id



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
    payload := *openapiclient.NewProfileUpdate() // ProfileUpdate | 
    contactId := "contactId_example" // string | \"ID of Contact to manage Profile\" (optional)
    xFields := "xFields_example" // string | An optional fields mask (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.PutProfileResource(context.Background()).Payload(payload).ContactId(contactId).XFields(xFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.PutProfileResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProfileResource`: Profile
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.PutProfileResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProfileResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ProfileUpdate**](ProfileUpdate.md) |  | 
 **contactId** | **string** | \&quot;ID of Contact to manage Profile\&quot; | 
 **xFields** | **string** | An optional fields mask | 

### Return type

[**Profile**](Profile.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

