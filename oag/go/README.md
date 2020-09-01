# Go API client for openapi

test api


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**DeleteAuthersId**](docs/AdminApi.md#deleteauthersid) | **Delete** /authers/{id} | 
*AdminApi* | [**DeleteBooksId**](docs/AdminApi.md#deletebooksid) | **Delete** /books/{id} | 
*AdminApi* | [**DeleteCategoriesId**](docs/AdminApi.md#deletecategoriesid) | **Delete** /categories/{id} | 
*AdminApi* | [**GetBooksId**](docs/AdminApi.md#getbooksid) | **Get** /books/{id} | Your GET endpoint
*AdminApi* | [**PatchAuthersId**](docs/AdminApi.md#patchauthersid) | **Patch** /authers/{id} | 
*AdminApi* | [**PatchBooksId**](docs/AdminApi.md#patchbooksid) | **Patch** /books/{id} | 
*AdminApi* | [**PatchCategoriesId**](docs/AdminApi.md#patchcategoriesid) | **Patch** /categories/{id} | 
*AdminApi* | [**PostAuthers**](docs/AdminApi.md#postauthers) | **Post** /authers | 
*AdminApi* | [**PostBooks**](docs/AdminApi.md#postbooks) | **Post** /books | 
*AdminApi* | [**PostCategories**](docs/AdminApi.md#postcategories) | **Post** /categories | 
*AllApi* | [**GetAuthers**](docs/AllApi.md#getauthers) | **Get** /authers | Your GET endpoint
*AllApi* | [**GetAuthersId**](docs/AllApi.md#getauthersid) | **Get** /authers/{id} | Your GET endpoint
*AllApi* | [**GetBooks**](docs/AllApi.md#getbooks) | **Get** /books | Your GET endpoint
*AllApi* | [**GetBooksId**](docs/AllApi.md#getbooksid) | **Get** /books/{id} | Your GET endpoint
*AllApi* | [**GetCategories**](docs/AllApi.md#getcategories) | **Get** /categories | Your GET endpoint
*AllApi* | [**GetCategoriesId**](docs/AllApi.md#getcategoriesid) | **Get** /categories/{id} | Your GET endpoint
*AllApi* | [**PostAuthers**](docs/AllApi.md#postauthers) | **Post** /authers | 
*AllApi* | [**PostBooks**](docs/AllApi.md#postbooks) | **Post** /books | 


## Documentation For Models

 - [Auther](docs/Auther.md)
 - [Book](docs/Book.md)
 - [Category](docs/Category.md)
 - [Result](docs/Result.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author


