# \DbsAPI

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDB**](DbsAPI.md#CreateDB) | **Post** /dbs/{databaseType} | Create a database
[**DeleteDbByUser**](DbsAPI.md#DeleteDbByUser) | **Delete** /dbs/{db} | Delete a single database owned by a user
[**FetchDbByUser**](DbsAPI.md#FetchDbByUser) | **Get** /dbs/{db} | Fetch a single database owned by a user
[**FetchDbsByUser**](DbsAPI.md#FetchDbsByUser) | **Get** /dbs | Fetch all databases owned by a user


## CreateDB

> InlineResponse2002 CreateDB(ctx, databaseType, optional)

Create a database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseType** | **string**| The type of database | 
 **optional** | ***CreateDBOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDBOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | [**optional.Interface of Database**](Database.md)|  | 

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


## DeleteDbByUser

> InlineResponse2002 DeleteDbByUser(ctx, db)

Delete a single database owned by a user

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


## FetchDbByUser

> InlineResponse2007 FetchDbByUser(ctx, db)

Fetch a single database owned by a user

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


## FetchDbsByUser

> InlineResponse2007 FetchDbsByUser(ctx, )

Fetch all databases owned by a user

### Required Parameters

This endpoint does not need any parameter.

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
