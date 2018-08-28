# \StorageApi

All URIs are relative to *https://vpc.speakin.mobi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Download**](StorageApi.md#Download) | **Get** /cloudapi/v1beta/storage/download | 
[**Upload**](StorageApi.md#Upload) | **Post** /cloudapi/v1beta/storage/upload | 


# **Download**
> *os.File Download(ctx, bucket, key)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucket** | **string**|  | 
  **key** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Upload**
> CallUploadResponse Upload(ctx, bucket, type_, durationMS, timestamp, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucket** | **string**|  | 
  **type_** | **string**|  | 
  **durationMS** | **int64**|  | 
  **timestamp** | **int64**|  | 
 **optional** | ***UploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**CallUploadResponse**](CallUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

