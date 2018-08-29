

### 安装

    go get -u github.com/speakin/go_sdk

### 使用

加载sdk包

```go
import sdk "github.com/speakin/go_sdk"

```

加载openapi

```go
import "github.com/speakin/go_sdk/openapi"
```

初始化客户端

[embedmd]:# (examples/example.go /client := sdk/ /\)/)
```go
client := sdk.NewClient(
		"your_access_key", // app access key
		"your_secret_key", // app secret key
		"your_bucket_access_key", // bucket access key
		"your_bucket_secret_key", // bucket secret key
	)
```

### 完整例子

[embedmd]:# (examples/example.go /package/ $)
```go
package main

import (
	"context"
	"fmt"
	sdk "github.com/speakin/sdk_go"
	"github.com/speakin/sdk_go/openapi"
	"github.com/antihax/optional"
	"os"
	"time"
	"io/ioutil"
)

func uploadFile(filename string, client *openapi.APIClient, bucket string) string {
	voice, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	resp, _, err := client.StorageApi.Upload(context.Background(),
		bucket, "wav", 0, time.Now().Unix(),
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

func main() {
	// 修改为你的app key和bucket key
	client := sdk.NewClient(
		"your_access_key", // app access key
		"your_secret_key", // app secret key
		"your_bucket_access_key", // bucket access key
		"your_bucket_secret_key", // bucket secret key
	)
	// 修改为你的bucket名字和app名字
	bucket := "bucket-test"
	appName := "app-test"
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
	// 注册
	for i, name := range userNames {
		req := openapi.VoiceprintRegisterRequest{
			AppName:      appName,
			UnionID:      name,
			Urls:         userFilesKey[i],
			SamplingRate: "16k",
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
				SamplingRate: "16k",
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
			SamplingRate: "16k",
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
```
