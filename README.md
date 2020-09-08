# Go API client for openapi

Video streaming server backed by decentralized filesystem.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1-beta
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

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HealthApi* | [**HealthGet**](docs/HealthApi.md#healthget) | **Get** /health | Check service status
*VideosApi* | [**VideosGet**](docs/VideosApi.md#videosget) | **Get** /videos | Query videos
*VideosApi* | [**VideosIdDelete**](docs/VideosApi.md#videosiddelete) | **Delete** /videos/{id} | Delete a video
*VideosApi* | [**VideosIdGet**](docs/VideosApi.md#videosidget) | **Get** /videos/{id} | Get a video
*VideosApi* | [**VideosIdPatch**](docs/VideosApi.md#videosidpatch) | **Patch** /videos/{id} | Update a video
*VideosApi* | [**VideosIdUploadPost**](docs/VideosApi.md#videosiduploadpost) | **Post** /videos/{id}/upload | Upload a video file
*VideosApi* | [**VideosPost**](docs/VideosApi.md#videospost) | **Post** /videos | Create a video


## Documentation For Models

 - [EntVideo](docs/EntVideo.md)
 - [HttputilsDataResponse](docs/HttputilsDataResponse.md)
 - [HttputilsErrorResponse](docs/HttputilsErrorResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



