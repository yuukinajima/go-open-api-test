# \AdminApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthersId**](AdminApi.md#DeleteAuthersId) | **Delete** /authers/{id} | 
[**DeleteBooksId**](AdminApi.md#DeleteBooksId) | **Delete** /books/{id} | 
[**DeleteCategoriesId**](AdminApi.md#DeleteCategoriesId) | **Delete** /categories/{id} | 
[**GetBooksId**](AdminApi.md#GetBooksId) | **Get** /books/{id} | Your GET endpoint
[**PatchAuthersId**](AdminApi.md#PatchAuthersId) | **Patch** /authers/{id} | 
[**PatchBooksId**](AdminApi.md#PatchBooksId) | **Patch** /books/{id} | 
[**PatchCategoriesId**](AdminApi.md#PatchCategoriesId) | **Patch** /categories/{id} | 
[**PostAuthers**](AdminApi.md#PostAuthers) | **Post** /authers | 
[**PostBooks**](AdminApi.md#PostBooks) | **Post** /books | 
[**PostCategories**](AdminApi.md#PostCategories) | **Post** /categories | 



## DeleteAuthersId

> Result DeleteAuthersId(ctx, id)



delete auther

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Result**](result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBooksId

> Result DeleteBooksId(ctx, id)



delete book

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Result**](result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategoriesId

> Result DeleteCategoriesId(ctx, id)



delete category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Result**](result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksId

> Book GetBooksId(ctx, id)

Your GET endpoint

get book info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Book**](book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthersId

> Auther PatchAuthersId(ctx, id, optional)



update auther info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***PatchAuthersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchAuthersIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auther** | [**optional.Interface of Auther**](Auther.md)|  | 

### Return type

[**Auther**](auther.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBooksId

> Book PatchBooksId(ctx, id, optional)



update book info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***PatchBooksIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchBooksIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **book** | [**optional.Interface of Book**](Book.md)|  | 

### Return type

[**Book**](book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCategoriesId

> Category PatchCategoriesId(ctx, id, optional)



update category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***PatchCategoriesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchCategoriesIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | [**optional.Interface of Category**](Category.md)|  | 

### Return type

[**Category**](category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthers

> Auther PostAuthers(ctx, optional)



create auther

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAuthersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAuthersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auther** | [**optional.Interface of Auther**](Auther.md)|  | 

### Return type

[**Auther**](auther.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooks

> Book PostBooks(ctx, optional)



create new book

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostBooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostBooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **book** | [**optional.Interface of Book**](Book.md)|  | 

### Return type

[**Book**](book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCategories

> Category PostCategories(ctx, optional)



create categories

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**optional.Interface of Category**](Category.md)|  | 

### Return type

[**Category**](category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

