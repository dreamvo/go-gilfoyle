# \StreamApi

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaPlaylistFile**](StreamApi.md#GetMediaPlaylistFile) | **Get** /medias/{media_id}/stream/{filename} | Get HLS playlist file of a media


# **GetMediaPlaylistFile**
> string GetMediaPlaylistFile(ctx, mediaId, filename)
Get HLS playlist file of a media

Get HLS playlist file of a media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media identifier | 
  **filename** | **string**| HLS filename | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

