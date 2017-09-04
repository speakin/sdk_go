package sdk_obj

type StartRecordRequest struct {
	GenText       bool   `bson:"gen_text"`
	VoiceBitCount int32  `bson:"voice_bit_count"`
	VoiceRate     int32  `bson:"voice_rate"`
	VoiceLang     string `bson:"voice_lang"`
	DataFormat    string `bson:"data_format"`
	TargetAction  string `bson:"target_action"`
}

type StartRecordResponse struct {
	RecordId string `bson:"record_id"`
	Text     string `bson:"text"`
}

type UploadRecordPartRequest struct {
	RecordId string `bson:"record_id"`
	IndexId  int32  `bson:"index_id"`
	Data     []byte `bson:"data"`
}

func (this *UploadRecordPartRequest) setRawData(data []byte) {
	//TODO
}

type UploadRecordPartResponse struct {
}

type UploadRecordDoneRequest struct {
	RecordId string `bson:"record_id"`
}

type UploadRecordDoneResponse struct {
}

type UploadRecordCancelRequest struct {
	RecordId string `bson:"record_id"`
}

type UploadRecordCancelResponse struct {
}

type RegisterRequest struct {
	RecordIdList []string `bson:"record_id_list"`
}

type RegisterResponse struct {
}

type VerifyRequest struct {
	RecordId string `bson:"record_id"`
}

type VerifyResponse struct {
	Result         bool    `bson:"result"`
	Score          float64 `bson:"score"`
	ThresholdScore float64 `bson:"threshold_score"`
}

type IdentityRequest struct {
	RecordId string `bson:"record_id"`
}

type IdentityResponse struct {
	UserIdList []string `bson:"user_id_list"`
}

type SimilarRequest struct {
	RecordId string `bson:"record_id"`
	TopCount int32  `bson:"top_count"`
}

type SimilarItem struct {
	UserId       string  `bson:"user_id"`
	SimilarValue float64 `bson:"similar_value"`
}

type SimilarResponse struct {
	SimilarList []*SimilarItem `bson:"similar_list"`
}
