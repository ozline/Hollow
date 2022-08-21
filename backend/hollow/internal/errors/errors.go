package errors

import (
	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	//SERVICE
	ErrMissingParams = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "Missing Params Data")
	ErrParamsIllegal = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "params invalid")

	//BIZ
	ErrUserCheckFailed = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Invalid")
	ErrUserNotFound    = errors.New(422, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "User Not Found")
	ErrUserExisted     = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "User Existed")

	//DATA
	ErrCode   = 422
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "Unknown Error")

	ErrUserNotExisted  = errors.New(ErrCode, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Not Existed")
	ErrFatherNotExistd = errors.New(ErrCode, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "Father Not Existed")

	ErrUserInvalid = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Invalid")
)
