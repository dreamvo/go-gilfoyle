# Go API client for swagger

Media streaming server backed by decentralized filesystem.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.1-beta
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HealthApi* | [**CheckHealth**](docs/HealthApi.md#checkhealth) | **Get** /health | Check service status
*MediasApi* | [**CreateMedia**](docs/MediasApi.md#createmedia) | **Post** /medias | Create a media
*MediasApi* | [**DeleteMedia**](docs/MediasApi.md#deletemedia) | **Delete** /medias/{id} | Delete a media
*MediasApi* | [**GetAllMedias**](docs/MediasApi.md#getallmedias) | **Get** /medias | Query medias
*MediasApi* | [**GetMedia**](docs/MediasApi.md#getmedia) | **Get** /medias/{id} | Get a media
*MediasApi* | [**UpdateMedia**](docs/MediasApi.md#updatemedia) | **Patch** /medias/{id} | Update a media
*MediasApi* | [**UploadMediaFile**](docs/MediasApi.md#uploadmediafile) | **Post** /medias/{id}/upload | Upload a media file


## Documentation For Models

 - [ApiCreateMedia](docs/ApiCreateMedia.md)
 - [ApiUpdateMedia](docs/ApiUpdateMedia.md)
 - [EntMedia](docs/EntMedia.md)
 - [HttputilsDataResponse](docs/HttputilsDataResponse.md)
 - [HttputilsErrorResponse](docs/HttputilsErrorResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



