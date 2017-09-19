package sdk_go

import (
	"github.com/speakin/sdk_go/sdk_obj"
)

type RecordStream struct {
	api      *SessionApi
	recordId string
	indexId  int32
	traceId  string
}

func newRecordStream(api *SessionApi, recordId, traceId string) *RecordStream {
	return &RecordStream{
		api:      api,
		recordId: recordId,
		traceId:  traceId,
	}
}

func (this *RecordStream) Write(data []byte) *ApiError {
	this.indexId += 1
	_, apiError := this.api.uploadRecordPart(&sdk_obj.UploadRecordPartRequest{
		IndexId:  this.indexId,
		RecordId: this.recordId,
		Data:     data,
	}, this.traceId)
	return apiError
}

func (this *RecordStream) Done() *ApiError {
	_, apiError := this.api.uploadRecordDone(&sdk_obj.UploadRecordDoneRequest{
		RecordId: this.recordId,
	}, this.traceId)
	return apiError
}

func (this *RecordStream) Cancel() *ApiError {
	_, apiError := this.api.uploadRecordCancel(&sdk_obj.UploadRecordCancelRequest{
		RecordId: this.recordId,
	}, this.traceId)
	return apiError
}

func (this *RecordStream) GetRecordId() string {
	return this.recordId
}
