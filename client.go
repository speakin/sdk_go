/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package voiceprint_cloud_go_git

import (
	"github.com/speakin/sdk_go/openapi"
	"net/http"
)

func NewClient(accessKey, secretKey, bucketAccessKey, bucketSecretKey string) *openapi.APIClient {
	cfg := openapi.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &TokenTransport{
			AccessKey:       accessKey,
			SecretKey:       secretKey,
			BucketAccessKey: bucketAccessKey,
			BucketSecretkey: bucketSecretKey,
			Transport:       http.DefaultTransport,
		},
	}
	return openapi.NewAPIClient(cfg)
}

func NewClientWithHost(host, accessKey, secretKey, bucketAccessKey, bucketSecretKey string) *openapi.APIClient {
	cfg := openapi.NewConfiguration()
	cfg.BasePath = host
	cfg.HTTPClient = &http.Client{
		Transport: &TokenTransport{
			AccessKey:       accessKey,
			SecretKey:       secretKey,
			BucketAccessKey: bucketAccessKey,
			BucketSecretkey: bucketSecretKey,
			Transport:       http.DefaultTransport,
		},
	}
	return openapi.NewAPIClient(cfg)
}
