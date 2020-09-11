# \AdminApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppByAdmin**](AdminApi.md#DeleteAppByAdmin) | **Delete** /admin/apps/{app} | Delete an application
[**DeleteDbByAdmin**](AdminApi.md#DeleteDbByAdmin) | **Delete** /admin/dbs/{db} | Delete a single database
[**DeleteUserByAdmin**](AdminApi.md#DeleteUserByAdmin) | **Delete** /admin/users/{userEmail} | Delete a single user
[**FetchAppByAdmin**](AdminApi.md#FetchAppByAdmin) | **Get** /admin/apps/{app} | Fetch a single application
[**FetchAppsByAdmin**](AdminApi.md#FetchAppsByAdmin) | **Get** /admin/apps | Fetch all applications with/without a filter defined by query params
[**FetchDbByAdmin**](AdminApi.md#FetchDbByAdmin) | **Get** /admin/dbs/{db} | Fetch a single database
[**FetchDbsByAdmin**](AdminApi.md#FetchDbsByAdmin) | **Get** /admin/dbs | Fetch all databases with/without a filter defined by query params
[**FetchNodeByAdmin**](AdminApi.md#FetchNodeByAdmin) | **Get** /admin/nodes/{type} | Fetch bind addresses(IP:Port) of a single microservice on all nodes
[**FetchNodesByAdmin**](AdminApi.md#FetchNodesByAdmin) | **Get** /admin/nodes | Fetch bind addresses(IP:Port) of all microservices on all nodes
[**FetchUserByAdmin**](AdminApi.md#FetchUserByAdmin) | **Get** /admin/users/{userEmail} | Fetch a single user
[**FetchUsersByAdmin**](AdminApi.md#FetchUsersByAdmin) | **Get** /admin/users | Fetch all users with/without a filter defined by query params
[**GrantSuperuserPrivilege**](AdminApi.md#GrantSuperuserPrivilege) | **Patch** /admin/users/{userEmail}/grant | Grant superuser privileges to a single user
[**RevokeSuperuserPrivilege**](AdminApi.md#RevokeSuperuserPrivilege) | **Patch** /admin/users/{userEmail}/revoke | Revoke superuser privileges from a single user



## DeleteAppByAdmin

> InlineResponse2002 DeleteAppByAdmin(ctx, app)

Delete an application

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


## DeleteDbByAdmin

> InlineResponse2002 DeleteDbByAdmin(ctx, db)

Delete a single database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string**| Name of the database | 

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


## DeleteUserByAdmin

> InlineResponse2002 DeleteUserByAdmin(ctx, userEmail)

Delete a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string**| Email ID of the user | 

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


## FetchAppByAdmin

> InlineResponse2003 FetchAppByAdmin(ctx, app)

Fetch a single application

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


## FetchAppsByAdmin

> InlineResponse2003 FetchAppsByAdmin(ctx, optional)

Fetch all applications with/without a filter defined by query params

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchAppsByAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchAppsByAdminOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Name of the application | 
 **owner** | **optional.String**| Owner of the application | 
 **language** | **optional.String**| Language in which the application is written | 
 **dockerImage** | **optional.String**| Docker Image used in building the application&#39;s container | 
 **hostIp** | **optional.String**| IPv4 address of the node in which the application is deployed | 
 **gitUrl** | **optional.String**| Application&#39;s Git Repository URL | 
 **containerPort** | **optional.String**| Port assigned by the node to the application&#39;s docker container | 

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


## FetchDbByAdmin

> InlineResponse2007 FetchDbByAdmin(ctx, db)

Fetch a single database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string**| Name of the database | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDbsByAdmin

> InlineResponse2007 FetchDbsByAdmin(ctx, optional)

Fetch all databases with/without a filter defined by query params

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchDbsByAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchDbsByAdminOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Name of the database | 
 **user** | **optional.String**| User of the database | 
 **owner** | **optional.String**| Owner of the database | 
 **language** | **optional.String**| Type of the database | 
 **hostIp** | **optional.String**| IPv4 address of the node in which the database is deployed | 
 **containerPort** | **optional.String**| Port assigned by the node to the database&#39;s docker container | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNodeByAdmin

> InlineResponse20013 FetchNodeByAdmin(ctx, type_)

Fetch bind addresses(IP:Port) of a single microservice on all nodes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Type of microservice | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNodesByAdmin

> InlineResponse20012 FetchNodesByAdmin(ctx, )

Fetch bind addresses(IP:Port) of all microservices on all nodes

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserByAdmin

> InlineResponse20011 FetchUserByAdmin(ctx, userEmail)

Fetch a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string**| Email ID of the user | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsersByAdmin

> InlineResponse20011 FetchUsersByAdmin(ctx, optional)

Fetch all users with/without a filter defined by query params

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchUsersByAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUsersByAdminOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Name of the user | 
 **email** | **optional.String**| Email of the user | 
 **admin** | **optional.Bool**| Field denoting superuser privileges | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantSuperuserPrivilege

> InlineResponse2002 GrantSuperuserPrivilege(ctx, userEmail)

Grant superuser privileges to a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string**| Email ID of the user | 

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


## RevokeSuperuserPrivilege

> InlineResponse2002 RevokeSuperuserPrivilege(ctx, userEmail)

Revoke superuser privileges from a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string**| Email ID of the user | 

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

