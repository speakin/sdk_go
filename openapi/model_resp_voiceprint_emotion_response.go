/*
 * 声纹云api
 *
 * api document
 *
 * API version: /cloudapi/v1beta
 * Contact: xiachengjia@speakin.mobi
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RespVoiceprintEmotionResponse struct {
	HasError bool `json:"hasError,omitempty"`
	ErrorId string `json:"errorId,omitempty"`
	ErrorDesc string `json:"errorDesc,omitempty"`
	Data VoiceprintEmotionResponse `json:"data,omitempty"`
}
