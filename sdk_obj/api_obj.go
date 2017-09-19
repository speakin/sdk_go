package sdk_obj

type StartSessionRequest struct {
	UserId      string `bson:"user_id"`
	CanRegister bool   `bson:"can_register"`
	CanVerify   bool   `bson:"can_verify"`
	CanIdentity bool   `bson:"can_identity"`
	Ttl         int32  `bson:"ttl"`
}

type StartSessionResponse struct {
	SessionId       string `bson:"session_id"`
	ExpireTimeStamp int64  `bson:"expire_time_stamp"`
	SessionSecret   string `bson:"session_secret"`
}

type ListModuleRequest struct {
}

type Module struct {
	VoiceBitCount int32  `bson:"voice_bit_count"`
	VoiceRate     int32  `bson:"voice_rate"`
	VoiceLang     string `bson:"voice_lang"`
}

type ListModuleRespone struct {
	ModuleList []Module `bson:"module_list"`
}

type QueryTraceRequest struct {
	Offset int32 `bson:"offset"`
	Limit  int32 `bson:"limit"`
}

type QueryTraceResponse struct {
	TotalCount int32  `bson:"total_count"`
	LogContent string `bson:"log_content"`
}
