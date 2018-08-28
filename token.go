/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package voiceprint_cloud_go_git

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenTokenAuth(secretKey string, data []byte) string {
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(data))
	digest := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString([]byte(digest))
}

func (tr *TokenTransport) GenerateToken(body []byte) string {
	sign := GenTokenAuth(tr.SecretKey, body)
	token := tr.AccessKey + ":" + sign
	return token
}

func (tr *TokenTransport) GenerateStorageToken(data []byte) string {
	sign := GenTokenAuth(tr.BucketSecretkey, data)
	token := tr.BucketAccessKey + ":" + sign
	return token
}

// use to generate Authorization token header
type TokenTransport struct {
	AccessKey       string
	SecretKey       string
	BucketAccessKey string
	BucketSecretkey string
	Transport       http.RoundTripper
}

func (tr *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	pathPrefix := req.URL.EscapedPath()
	// only generate token with json data
	if strings.HasPrefix(pathPrefix, "/cloudapi/v1beta/voiceprint/") {
		reader, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		buf, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		req.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewBuffer(buf)), nil
		}
		req.Header.Set("Authorization", "SpeakIn "+tr.GenerateToken(buf))
	} else if strings.HasPrefix(pathPrefix, "/cloudapi/v1beta/storage/") {
		req.Header.Set("Authorization", "SpeakIn "+
			tr.GenerateStorageToken([]byte(fmt.Sprintf("%s?%s", req.URL.Path, req.URL.RawQuery))))
	}

	return tr.Transport.RoundTrip(req)
}
