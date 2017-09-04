package cloud_sdk

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type _RequestWarp struct {
	Id            string `bson:"id"`
	IdType        string `bson:"id_type"`
	CallTimeStamp int64  `bson:"t"`
	Data          []byte `bson:"data"`
	SkipCrypt     bool   `bson:"skip_crypt"`
	Sign          []byte `bson:"sign"`
	secret        string
}

func newRequestWarp(id, idType, secret string, skipCrypt bool) *_RequestWarp {
	return &_RequestWarp{
		Id:        id,
		IdType:    idType,
		secret:    secret,
		SkipCrypt: skipCrypt,
	}
}

func (this *_RequestWarp) setDataToBson(data interface{}) error {
	content, err := bson.Marshal(data)
	if nil != err {
		return err
	}
	this.Data = content
	return nil
}

func (this *_RequestWarp) toRequestBody() ([]byte, error) {
	callTimeStamp := time.Now().UnixNano() / 1e6
	signContent := fmt.Sprintf("%s%s%d%s", this.Id, this.IdType, callTimeStamp, string(this.Data))
	signResult, err := sign(this.secret, signContent)
	if nil != err {
		return nil, err
	}
	if false == this.SkipCrypt {
		var err error
		this.Data, err = encrypt(this.secret, this.Data)
		if nil != err {
			return nil, err
		}
	}
	this.Sign = signResult
	this.CallTimeStamp = callTimeStamp
	return bson.Marshal(this)
}
