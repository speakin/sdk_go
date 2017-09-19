package sdk_go

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var httpClient *http.Client

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func encrypt(secret string, data []byte) ([]byte, error) {
	b, err := aes.NewCipher([]byte(secret))
	if nil != err {
		return nil, err
	}
	content := pkcs5Padding(data, b.BlockSize())
	result := make([]byte, len(content)+aes.BlockSize)
	iv := result[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)
	mode := cipher.NewCBCEncrypter(b, iv)
	mode.CryptBlocks(result[aes.BlockSize:], content)
	return result, nil
}

func decrypt(secret string, data []byte) ([]byte, error) {
	b, err := aes.NewCipher([]byte(secret))
	if nil != err {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(b, data[:aes.BlockSize])
	result := make([]byte, len(data)-aes.BlockSize)
	mode.CryptBlocks(result, data[aes.BlockSize:])
	result = pkcs5UnPadding(result)
	return result, nil
}

func sign(secret, data string) ([]byte, error) {
	h := sha1.New()
	h.Write([]byte(data))
	h.Write([]byte(secret))
	return h.Sum(nil), nil
}

func callApi(request, response interface{}, httpUrl string, id, idType, secret string, skipCrypt bool) *ApiError {
	reqWarp := newRequestWarp(id, idType, secret, skipCrypt)
	err := reqWarp.setDataToBson(request)
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	reqBody, err := reqWarp.toRequestBody()
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	res, err := httpClient.Post(httpUrl, "application/bson", bytes.NewBuffer(reqBody))
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	defer res.Body.Close()
	respData, err := ioutil.ReadAll(res.Body)
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	respWarp, err := newResponseWarp(reqWarp.secret, respData)
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	if respWarp.HasError {
		return &respWarp.Error
	}
	err = bson.Unmarshal([]byte(respWarp.Data), response)
	if nil != err {
		return &ApiError{
			ErrorId:   COMMON_ERROR_UNKWON,
			ErrorDesc: err.Error(),
		}
	}
	return nil
}

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			DisableCompression:    true,
			MaxIdleConnsPerHost:   1024,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 10 * time.Second,
	}
}
