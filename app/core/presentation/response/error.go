package response

import "github.com/yudwig/echo-sample/app/core/domain/entity/errors"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(err error) Error {
	if err == nil {
		return Error{
			Code:    0,
			Message: "",
		}
	}
	return Error{
		Code:    getErrorCode(err.Error()),
		Message: err.Error(),
	}
}

var errorCodes = map[string]int{
	errors.EmptyUserNameError: 1000,
}

func getErrorCode(msg string) int {
	code, err := errorCodes[msg]
	if err == false {
		return -1
	}
	return code
}
