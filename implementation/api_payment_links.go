/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-01-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// PaymentLinksApiService PaymentLinksApi service
type PaymentLinksApiService service

type ApiCreatePaymentLinkRequest struct {
	ctx context.Context
	ApiService *PaymentLinksApiService
	xClientId *string
	xClientSecret *string
	xApiVersion *string
	xIdempotencyReplayed *bool
	xIdempotencyKey *string
	xRequestId *string
	cFLinkRequest *CFLinkRequest
}

func (r ApiCreatePaymentLinkRequest) XClientId(xClientId string) ApiCreatePaymentLinkRequest {
	r.xClientId = &xClientId
	return r
}
func (r ApiCreatePaymentLinkRequest) XClientSecret(xClientSecret string) ApiCreatePaymentLinkRequest {
	r.xClientSecret = &xClientSecret
	return r
}
func (r ApiCreatePaymentLinkRequest) XApiVersion(xApiVersion string) ApiCreatePaymentLinkRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiCreatePaymentLinkRequest) XIdempotencyReplayed(xIdempotencyReplayed bool) ApiCreatePaymentLinkRequest {
	r.xIdempotencyReplayed = &xIdempotencyReplayed
	return r
}
func (r ApiCreatePaymentLinkRequest) XIdempotencyKey(xIdempotencyKey string) ApiCreatePaymentLinkRequest {
	r.xIdempotencyKey = &xIdempotencyKey
	return r
}
func (r ApiCreatePaymentLinkRequest) XRequestId(xRequestId string) ApiCreatePaymentLinkRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreatePaymentLinkRequest) CFLinkRequest(cFLinkRequest CFLinkRequest) ApiCreatePaymentLinkRequest {
	r.cFLinkRequest = &cFLinkRequest
	return r
}

func (r ApiCreatePaymentLinkRequest) Execute() (*CFLink, *http.Response, error) {
	return r.ApiService.CreatePaymentLinkExecute(r)
}

