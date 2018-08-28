/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package voiceprint_cloud_go_git

import (
	"context"
	"github.com/speakin/sdk_go/openapi"
	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
	"testing"
	"time"
)

func uploadFileUnit(t *testing.T, client *openapi.APIClient, fileName, bucket string) string {
	voice, err := os.Open(fileName)
	assert.NoError(t, err)

	resp, _, err := client.StorageApi.Upload(
		context.Background(),
		bucket, "wav", 0, time.Now().Unix(),
		&openapi.UploadOpts{Body: optional.NewInterface(voice)})
	assert.NoError(t, err)
	if resp.HasError {
		t.Log(resp.ErrorDesc)
		t.FailNow()
	}
	return resp.Data.Key
}

func TestFlow(t *testing.T) {
	var config Config
	fileData, err := ioutil.ReadFile("config_test.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}

	client := NewClient(
		config.App.AccessKey,
		config.App.SecretKey,
		config.Bucket.AccessKey,
		config.Bucket.SecretKey,
	)

	// upload voice file
	userKeys := make([][]string, len(config.TestData.Users))

	for i, user := range config.TestData.Users {
		for _, filename := range user.Files {
			userKeys[i] = append(userKeys[i], uploadFileUnit(t, client, filename, config.Bucket.Name))
		}
	}
	// register
	{
		for i, user := range config.TestData.Users {
			req := openapi.VoiceprintRegisterRequest{
				AppName:      config.App.Name,
				UnionID:      user.Name,
				Urls:         userKeys[i][:3],
				SamplingRate: config.TestData.SimpleRate,
				Timestamp:    time.Now().Unix(),
				Replace:      true,
			}
			resp, _, err := client.VoiceprintApi.Register(context.Background(),
				&openapi.RegisterOpts{VoiceprintRegisterRequest: optional.NewInterface(req)})
			assert.NoError(t, err)
			if resp.HasError {
				t.Log(resp.ErrorDesc)
				t.FailNow()
			}
		}
	}
	// query
	{
		for _, user := range config.TestData.Users {
			req := openapi.VoiceprintQueryRequest{
				AppName:   config.App.Name,
				UnionID:   user.Name,
				Timestamp: time.Now().Unix(),
			}
			resp, _, err := client.VoiceprintApi.Query(context.Background(),
				&openapi.QueryOpts{VoiceprintQueryRequest: optional.NewInterface(req)})
			assert.NoError(t, err)
			if resp.HasError {
				t.Log(resp.ErrorDesc)
				t.FailNow()
			}
			t.Logf("user %s register: %v", user.Name, resp.Data.IsRegister)
			assert.True(t, resp.Data.IsRegister)
		}
	}

	// verify
	for i, user := range config.TestData.Users {
		ownScopes := make([]float64, 0, len(user.Files[3:]))
		otherUsersBefore := config.TestData.Users[:i]
		otherUsersAfter := config.TestData.Users[i+1:]
		otherScopes := []float64{}
		for j, key := range userKeys[i][3:] {
			req := openapi.VoiceprintVerifyRequest{
				AppName:      config.App.Name,
				UnionID:      user.Name,
				Url:          key,
				SamplingRate: "16k",
				Timestamp:    time.Now().Unix(),
			}
			resp, _, err := client.VoiceprintApi.Verify(context.Background(),
				&openapi.VerifyOpts{VoiceprintVerifyRequest: optional.NewInterface(req)})
			assert.NoError(t, err)
			if resp.HasError {
				t.Log(resp.ErrorDesc)
				t.FailNow()
			}
			t.Logf("user %s compare file name %s score: %v", user.Name, user.Files[3+j], resp.Data.Score)
			ownScopes = append(ownScopes, resp.Data.Score)
		}

		for j, otherUser := range otherUsersBefore {
			for k, key := range userKeys[j][3:] {
				req := openapi.VoiceprintVerifyRequest{
					AppName:      config.App.Name,
					UnionID:      user.Name,
					Url:          key,
					SamplingRate: "16k",
					Timestamp:    time.Now().Unix(),
				}
				resp, _, err := client.VoiceprintApi.Verify(context.Background(),
					&openapi.VerifyOpts{VoiceprintVerifyRequest: optional.NewInterface(req)})
				assert.NoError(t, err)
				if resp.HasError {
					t.Log(resp.ErrorDesc)
					t.FailNow()
				}
				t.Logf("user %s compare file name %s score: %v", user.Name, otherUser.Files[3+k], resp.Data.Score)
				otherScopes = append(otherScopes, resp.Data.Score)
			}
		}

		for j, otherUser := range otherUsersAfter {
			for k, key := range userKeys[j+i+1][3:] {
				req := openapi.VoiceprintVerifyRequest{
					AppName:      config.App.Name,
					UnionID:      user.Name,
					Url:          key,
					SamplingRate: "16k",
					Timestamp:    time.Now().Unix(),
				}
				resp, _, err := client.VoiceprintApi.Verify(context.Background(),
					&openapi.VerifyOpts{VoiceprintVerifyRequest: optional.NewInterface(req)})
				assert.NoError(t, err)
				if resp.HasError {
					t.Log(resp.ErrorDesc)
					t.FailNow()
				}
				t.Logf("user %s compare file name %s score: %v", user.Name, otherUser.Files[3+k], resp.Data.Score)
				otherScopes = append(otherScopes, resp.Data.Score)
			}
		}
		sort.Float64s(ownScopes)
		sort.Float64s(otherScopes)

		assert.True(t, ownScopes[0] > otherScopes[len(otherScopes)-1])
	}

	// multiple test
	var names []string
	for _, user := range config.TestData.Users {
		names = append(names, user.Name)
	}
	for i, user := range config.TestData.Users {
		for k, key := range userKeys[i][3:] {
			req := openapi.Voiceprint1tonVerifyRequest{
				AppName:      config.App.Name,
				UnionIDs:     names,
				Url:          key,
				SamplingRate: "16k",
				Timestamp:    time.Now().Unix(),
			}
			resp, _, err := client.VoiceprintApi.Verify1ton(context.Background(),
				&openapi.Verify1tonOpts{Voiceprint1tonVerifyRequest: optional.NewInterface(req)})
			assert.NoError(t, err)
			if resp.HasError {
				t.Log(resp.ErrorDesc)
				t.FailNow()
			}
			t.Logf("user %s multiple compare file name %s target: %s score: %v", user.Name, user.Files[3+k], resp.Data.UnionID, resp.Data.Score)
			assert.Equal(t, user.Name, resp.Data.UnionID)
		}
	}
}
