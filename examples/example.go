/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/antihax/optional"
	sdk "github.com/speakin/sdk_go"
	"github.com/speakin/sdk_go/openapi"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name      string `yaml:"name"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
	} `yaml:"app"`
	Bucket struct {
		Name      string `yaml:"name"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
	} `yaml:"bucket"`
	SamplingRate string `yaml:"samplingRate"`
}

var (
	configFile = flag.String("c", "config.yaml", "config file")
)

func main() {
	flag.Parse()
	var config Config

	fmt.Println(*configFile)
	fileData, err := ioutil.ReadFile(*configFile)
	fmt.Printf("%s", string(fileData))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}

	// 修改为你的app key和bucket key
	client := sdk.NewClient(
		config.App.AccessKey,
		config.App.SecretKey,
		config.Bucket.AccessKey,
		config.Bucket.SecretKey,
	)
	// 修改为你的bucket名字和app名字
	bucket := config.Bucket.Name
	appName := config.App.Name
	// voice file
	userNames := []string{"user_a", "user_b"}
	userFiles := [][]string{
		{
			"../testdata/audio-data/u1/01_16k.wav",
			"../testdata/audio-data/u1/02_16k.wav",
			"../testdata/audio-data/u1/03_16k.wav",
		},
		{
			"../testdata/audio-data/u2/01_16k.wav",
			"../testdata/audio-data/u2/02_16k.wav",
			"../testdata/audio-data/u2/03_16k.wav",
		},
	}
	userFilesKey := make([][]string, 2)
	// 上传所有注册文件
	for i := range userNames {
		for _, filename := range userFiles[i] {
			userFilesKey[i] = append(userFilesKey[i], uploadFile(filename, client, bucket))
		}
	}
	// vad检测所有上传的文件
	for i := range userFilesKey {
		for _, key := range userFilesKey[i] {
			fmt.Printf("vadcheck key: %s\n", key)
			// vadcheck
			req := openapi.VoiceprintVadcheckRequest{
				AppName:      appName,
				Url:          key,
				SamplingRate: config.SamplingRate,
				Duration:     1.5,
				Timestamp:    time.Now().Unix(),
			}
			resp, _, err := client.VoiceprintApi.Vadcheck(context.Background(),
				&openapi.VadcheckOpts{VoiceprintVadcheckRequest: optional.NewInterface(req)})
			if err != nil {
				panic(err)
			}
			if resp.HasError {
				fmt.Printf("%s:%s\n", resp.ErrorId, resp.ErrorDesc)
				os.Exit(1)
			}
			if resp.Data.Code != "CHECKVAD_OK" {
				fmt.Printf("vadcheck: %s\n", resp.Data.Code)
			}
		}
	}
	// 注册
	for i, name := range userNames {
		req := openapi.VoiceprintRegisterRequest{
			AppName:      appName,
			UnionID:      name,
			Urls:         userFilesKey[i],
			SamplingRate: config.SamplingRate,
			Timestamp:    time.Now().Unix(),
			Replace:      true,
		}
		resp, _, err := client.VoiceprintApi.Register(context.Background(),
			&openapi.RegisterOpts{VoiceprintRegisterRequest: optional.NewInterface(req)})
		if err != nil {
			panic(err)
		}
		if resp.HasError {
			fmt.Printf("%s:%s\n", resp.ErrorId, resp.ErrorDesc)
			os.Exit(1)
		}
	}
	// 校验注册结果
	for _, name := range userNames {
		req := openapi.VoiceprintQueryRequest{
			AppName:   appName,
			UnionID:   name,
			Timestamp: time.Now().Unix(),
		}
		resp, _, err := client.VoiceprintApi.Query(context.Background(),
			&openapi.QueryOpts{VoiceprintQueryRequest: optional.NewInterface(req)})
		if err != nil {
			panic(err)
		}
		if resp.HasError {
			fmt.Printf("%s:%s\n", resp.ErrorId, resp.ErrorDesc)
			os.Exit(1)
		}
		fmt.Printf("user %s register: %v\n", name, resp.Data.IsRegister)
	}
	// 1比1比对
	checkFiles := []string{
		"../testdata/audio-data/u1/04_16k.wav",
		"../testdata/audio-data/u2/04_16k.wav",
	}
	// 上传upload文件
	checkFilesKeys := []string{
		uploadFile(checkFiles[0], client, bucket),
		uploadFile(checkFiles[1], client, bucket),
	}
	// 1比1验证文件
	for _, name := range userNames {
		for i, key := range checkFilesKeys {
			req := openapi.VoiceprintVerifyRequest{
				AppName:      appName,
				UnionID:      name,
				Url:          key,
				SamplingRate: config.SamplingRate,
				Timestamp:    time.Now().Unix(),
			}
			resp, _, err := client.VoiceprintApi.Verify(context.Background(),
				&openapi.VerifyOpts{VoiceprintVerifyRequest: optional.NewInterface(req)})
			if err != nil {
				panic(err)
			}
			if resp.HasError {
				fmt.Printf("%s:%s\n", resp.ErrorId, resp.ErrorDesc)
				os.Exit(1)
			}

			fmt.Printf("user %s verify file %s score: %v\n", name, checkFiles[i], resp.Data.Score)
		}
	}
	// 1比n验证
	for i, key := range checkFilesKeys {
		req := openapi.Voiceprint1tonVerifyRequest{
			AppName:      appName,
			UnionIDs:     userNames,
			Url:          key,
			SamplingRate: config.SamplingRate,
			Timestamp:    time.Now().Unix(),
		}
		resp, _, err := client.VoiceprintApi.Verify1ton(context.Background(),
			&openapi.Verify1tonOpts{Voiceprint1tonVerifyRequest: optional.NewInterface(req)})
		if err != nil {
			panic(err)
		}
		if resp.HasError {
			fmt.Printf("%s:%s\n", resp.ErrorId, resp.ErrorDesc)
			os.Exit(1)
		}
		fmt.Printf("verify file %s match user %s score: %v\n", checkFiles[i], resp.Data.UnionID, resp.Data.Score)
	}
	// 下载
	{
		_, resp, err := client.StorageApi.Download(context.Background(),
			bucket, checkFilesKeys[0])
		if err != nil {
			panic(err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile("../testdata/tmp/tmp.wav", b, 0644)
	}

}

func uploadFile(filename string, client *openapi.APIClient, bucket string) string {
	voice, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	resp, _, err := client.StorageApi.Upload(context.Background(),
		bucket, "wav", time.Now().Unix(),
		&openapi.UploadOpts{Body: optional.NewInterface(voice)})
	if err != nil {
		panic(err)
	}
	if resp.HasError {
		fmt.Printf("%s:%s", resp.ErrorId, resp.ErrorDesc)
		os.Exit(1)
	}
	return resp.Data.Key
}
