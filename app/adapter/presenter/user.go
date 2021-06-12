package presenter

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
	"github.com/yudwig/echo-sample/app/core/presentation/response"
)

type UserPresenter struct {
}

func NewUserPresenter() UserPresenter {
	return UserPresenter{}
}

func createUserInfo(user user.User) response.UserInfo {
	return response.UserInfo{
		Id:   user.Id,
		Name: user.Name,
	}
}

func (p UserPresenter) MakeCreateUserResponse(user user.User, err error) *response.CreateUserResponse {
	return &response.CreateUserResponse{
		User:  createUserInfo(user),
		Error: response.NewError(err),
	}
}

func (p UserPresenter) MakeGetUsersResponse(users []user.User, err error) *response.GetUsersResponse {
	var userInfos []response.UserInfo
	for _, u := range users {
		userInfos = append(userInfos, createUserInfo(u))
	}
	return &response.GetUsersResponse{
		Users: userInfos,
		Error: response.NewError(err),
	}
}

func (p UserPresenter) MakeUpdateUserResponse(user user.User, err error) *response.UpdateUserResponse {
	return &response.UpdateUserResponse{
		User:  createUserInfo(user),
		Error: response.NewError(err),
	}
}

func (p UserPresenter) MakeDeleteUserResponse(err error) *response.DeleteUserResponse {
	return &response.DeleteUserResponse{
		Error: response.NewError(err),
	}
}
