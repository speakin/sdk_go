package sdk_go

import (
	"fmt"
	"github.com/speakin/sdk_go/sdk_obj"
	"net/url"
)

type Api struct {
	appId     string
	appSecret string
	baseUrl   string
}

func NewApi(appId, appSecret, baseUrl string) *Api {
	return &Api{
		appId:     appId,
		appSecret: appSecret,
		baseUrl:   baseUrl,
	}
}

func (this *Api) ListModule(request *sdk_obj.ListModuleRequest, traceId ...string) (*sdk_obj.ListModuleRespone, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/list_module?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.ListModuleRespone{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *Api) StartSession(request *sdk_obj.StartSessionRequest, traceId ...string) (*sdk_obj.StartSessionResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/start_session?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.StartSessionResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *Api) GetUserApi() *UserApi {
	return newUserApi(this.appId, this.appSecret, this.baseUrl)
}
