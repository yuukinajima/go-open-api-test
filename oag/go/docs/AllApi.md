# \AllApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthers**](AllApi.md#GetAuthers) | **Get** /authers | Your GET endpoint
[**GetAuthersId**](AllApi.md#GetAuthersId) | **Get** /authers/{id} | Your GET endpoint
[**GetBooks**](AllApi.md#GetBooks) | **Get** /books | Your GET endpoint
[**GetBooksId**](AllApi.md#GetBooksId) | **Get** /books/{id} | Your GET endpoint
[**GetCategories**](AllApi.md#GetCategories) | **Get** /categories | Your GET endpoint
[**GetCategoriesId**](AllApi.md#GetCategoriesId) | **Get** /categories/{id} | Your GET endpoint
[**PostAuthers**](AllApi.md#PostAuthers) | **Post** /authers | 
[**PostBooks**](AllApi.md#PostBooks) | **Post** /books | 



## GetAuthers

> []Auther GetAuthers(ctx, )

Your GET endpoint

get authers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Auther**](auther.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthersId

> Auther GetAuthersId(ctx, id)

Your GET endpoint

get auther info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Auther**](auther.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooks

> []Book GetBooks(ctx, )

Your GET endpoint

book list

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Book**](book.md)

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


## GetCategories

> []Category GetCategories(ctx, )

Your GET endpoint

get categories

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Category**](category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoriesId

> Category GetCategoriesId(ctx, id)

Your GET endpoint

get category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Category**](category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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