/*
CreatePaymentLink Create Payment Link

Use this API to create a new payment link. The created payment link url will be available in the API response parameter link_url.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePaymentLinkRequest
*/
func (a *PaymentLinksApiService) CreatePaymentLink(ctx context.Context) ApiCreatePaymentLinkRequest {
	return ApiCreatePaymentLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CFLink
func (a *PaymentLinksApiService) CreatePaymentLinkExecute(r ApiCreatePaymentLinkRequest) (*CFLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CFLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksApiService.CreatePaymentLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xClientId == nil {
		return localVarReturnValue, nil, reportError("xClientId is required and must be specified")
	}
	if r.xClientSecret == nil {
		return localVarReturnValue, nil, reportError("xClientSecret is required and must be specified")
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
	localVarHeaderParams["x-client-id"] = parameterToString(*r.xClientId, "")
	localVarHeaderParams["x-client-secret"] = parameterToString(*r.xClientSecret, "")
	if r.xApiVersion != nil {
		localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	}
	if r.xIdempotencyReplayed != nil {
		localVarHeaderParams["x-idempotency-replayed"] = parameterToString(*r.xIdempotencyReplayed, "")
	}
	if r.xIdempotencyKey != nil {
		localVarHeaderParams["x-idempotency-key"] = parameterToString(*r.xIdempotencyKey, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	}
	// body params
	localVarPostBody = r.cFLinkRequest
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

type ApiGetPaymentLinkDetailsRequest struct {
	ctx context.Context
	ApiService *PaymentLinksApiService
	xClientId *string
	xClientSecret *string
	linkId string
	xApiVersion *string
	xIdempotencyReplayed *bool
	xIdempotencyKey *string
	xRequestId *string
}

func (r ApiGetPaymentLinkDetailsRequest) XClientId(xClientId string) ApiGetPaymentLinkDetailsRequest {
	r.xClientId = &xClientId
	return r
}
func (r ApiGetPaymentLinkDetailsRequest) XClientSecret(xClientSecret string) ApiGetPaymentLinkDetailsRequest {
	r.xClientSecret = &xClientSecret
	return r
}
func (r ApiGetPaymentLinkDetailsRequest) XApiVersion(xApiVersion string) ApiGetPaymentLinkDetailsRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetPaymentLinkDetailsRequest) XIdempotencyReplayed(xIdempotencyReplayed bool) ApiGetPaymentLinkDetailsRequest {
	r.xIdempotencyReplayed = &xIdempotencyReplayed
	return r
}
func (r ApiGetPaymentLinkDetailsRequest) XIdempotencyKey(xIdempotencyKey string) ApiGetPaymentLinkDetailsRequest {
	r.xIdempotencyKey = &xIdempotencyKey
	return r
}
func (r ApiGetPaymentLinkDetailsRequest) XRequestId(xRequestId string) ApiGetPaymentLinkDetailsRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiGetPaymentLinkDetailsRequest) Execute() (*CFLink, *http.Response, error) {
	return r.ApiService.GetPaymentLinkDetailsExecute(r)
}

/*
GetPaymentLinkDetails Fetch Payment Link Details

Use this API to view all details and status of a payment link.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param linkId
 @return ApiGetPaymentLinkDetailsRequest
*/
func (a *PaymentLinksApiService) GetPaymentLinkDetails(ctx context.Context, linkId string) ApiGetPaymentLinkDetailsRequest {
	return ApiGetPaymentLinkDetailsRequest{
		ApiService: a,
		ctx: ctx,
		linkId: linkId,
	}
}

// Execute executes the request
//  @return CFLink
func (a *PaymentLinksApiService) GetPaymentLinkDetailsExecute(r ApiGetPaymentLinkDetailsRequest) (*CFLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CFLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksApiService.GetPaymentLinkDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/links/{link_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"link_id"+"}", url.PathEscape(parameterToString(r.linkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xClientId == nil {
		return localVarReturnValue, nil, reportError("xClientId is required and must be specified")
	}
	if r.xClientSecret == nil {
		return localVarReturnValue, nil, reportError("xClientSecret is required and must be specified")
	}

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
	localVarHeaderParams["x-client-id"] = parameterToString(*r.xClientId, "")
	localVarHeaderParams["x-client-secret"] = parameterToString(*r.xClientSecret, "")
	if r.xApiVersion != nil {
		localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	}
	if r.xIdempotencyReplayed != nil {
		localVarHeaderParams["x-idempotency-replayed"] = parameterToString(*r.xIdempotencyReplayed, "")
	}
	if r.xIdempotencyKey != nil {
		localVarHeaderParams["x-idempotency-key"] = parameterToString(*r.xIdempotencyKey, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
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

type ApiGetPaymentLinkOrdersRequest struct {
	ctx context.Context
	ApiService *PaymentLinksApiService
	xClientId *string
	xClientSecret *string
	linkId string
	xApiVersion *string
	xIdempotencyReplayed *bool
	xIdempotencyKey *string
	xRequestId *string
}

func (r ApiGetPaymentLinkOrdersRequest) XClientId(xClientId string) ApiGetPaymentLinkOrdersRequest {
	r.xClientId = &xClientId
	return r
}
func (r ApiGetPaymentLinkOrdersRequest) XClientSecret(xClientSecret string) ApiGetPaymentLinkOrdersRequest {
	r.xClientSecret = &xClientSecret
	return r
}
func (r ApiGetPaymentLinkOrdersRequest) XApiVersion(xApiVersion string) ApiGetPaymentLinkOrdersRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetPaymentLinkOrdersRequest) XIdempotencyReplayed(xIdempotencyReplayed bool) ApiGetPaymentLinkOrdersRequest {
	r.xIdempotencyReplayed = &xIdempotencyReplayed
	return r
}
func (r ApiGetPaymentLinkOrdersRequest) XIdempotencyKey(xIdempotencyKey string) ApiGetPaymentLinkOrdersRequest {
	r.xIdempotencyKey = &xIdempotencyKey
	return r
}
func (r ApiGetPaymentLinkOrdersRequest) XRequestId(xRequestId string) ApiGetPaymentLinkOrdersRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiGetPaymentLinkOrdersRequest) Execute() ([]CFLinkOrders, *http.Response, error) {
	return r.ApiService.GetPaymentLinkOrdersExecute(r)
}

/*
GetPaymentLinkOrders Get Orders for a Payment Link

Use this API to view all order details for a payment link.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param linkId
 @return ApiGetPaymentLinkOrdersRequest
*/
func (a *PaymentLinksApiService) GetPaymentLinkOrders(ctx context.Context, linkId string) ApiGetPaymentLinkOrdersRequest {
	return ApiGetPaymentLinkOrdersRequest{
		ApiService: a,
		ctx: ctx,
		linkId: linkId,
	}
}

// Execute executes the request
//  @return []CFLinkOrders
func (a *PaymentLinksApiService) GetPaymentLinkOrdersExecute(r ApiGetPaymentLinkOrdersRequest) ([]CFLinkOrders, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CFLinkOrders
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksApiService.GetPaymentLinkOrders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/links/{link_id}/orders"
	localVarPath = strings.Replace(localVarPath, "{"+"link_id"+"}", url.PathEscape(parameterToString(r.linkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xClientId == nil {
		return localVarReturnValue, nil, reportError("xClientId is required and must be specified")
	}
	if r.xClientSecret == nil {
		return localVarReturnValue, nil, reportError("xClientSecret is required and must be specified")
	}

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
	localVarHeaderParams["x-client-id"] = parameterToString(*r.xClientId, "")
	localVarHeaderParams["x-client-secret"] = parameterToString(*r.xClientSecret, "")
	if r.xApiVersion != nil {
		localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	}
	if r.xIdempotencyReplayed != nil {
		localVarHeaderParams["x-idempotency-replayed"] = parameterToString(*r.xIdempotencyReplayed, "")
	}
	if r.xIdempotencyKey != nil {
		localVarHeaderParams["x-idempotency-key"] = parameterToString(*r.xIdempotencyKey, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
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