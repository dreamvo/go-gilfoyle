# \AttachmentsApi

All URIs are relative to *http://demo-v1.gilfoyle.dreamvo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaAttachment**](AttachmentsApi.md#AddMediaAttachment) | **Post** /medias/{media_id}/attachments | Add attachment to a media
[**DeleteMediaAttachment**](AttachmentsApi.md#DeleteMediaAttachment) | **Delete** /medias/{media_id}/attachments/{key} | Delete attachment of a media
[**GetMediaAttachments**](AttachmentsApi.md#GetMediaAttachments) | **Get** /medias/{media_id}/attachments | Get attachments of a media


# **AddMediaAttachment**
> UtilDataResponse AddMediaAttachment(ctx, mediaId, file)
Add attachment to a media

Add attachment to a media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media identifier | 
  **file** | ***os.File**| Attachment file | 

### Return type

[**UtilDataResponse**](util.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMediaAttachment**
> UtilDataResponse DeleteMediaAttachment(ctx, mediaId, key)
Delete attachment of a media

Delete attachment of a media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media identifier | 
  **key** | **string**| Attachment unique identifier | 

### Return type

[**UtilDataResponse**](util.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaAttachments**
> UtilDataResponse GetMediaAttachments(ctx, mediaId)
Get attachments of a media

Get attachments of a media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media identifier | 

### Return type

[**UtilDataResponse**](util.DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

