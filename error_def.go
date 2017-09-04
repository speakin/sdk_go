package cloud_sdk

type ApiError struct {
	ErrorId   string `bson:"id"`
	ErrorDesc string `bson:"desc"`
}

const (
	COMMON_ERROR_MISS_PARAM        = "common.miss_param"
	COMMON_ERROR_WRONG_TIME_STAMP  = "common.wrong_time_stamp"
	COMMON_ERROR_WRONG_SIGN        = "common.wrong_sign"
	COMMON_ERROR_WRONG_DATA        = "common.wrong_data"
	COMMON_ERROR_UNKWON            = "common.unkwon"
	COMMON_ERROR_UNKWON_APP_ID     = "common.unkwon_app_id"
	COMMON_ERROR_UNKWON_SESSION_ID = "common.unkwon_session_id"
	COMMON_ERROR_INVALID_ID_TYPE   = "common.invalid_id_type"

	USER_WRONG_TYPE = "user.wrong_type"
	USER_NOT_EXIST  = "user.not_exist"
	USER_NO_PARENT  = "user.no_parent"
	USER_NO_CHILD   = "user.no_child"
	USER_NOT_VALID  = "user.not_valid"

	RECORD_PRE_NOT_DONE          = "record.pre_not_done"
	RECORD_WRONG_ID              = "record.wrong_id"
	RECORD_WRONG_TARGET          = "record.wrong_target"
	RECORD_TARGET_NOT_ALLOW      = "record.target_not_allow"
	RECORD_NO_MODULE             = "record.no_module"
	RECORD_NOT_START             = "record.not_start"
	RECORD_UNSUPPORT_DATA_FORMAT = "record.unsupport_data_format"
	RECORD_SNR_TOO_LOW           = "record.snr_too_low"
	RECORD_SPEECH_TOO_SHORT      = "record.speech_too_short"
	RECORD_VOLUMN_TOO_LOW        = "record.volumn_too_low"
	RECORD_WRONG_DATA            = "record.wrong_data"
	RECORD_WRONG_VOICE_BIT_COUNT = "record.wrong_voice_bit_count"
	RECORD_WRONG_VOICE_RATE      = "record.wrong_voice_rate"

	REGISTER_WRONG_RECORD_ID  = "register.wrong_record_id"
	REGISTER_NEED_MORE_RECORD = "register.need_more_record"
	REGISTER_MULTI_MODULE     = "register.multi_module"
	REGISTER_NO_MODULE_FOUND  = "register.no_module_found"

	VERIFY_WRONG_RECORD_ID = "verify.wrong_record_id"
	VERIFY_NO_MODULE_FOUND = "verify.no_module_found"
	VERIFY_NEED_REGISTER   = "verify.need_register"

	IDENTITY_WRONG_RECORD_ID = "identity.wrong_record_id"
	IDENTITY_NO_MODULE_FOUND = "identity.no_module_found"
)
