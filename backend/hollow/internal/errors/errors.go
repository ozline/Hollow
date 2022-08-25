package errors

import (
	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type Error struct {
	Code    int
	Reason  string
	Message string
}

var (
	ErrCodeNormal              = 400
	ErrCodeInternalServerError = 500
	ErrCodeNeedMFA             = 417

	//SERVICE
	ErrMissingParams = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "Missing Params Data")
	ErrParamsIllegal = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "params invalid")

	//BIZ
	ErrUserCheckFailed = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Invalid")
	ErrUserNotFound    = errors.New(ErrCodeNormal, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "User Not Found")
	ErrUserExisted     = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "User Existed")

	//DATA
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "Unknown Error")

	ErrUserNotExisted   = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Not Existed")
	ErrFatherNotExisted = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "Father Not Existed")
	ErrLeafNotExisted   = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "Leaf Not Existed")

	ErrUserInvalid    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Invalid")
	ErrUserNotMatched = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Not Matched")

	ErrCommentNotFound = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "Comment Not Found")
	ErrHaveLiked       = errors.New(ErrCodeNormal, v1.ErrorReason_REPEATED_OPERATION.String(), "You have liked this comment")

	ErrNeedMFA          = errors.New(ErrCodeNeedMFA, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "Need MFA")
	ErrHaveActivatedMFA = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "You have Activated MFA")
	ErrNotEnabledMFA    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "You have not enabled MFA")
	ErrMFAVerifyFailed  = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "MFA Verify Failed")
)

func GenerateError(err Error) *errors.Error {
	return errors.New(err.Code, err.Reason, err.Message)
}

// 产生默认错误信息
func GenerateErrorNormal(err error) *errors.Error {
	return errors.New(ErrCodeInternalServerError, v1.ErrorReason_NORMAL_ERROR.String(), err.Error())
}
