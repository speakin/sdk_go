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

type VoiceprintSplitResponse struct {
	// 开始时间
	StartMS int32 `json:"startMS,omitempty"`
	// 结束时间
	EndMS int32 `json:"endMS,omitempty"`
}
