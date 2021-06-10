package response

type UserInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User  UserInfo `json:"user"`
	Error Error    `json:"error"`
}

type GetUsersResponse struct {
	Users []UserInfo `json:"users"`
	Error Error      `json:"error"`
}

type UpdateUserResponse struct {
	User  UserInfo `json:"user"`
	Error Error    `json:"error"`
}

type DeleteUserResponse struct {
	Error Error `json:"error"`
}
