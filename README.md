

### 安装

    go get -u github.com/speakin/sdk_go

### 使用

加载sdk包

```go
import sdk "github.com/speakin/sdk_go"

```

加载openapi

```go
import "github.com/speakin/sdk_go/openapi"
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

client := sdk.NewClientWithHost(
        "host",
		"your_access_key", // app access key
		"your_secret_key", // app secret key
		"your_bucket_access_key", // bucket access key
		"your_bucket_secret_key", // bucket secret key
	)
```

### 完整例子
``` shell
examples/example.go
```

