# \UserApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /user | Delete user
[**FetchUser**](UserApi.md#FetchUser) | **Get** /user | Fetch logged in user&#39;s info
[**UpdatePassword**](UserApi.md#UpdatePassword) | **Put** /user/password | Update the password of the logged in user



## DeleteUser

> InlineResponse2009 DeleteUser(ctx, )

Delete user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> InlineResponse2008 FetchUser(ctx, )

Fetch logged in user's info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassword

> InlineResponse20010 UpdatePassword(ctx, inlineObject)

Update the password of the logged in user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

