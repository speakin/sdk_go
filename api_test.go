package sdk_go

import (
	"fmt"
	"github.com/speakin/sdk_go/sdk_obj"
	"io/ioutil"
	"testing"
)

const (
	BASE_URL   = "http://api3.speakin.mobi"
	APP_ID     = "speakin_test"
	APP_SECRET = "mr1imt1etu7ryets9eoergua87h0pa4n"
)

var api *Api

func TestGlobalApi(t *testing.T) {
	api := NewApi(APP_ID, APP_SECRET, BASE_URL)
	resp, apiErr := api.ListModule(&sdk_obj.ListModuleRequest{})
	fmt.Println(resp, apiErr)
}

func TestSessionApi(t *testing.T) {
	userId := "abc"
	api := NewApi(APP_ID, APP_SECRET, BASE_URL)
	_, apiErr := api.GetUserApi().SetAppUser(&sdk_obj.SetAppUserRequest{
		UserId:   userId,
		UserType: sdk_obj.USER_TYPE_PEOPLE,
		Valid:    true,
	})
	if nil != apiErr {
		t.Error(apiErr)
	}
	sessionInfo, apiErr := api.StartSession(&sdk_obj.StartSessionRequest{
		CanIdentity: true,
		CanRegister: true,
		CanVerify:   true,
		Ttl:         7200,
		UserId:      userId,
	})
	if nil != apiErr {
		t.Error(apiErr)
	}
	sessApi := NewSessionApi(sessionInfo.SessionId, sessionInfo.SessionSecret, BASE_URL, true)
	recordId, apiErr := upload(sessApi, sdk_obj.TARGET_ACTION_REGISTER)
	if nil != apiErr {
		t.Error(apiErr)
	}
	_, apiErr = sessApi.Register(&sdk_obj.RegisterRequest{
		RecordIdList: []string{recordId},
	})
	if nil != apiErr {
		t.Error(apiErr)
	}
	recordId, apiErr = upload(sessApi, sdk_obj.TARGET_ACTION_VERIFY)
	if nil != apiErr {
		t.Error(apiErr)
	}
	verifyResp, apiErr := sessApi.Verify(&sdk_obj.VerifyRequest{
		RecordId: recordId,
	})
	fmt.Println(verifyResp.Score, verifyResp.ThresholdScore)
}

func upload(sessApi *SessionApi, target string) (string, *ApiError) {
	rs, apiErr := sessApi.OpenRecordStream(&sdk_obj.StartRecordRequest{
		DataFormat:    sdk_obj.DATA_FORMAT_WAV,
		TargetAction:  target,
		VoiceBitCount: 16,
		VoiceRate:     16000,
		VoiceLang:     "common-short",
	})
	if nil != apiErr {
		return "", apiErr
	}
	content, err := ioutil.ReadFile("test_data/SI_device1_1262_Samsung1_advertisement_2-concated.wav")
	if nil != err {
		return "", apiErr
	}
	apiErr = rs.Write(content)
	if nil != apiErr {
		return "", apiErr
	}
	apiErr = rs.Done()
	if nil != apiErr {
		return "", apiErr
	}
	return rs.GetRecordId(), nil
}
