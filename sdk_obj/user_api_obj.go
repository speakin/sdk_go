package sdk_obj

type SetAppUserRequest struct {
	UserId           string `bson:"user_id"`
	UserType         string `bson:"user_type"`
	Valid            bool   `bson:"valid"`
	AccessAllAppUser bool   `bson:"access_all_app_user"`
}

type SetAppUserResponse struct {
}

type GetAppUserRequest struct {
	UserId string `bson:"user_id"`
}

type GetAppUserResponse struct {
	UserId           string `bson:"user_id"`
	UserType         string `bson:"user_type"`
	Valid            bool   `bson:"valid"`
	AccessAllAppUser bool   `bson:"access_all_app_user"`
}

type AddChildAppUserRequest struct {
	UserId      string `bson:"user_id"`
	ChildUserId string `bson:"child_user_id"`
}

type AddChildAppUserResponse struct {
}

type RemoveChildAppUserRequest struct {
	UserId      string `bson:"user_id"`
	ChildUserId string `bson:"child_user_id"`
}

type RemoveChildAppUserResponse struct {
}

type GetChildAppUserCountRequest struct {
	UserId string `bson:"user_id"`
}

type GetChildAppUserCountResponse struct {
	Count int32 `bson:"count"`
}

type ListChildAppUserRequest struct {
	UserId string `bson:"user_id"`
	Offset int32  `bson:"offset"`
	Limit  int32  `bson:"limit"`
}

type ListChildAppUserResponse struct {
	ChildUserIdList []string `bson:"child_user_id_list"`
}

type ContainChildAppUserRequest struct {
	UserId      string `bson:"user_id"`
	ChildUserId string `bson:"child_user_id"`
}

type ContainChildAppUserResponse struct {
	Contain bool `bson:"contain"`
}
