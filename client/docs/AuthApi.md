# \AuthApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Refresh**](AuthApi.md#Refresh) | **Get** /auth/refresh | Refresh JWT token using existing token


## Refresh

> LoginResponse Refresh(ctx, authorization)

Refresh JWT token using existing token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Bearer Token Authentication | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
