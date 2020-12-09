# \StreamApi

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamMedia**](StreamApi.md#StreamMedia) | **Get** /medias/{media_id}/stream/{preset} | Get stream from media file


# **StreamMedia**
> UtilDataResponse StreamMedia(ctx, mediaId, preset)
Get stream from media file

Get stream from media file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media identifier | 
  **preset** | **string**| Encoder preset | 

### Return type

[**UtilDataResponse**](util.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

