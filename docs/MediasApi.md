# \MediasApi

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMedia**](MediasApi.md#CreateMedia) | **Post** /medias | Create a media
[**DeleteMedia**](MediasApi.md#DeleteMedia) | **Delete** /medias/{id} | Delete a media
[**GetAllMedias**](MediasApi.md#GetAllMedias) | **Get** /medias | Query medias
[**GetMedia**](MediasApi.md#GetMedia) | **Get** /medias/{id} | Get a media
[**UpdateMedia**](MediasApi.md#UpdateMedia) | **Patch** /medias/{id} | Update a media
[**UploadMediaFile**](MediasApi.md#UploadMediaFile) | **Post** /medias/{id}/upload | Upload a media file


# **CreateMedia**
> InlineResponse2001 CreateMedia(ctx, media)
Create a media

Create a new media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **media** | [**ApiCreateMedia**](ApiCreateMedia.md)| Media data | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMedia**
> HttputilsDataResponse DeleteMedia(ctx, id)
Delete a media

Delete one media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 

### Return type

[**HttputilsDataResponse**](httputils.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMedias**
> InlineResponse200 GetAllMedias(ctx, optional)
Query medias

Get latest created medias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MediasApiGetAllMediasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediasApiGetAllMediasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Max number of results | 
 **offset** | **optional.Int32**| Number of results to ignore | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMedia**
> InlineResponse2001 GetMedia(ctx, id)
Get a media

Get one media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMedia**
> InlineResponse2001 UpdateMedia(ctx, id, media)
Update a media

Update an existing media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 
  **media** | [**ApiUpdateMedia**](ApiUpdateMedia.md)| Media data | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadMediaFile**
> InlineResponse2001 UploadMediaFile(ctx, id, file)
Upload a media file

Upload a new media file for a given media ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media identifier | 
  **file** | ***os.File**| Media file | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

