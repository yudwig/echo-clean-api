package response

import "github.com/yudwig/echo-sample/app/core/domain/entity/errors"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(message string) Error {
	return Error{
		Code:    getErrorCode(message),
		Message: message,
	}
}

var errorCodes = map[string]int{
	errors.EmptyUserNameError: 1000,
}

func getErrorCode(message string) int {
	return errorCodes[message]
}
