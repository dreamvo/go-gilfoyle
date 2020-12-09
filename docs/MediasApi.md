# \MediasApi

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMedia**](MediasApi.md#CreateMedia) | **Post** /medias | Create a media
[**DeleteMedia**](MediasApi.md#DeleteMedia) | **Delete** /medias/{id} | Delete a media
[**GetAllMedias**](MediasApi.md#GetAllMedias) | **Get** /medias | Query medias
[**GetMedia**](MediasApi.md#GetMedia) | **Get** /medias/{id} | Get a media
[**UpdateMedia**](MediasApi.md#UpdateMedia) | **Patch** /medias/{id} | Update a media
[**UploadAudio**](MediasApi.md#UploadAudio) | **Post** /medias/{id}/upload/audio | Upload a audio file
[**UploadVideo**](MediasApi.md#UploadVideo) | **Post** /medias/{id}/upload/video | Upload a video file


# **CreateMedia**
> InlineResponse2002 CreateMedia(ctx, media)
Create a media

Create a new media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **media** | [**ApiCreateMedia**](ApiCreateMedia.md)| Media data | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMedia**
> UtilDataResponse DeleteMedia(ctx, id)
Delete a media

Delete one media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 

### Return type

[**UtilDataResponse**](util.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMedias**
> InlineResponse2001 GetAllMedias(ctx, optional)
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

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMedia**
> InlineResponse2002 GetMedia(ctx, id)
Get a media

Get one media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMedia**
> InlineResponse2002 UpdateMedia(ctx, id, media)
Update a media

Update an existing media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media ID | 
  **media** | [**ApiUpdateMedia**](ApiUpdateMedia.md)| Media data | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadAudio**
> InlineResponse2003 UploadAudio(ctx, id, file)
Upload a audio file

Upload a new audio file for a given media ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media identifier | 
  **file** | ***os.File**| Audio file | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadVideo**
> InlineResponse2003 UploadVideo(ctx, id, file)
Upload a video file

Upload a new video file for a given media ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Media identifier | 
  **file** | ***os.File**| Video file | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

