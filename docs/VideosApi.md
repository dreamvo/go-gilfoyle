# \VideosApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1VideosGet**](VideosApi.md#V1VideosGet) | **Get** /v1/videos | Query videos
[**V1VideosIdDelete**](VideosApi.md#V1VideosIdDelete) | **Delete** /v1/videos/{id} | Delete a video
[**V1VideosIdGet**](VideosApi.md#V1VideosIdGet) | **Get** /v1/videos/{id} | Get a video
[**V1VideosIdPatch**](VideosApi.md#V1VideosIdPatch) | **Patch** /v1/videos/{id} | Update a video
[**V1VideosIdUploadPost**](VideosApi.md#V1VideosIdUploadPost) | **Post** /v1/videos/{id}/upload | Upload a video file
[**V1VideosPost**](VideosApi.md#V1VideosPost) | **Post** /v1/videos | Create a video


# **V1VideosGet**
> InlineResponse200 V1VideosGet(ctx, optional)
Query videos

get latest videos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VideosApiV1VideosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideosApiV1VideosGetOpts struct

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

# **V1VideosIdDelete**
> HttputilsDataResponse V1VideosIdDelete(ctx, id)
Delete a video

Delete one video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Video ID | 

### Return type

[**HttputilsDataResponse**](httputils.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VideosIdGet**
> InlineResponse2001 V1VideosIdGet(ctx, id)
Get a video

get one video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Video ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VideosIdPatch**
> InlineResponse2001 V1VideosIdPatch(ctx, id, title)
Update a video

Update an existing video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Video ID | 
  **title** | **string**| Video title | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VideosIdUploadPost**
> InlineResponse2001 V1VideosIdUploadPost(ctx, id, file)
Upload a video file

Upload a new video file for a given video ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Video ID | 
  **file** | ***os.File**| Video file | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VideosPost**
> InlineResponse2001 V1VideosPost(ctx, title)
Create a video

Create a new video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **title** | **string**| Video title | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

