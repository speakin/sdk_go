package sdk_go

import (
	"fmt"
	"github.com/speakin/sdk_go/sdk_obj"
	"net/url"
)

type UserApi struct {
	appId     string
	appSecret string
	baseUrl   string
}

func newUserApi(appId, appSecret, baseUrl string) *UserApi {
	return &UserApi{
		appId:     appId,
		appSecret: appSecret,
		baseUrl:   baseUrl,
	}
}

func (this *UserApi) SetAppUser(request *sdk_obj.SetAppUserRequest, traceId ...string) (*sdk_obj.SetAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/set?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.SetAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) GetAppUser(request *sdk_obj.GetAppUserRequest, traceId ...string) (*sdk_obj.GetAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/get?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.GetAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) AddChildAppUser(request *sdk_obj.AddChildAppUserRequest, traceId ...string) (*sdk_obj.AddChildAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/add_child?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.AddChildAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) RemoveChildAppUser(request *sdk_obj.RemoveChildAppUserRequest, traceId ...string) (*sdk_obj.RemoveChildAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/remove_child?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.RemoveChildAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) GetChildAppUserCount(request *sdk_obj.GetChildAppUserCountRequest, traceId ...string) (*sdk_obj.GetChildAppUserCountResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/get_child_count?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.GetChildAppUserCountResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) ListChildAppUser(request *sdk_obj.ListChildAppUserRequest, traceId ...string) (*sdk_obj.ListChildAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/list_child?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.ListChildAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}

func (this *UserApi) ContainChildAppUser(request *sdk_obj.ContainChildAppUserRequest, traceId ...string) (*sdk_obj.ContainChildAppUserResponse, *ApiError) {
	tid := ""
	if len(traceId) > 0 {
		tid = traceId[0]
	}
	httpUrl := fmt.Sprintf("%s/v1/user/contain_child?traceId=%s", this.baseUrl, url.QueryEscape(tid))
	response := &sdk_obj.ContainChildAppUserResponse{}
	apiError := callApi(request, response, httpUrl, this.appId, sdk_obj.ID_TYPE_APP, this.appSecret, false)
	return response, apiError
}
