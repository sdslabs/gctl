# \AppsAPI

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsAPI.md#CreateApp) | **Post** /apps/{language} | Create an application
[**DeleteAppByUser**](AppsAPI.md#DeleteAppByUser) | **Delete** /apps/{app} | Delete an application owned by a user
[**FetchAppByUser**](AppsAPI.md#FetchAppByUser) | **Get** /apps/{app} | Fetch a single application owned by a user
[**FetchAppsByUser**](AppsAPI.md#FetchAppsByUser) | **Get** /apps | Fetch all applications owned by a user
[**FetchLogsByUser**](AppsAPI.md#FetchLogsByUser) | **Get** /apps/{app}/logs | Fetch logs of an application
[**FetchMetricsByUser**](AppsAPI.md#FetchMetricsByUser) | **Get** /apps/{app}/metrics | Fetch metrics of an application
[**RebuildAppByUser**](AppsAPI.md#RebuildAppByUser) | **Patch** /apps/{app}/rebuild | Rebuild an application
[**UpdateAppByUser**](AppsAPI.md#UpdateAppByUser) | **Put** /apps/{app} | Update an application owned by a user



## CreateApp

> InlineResponse2002 CreateApp(ctx, language, application)

Create an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string**| The programming language in which the application is written | 
**application** | [**Application**](Application.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppByUser

> InlineResponse2002 DeleteAppByUser(ctx, app)

Delete an application owned by a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAppByUser

> InlineResponse2003 FetchAppByUser(ctx, app)

Fetch a single application owned by a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAppsByUser

> InlineResponse2003 FetchAppsByUser(ctx, )

Fetch all applications owned by a user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLogsByUser

> InlineResponse2005 FetchLogsByUser(ctx, app, optional)

Fetch logs of an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 
 **optional** | ***FetchLogsByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchLogsByUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tail** | **optional.Int32**| Fetch the last **n** logs (Fetches all logs if not specified) | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMetricsByUser

> InlineResponse2006 FetchMetricsByUser(ctx, app, optional)

Fetch metrics of an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 
 **optional** | ***FetchMetricsByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchMetricsByUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seconds** | **optional.Int32**| Fetch metrics in the last **n** seconds | 
 **minutes** | **optional.Int32**| Fetch metrics in the last **n** minutes | 
 **hours** | **optional.Int32**| Fetch metrics in the last **n** hours | 
 **days** | **optional.Int32**| Fetch metrics in the last **n** days | 
 **weeks** | **optional.Int32**| Fetch metrics in the last **n** weeks | 
 **months** | **optional.Int32**| Fetch metrics in the last **n** months | 
 **years** | **optional.Int32**| Fetch metrics in the last **n** years | 
 **decades** | **optional.Int32**| Fetch metrics in the last **n** decades | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildAppByUser

> InlineResponse2002 RebuildAppByUser(ctx, app)

Rebuild an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppByUser

> InlineResponse2002 UpdateAppByUser(ctx, app, application)

Update an application owned by a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the application | 
**application** | [**Application**](Application.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

