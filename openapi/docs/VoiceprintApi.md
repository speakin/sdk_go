# \VoiceprintApi

All URIs are relative to *https://vpc.speakin.mobi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Agc**](VoiceprintApi.md#Agc) | **Post** /cloudapi/v1beta/voiceprint/agc | 
[**Clips**](VoiceprintApi.md#Clips) | **Post** /cloudapi/v1beta/voiceprint/clips | 
[**Ctcdasr**](VoiceprintApi.md#Ctcdasr) | **Post** /cloudapi/v1beta/voiceprint/ctcdasr | 
[**Delete**](VoiceprintApi.md#Delete) | **Post** /cloudapi/v1beta/voiceprint/delete | 
[**Denoise**](VoiceprintApi.md#Denoise) | **Post** /cloudapi/v1beta/voiceprint/denoise | 
[**Dynamicthreshold**](VoiceprintApi.md#Dynamicthreshold) | **Post** /cloudapi/v1beta/voiceprint/dynamicthreshold | 
[**Mos**](VoiceprintApi.md#Mos) | **Post** /cloudapi/v1beta/voiceprint/mos | 
[**Normalize**](VoiceprintApi.md#Normalize) | **Post** /cloudapi/v1beta/voiceprint/normalize | 
[**Query**](VoiceprintApi.md#Query) | **Post** /cloudapi/v1beta/voiceprint/query | 
[**Register**](VoiceprintApi.md#Register) | **Post** /cloudapi/v1beta/voiceprint/register | 
[**Split**](VoiceprintApi.md#Split) | **Post** /cloudapi/v1beta/voiceprint/split | 
[**Threshold**](VoiceprintApi.md#Threshold) | **Post** /cloudapi/v1beta/voiceprint/threshold | 
[**Vadcheck**](VoiceprintApi.md#Vadcheck) | **Post** /cloudapi/v1beta/voiceprint/vadcheck | 
[**Verify**](VoiceprintApi.md#Verify) | **Post** /cloudapi/v1beta/voiceprint/verify | 
[**Verify1ton**](VoiceprintApi.md#Verify1ton) | **Post** /cloudapi/v1beta/voiceprint/verify1ton | 
[**VerifyMulti**](VoiceprintApi.md#VerifyMulti) | **Post** /cloudapi/v1beta/voiceprint/verify_multi | 
[**Verifytopn**](VoiceprintApi.md#Verifytopn) | **Post** /cloudapi/v1beta/voiceprint/verifytopn | 


# **Agc**
> RespVoiceprintAgcResponse Agc(ctx, optional)


增益

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgcOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintAgcRequest** | [**optional.Interface of VoiceprintAgcRequest**](VoiceprintAgcRequest.md)|  | 

### Return type

[**RespVoiceprintAgcResponse**](RespVoiceprintAgcResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Clips**
> RespVoiceprintClipsResponse Clips(ctx, optional)


有效音

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClipsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintClipsRequest** | [**optional.Interface of VoiceprintClipsRequest**](VoiceprintClipsRequest.md)|  | 

### Return type

[**RespVoiceprintClipsResponse**](RespVoiceprintClipsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ctcdasr**
> RespVoiceprintCtcdasrResponse Ctcdasr(ctx, optional)


数字asr

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CtcdasrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CtcdasrOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintCtcdasrRequest** | [**optional.Interface of VoiceprintCtcdasrRequest**](VoiceprintCtcdasrRequest.md)|  | 

### Return type

[**RespVoiceprintCtcdasrResponse**](RespVoiceprintCtcdasrResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> RespVoiceprintDeleteResponse Delete(ctx, optional)


声纹查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintDeleteRequest** | [**optional.Interface of VoiceprintDeleteRequest**](VoiceprintDeleteRequest.md)|  | 

### Return type

[**RespVoiceprintDeleteResponse**](RespVoiceprintDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Denoise**
> RespVoiceprintDenoiseResponse Denoise(ctx, optional)


降噪算法

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DenoiseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DenoiseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintDenoiseRequest** | [**optional.Interface of VoiceprintDenoiseRequest**](VoiceprintDenoiseRequest.md)|  | 

### Return type

[**RespVoiceprintDenoiseResponse**](RespVoiceprintDenoiseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Dynamicthreshold**
> RespVoiceprintDynamicThresholdResponse Dynamicthreshold(ctx, optional)


查询动态阈值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DynamicthresholdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicthresholdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintDynamicThresholdRequest** | [**optional.Interface of VoiceprintDynamicThresholdRequest**](VoiceprintDynamicThresholdRequest.md)|  | 

### Return type

[**RespVoiceprintDynamicThresholdResponse**](RespVoiceprintDynamicThresholdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Mos**
> RespVoiceprintMosResponse Mos(ctx, optional)


计算mos值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MosOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintMosRequest** | [**optional.Interface of VoiceprintMosRequest**](VoiceprintMosRequest.md)|  | 

### Return type

[**RespVoiceprintMosResponse**](RespVoiceprintMosResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Normalize**
> RespVoiceprintNormalizeResponse Normalize(ctx, optional)


分数归一化

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NormalizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NormalizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintNormalizeRequest** | [**optional.Interface of VoiceprintNormalizeRequest**](VoiceprintNormalizeRequest.md)|  | 

### Return type

[**RespVoiceprintNormalizeResponse**](RespVoiceprintNormalizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Query**
> RespVoiceprintQueryResponse Query(ctx, optional)


声纹查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintQueryRequest** | [**optional.Interface of VoiceprintQueryRequest**](VoiceprintQueryRequest.md)|  | 

### Return type

[**RespVoiceprintQueryResponse**](RespVoiceprintQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Register**
> RespVoiceprintRegisterResponse Register(ctx, optional)


声纹注册

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintRegisterRequest** | [**optional.Interface of VoiceprintRegisterRequest**](VoiceprintRegisterRequest.md)|  | 

### Return type

[**RespVoiceprintRegisterResponse**](RespVoiceprintRegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Split**
> RespVoiceprintSplitResponse Split(ctx, optional)


人声分离

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SplitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintSplitRequest** | [**optional.Interface of VoiceprintSplitRequest**](VoiceprintSplitRequest.md)|  | 

### Return type

[**RespVoiceprintSplitResponse**](RespVoiceprintSplitResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Threshold**
> RespVoiceprintThresholdResponse Threshold(ctx, optional)


查询阈值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ThresholdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ThresholdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintThresholdRequest** | [**optional.Interface of VoiceprintThresholdRequest**](VoiceprintThresholdRequest.md)|  | 

### Return type

[**RespVoiceprintThresholdResponse**](RespVoiceprintThresholdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Vadcheck**
> RespVoiceprintVadcheckResponse Vadcheck(ctx, optional)


VAD检测

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VadcheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VadcheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintVadcheckRequest** | [**optional.Interface of VoiceprintVadcheckRequest**](VoiceprintVadcheckRequest.md)|  | 

### Return type

[**RespVoiceprintVadcheckResponse**](RespVoiceprintVadcheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verify**
> RespVoiceprintVerifyResponse Verify(ctx, optional)


声纹验证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintVerifyRequest** | [**optional.Interface of VoiceprintVerifyRequest**](VoiceprintVerifyRequest.md)|  | 

### Return type

[**RespVoiceprintVerifyResponse**](RespVoiceprintVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verify1ton**
> RespVoiceprint1tonVerifyResponse Verify1ton(ctx, optional)


声纹验证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Verify1tonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Verify1tonOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprint1tonVerifyRequest** | [**optional.Interface of Voiceprint1tonVerifyRequest**](Voiceprint1tonVerifyRequest.md)|  | 

### Return type

[**RespVoiceprint1tonVerifyResponse**](RespVoiceprint1tonVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMulti**
> RespVoiceprintVerifyMultiResponse VerifyMulti(ctx, optional)


声纹验证1对多

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintVerifyMultiRequest** | [**optional.Interface of VoiceprintVerifyMultiRequest**](VoiceprintVerifyMultiRequest.md)|  | 

### Return type

[**RespVoiceprintVerifyMultiResponse**](RespVoiceprintVerifyMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verifytopn**
> RespVoiceprinttopnVerifyResponse Verifytopn(ctx, optional)


声纹验证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifytopnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifytopnOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprinttopnVerifyRequest** | [**optional.Interface of VoiceprinttopnVerifyRequest**](VoiceprinttopnVerifyRequest.md)|  | 

### Return type

[**RespVoiceprinttopnVerifyResponse**](RespVoiceprinttopnVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

