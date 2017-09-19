package sdk_go

import (
	"fmt"
	"github.com/speakin/sdk_go/sdk_obj"
	"net/url"
)

type SessionApi struct {
	sessionId     string
	sessionSecret string
	baseUrl       string
	skipCrypt     bool
}

func NewSessionApi(sessionId, sessionSecret, baseUrl string, skipCrypt bool) *SessionApi {
	return &SessionApi{
		sessionId:     sessionId,
		sessionSecret: sessionSecret,
		baseUrl:       baseUrl,
		skipCrypt:     skipCrypt,
	}
}

func (this *SessionApi) StartRecord(request *sdk_obj.StartRecordRequest, traceId ...string) (*sdk_obj.StartRecordResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/record/start?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.StartRecordResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) OpenRecordStream(recordId string, traceId ...string) *RecordStream {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	return newRecordStream(this, recordId, tid)
}

func (this *SessionApi) uploadRecordPart(request *sdk_obj.UploadRecordPartRequest, traceId ...string) (*sdk_obj.UploadRecordPartResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/record/upload_part?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.UploadRecordPartResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) uploadRecordDone(request *sdk_obj.UploadRecordDoneRequest, traceId ...string) (*sdk_obj.UploadRecordDoneResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/record/done?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.UploadRecordDoneResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) uploadRecordCancel(request *sdk_obj.UploadRecordCancelRequest, traceId ...string) (*sdk_obj.UploadRecordCancelResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/record/cancel?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.UploadRecordCancelResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) Register(request *sdk_obj.RegisterRequest, traceId ...string) (*sdk_obj.RegisterResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/register?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.RegisterResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) Verify(request *sdk_obj.VerifyRequest, traceId ...string) (*sdk_obj.VerifyResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/verify?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.VerifyResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) Identity(request *sdk_obj.IdentityRequest, traceId ...string) (*sdk_obj.IdentityResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/identity?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.IdentityResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}

func (this *SessionApi) Similar(request *sdk_obj.SimilarRequest, traceId ...string) (*sdk_obj.SimilarResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/session/similar?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.SimilarResponse{}
	apiError := callApi(request, response, httpUrl, this.sessionId, sdk_obj.ID_TYPE_SESSION, this.sessionSecret, this.skipCrypt)
	return response, apiError
}
