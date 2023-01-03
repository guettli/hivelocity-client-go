/*
Hivelocity API

Interact with Hivelocity

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BillingInfoApiService BillingInfoApi service
type BillingInfoApiService service

type ApiGetBillingInfoResourceRequest struct {
	ctx        context.Context
	ApiService *BillingInfoApiService
	xFields    *string
}

// An optional fields mask
func (r ApiGetBillingInfoResourceRequest) XFields(xFields string) ApiGetBillingInfoResourceRequest {
	r.xFields = &xFields
	return r
}

func (r ApiGetBillingInfoResourceRequest) Execute() ([]BillingInfo, *http.Response, error) {
	return r.ApiService.GetBillingInfoResourceExecute(r)
}

/*
GetBillingInfoResource Return a list with all Billing Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBillingInfoResourceRequest
*/
func (a *BillingInfoApiService) GetBillingInfoResource(ctx context.Context) ApiGetBillingInfoResourceRequest {
	return ApiGetBillingInfoResourceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []BillingInfo
func (a *BillingInfoApiService) GetBillingInfoResourceExecute(r ApiGetBillingInfoResourceRequest) ([]BillingInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BillingInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoApiService.GetBillingInfoResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing-info/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xFields != nil {
		localVarHeaderParams["X-Fields"] = parameterToString(*r.xFields, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostBillingInfoResourceRequest struct {
	ctx        context.Context
	ApiService *BillingInfoApiService
	payload    *BillingInfoVerification
	xFields    *string
}

func (r ApiPostBillingInfoResourceRequest) Payload(payload BillingInfoVerification) ApiPostBillingInfoResourceRequest {
	r.payload = &payload
	return r
}

// An optional fields mask
func (r ApiPostBillingInfoResourceRequest) XFields(xFields string) ApiPostBillingInfoResourceRequest {
	r.xFields = &xFields
	return r
}

func (r ApiPostBillingInfoResourceRequest) Execute() (*BillingInfo, *http.Response, error) {
	return r.ApiService.PostBillingInfoResourceExecute(r)
}

/*
PostBillingInfoResource Create verification for credit card with all Billing Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostBillingInfoResourceRequest
*/
func (a *BillingInfoApiService) PostBillingInfoResource(ctx context.Context) ApiPostBillingInfoResourceRequest {
	return ApiPostBillingInfoResourceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BillingInfo
func (a *BillingInfoApiService) PostBillingInfoResourceExecute(r ApiPostBillingInfoResourceRequest) (*BillingInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BillingInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoApiService.PostBillingInfoResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing-info/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xFields != nil {
		localVarHeaderParams["X-Fields"] = parameterToString(*r.xFields, "")
	}
	// body params
	localVarPostBody = r.payload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutBillingInfoResourceRequest struct {
	ctx        context.Context
	ApiService *BillingInfoApiService
	payload    *BillingInfoVerification
	xFields    *string
}

func (r ApiPutBillingInfoResourceRequest) Payload(payload BillingInfoVerification) ApiPutBillingInfoResourceRequest {
	r.payload = &payload
	return r
}

// An optional fields mask
func (r ApiPutBillingInfoResourceRequest) XFields(xFields string) ApiPutBillingInfoResourceRequest {
	r.xFields = &xFields
	return r
}

func (r ApiPutBillingInfoResourceRequest) Execute() (*Credit, *http.Response, error) {
	return r.ApiService.PutBillingInfoResourceExecute(r)
}

/*
PutBillingInfoResource Verify credit card with all Billing Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutBillingInfoResourceRequest
*/
func (a *BillingInfoApiService) PutBillingInfoResource(ctx context.Context) ApiPutBillingInfoResourceRequest {
	return ApiPutBillingInfoResourceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return Credit
func (a *BillingInfoApiService) PutBillingInfoResourceExecute(r ApiPutBillingInfoResourceRequest) (*Credit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Credit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoApiService.PutBillingInfoResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing-info/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xFields != nil {
		localVarHeaderParams["X-Fields"] = parameterToString(*r.xFields, "")
	}
	// body params
	localVarPostBody = r.payload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
