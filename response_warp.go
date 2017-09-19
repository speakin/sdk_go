package sdk_go

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type _ResponseWarp struct {
	HasError        bool     `bson:"has_error"`
	Error           ApiError `bson:"error"`
	Data            []byte   `bson:"data"`
	ExcuteTimeStamp int64    `bson:"t"`
	Sign            []byte   `bson:"sign"`
}

func newResponseWarp(secret string, content []byte) (*_ResponseWarp, error) {
	ret := &_ResponseWarp{}
	err := bson.Unmarshal(content, ret)
	if nil != err {
		return nil, err
	}
	if ret.HasError && strings.HasPrefix(ret.Error.ErrorId, "common.") {
		return ret, nil
	}
	ret.Data, err = decrypt(secret, ret.Data)
	if nil != err {
		return nil, err
	}
	err = ret.checkSign(secret)
	if nil != err {
		return nil, err
	}
	return ret, nil
}

func (this *_ResponseWarp) checkSign(secret string) error {
	if this.HasError && strings.HasPrefix(this.Error.ErrorId, "common.") {
		return nil
	}
	hasErrorStr := "true"
	if false == this.HasError {
		hasErrorStr = "false"
	}
	signContent := fmt.Sprintf("%d%s%s%s%s", this.ExcuteTimeStamp, hasErrorStr, this.Error.ErrorId, this.Error.ErrorDesc, string(this.Data))
	signResult, err := sign(secret, signContent)
	if nil != err {
		return err
	}
	if len(signResult) != len(this.Sign) {
		this.Error = ApiError{
			ErrorId:   COMMON_ERROR_WRONG_SIGN,
			ErrorDesc: "wrong sign",
		}
	} else {
		l := len(signResult)
		for i := 0; i < l; i++ {
			if signResult[i] != this.Sign[i] {
				this.Error = ApiError{
					ErrorId:   COMMON_ERROR_WRONG_SIGN,
					ErrorDesc: "wrong sign",
				}
			}
		}
	}
	return nil
}
