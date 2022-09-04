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
	ErrMissingParams = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "缺失必要参数")
	ErrParamsIllegal = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "参数非法")

	//BIZ
	ErrUserCheckFailed     = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "账号验证失败")
	ErrUserNotFound        = errors.New(ErrCodeNormal, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "用户没有找到")
	ErrUserExisted         = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "用户名已经存在")
	ErrMsgCodeVerifyFailed = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "验证码验证失败")
	ErrMsgCodeInvalid      = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "验证码已过期或不存在")

	//DATA
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "未知错误")

	ErrUserNotExisted   = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不存在")
	ErrFatherNotExisted = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "父楼不存在")
	ErrLeafNotExisted   = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "帖子不存在")

	ErrUserInvalid    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不可用")
	ErrUserNotMatched = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不匹配")

	ErrCommentNotFound = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "评论没有找到")
	ErrHaveLiked       = errors.New(ErrCodeNormal, v1.ErrorReason_REPEATED_OPERATION.String(), "你已经点过赞了")

	ErrNeedMFA          = errors.New(ErrCodeNeedMFA, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "需要MFA")
	ErrHaveActivatedMFA = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "已经激活了MFA")
	ErrNotEnabledMFA    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "尚未激活MFA")
	ErrMFAVerifyFailed  = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "动态码验证失败")
)

func GenerateError(err Error) *errors.Error {
	return errors.New(err.Code, err.Reason, err.Message)
}

// 产生默认错误信息
func GenerateErrorNormal(err error) *errors.Error {
	return errors.New(ErrCodeInternalServerError, v1.ErrorReason_NORMAL_ERROR.String(), err.Error())
}

// 从文本直接新建错误
func GenerateErrorString(err string) *errors.Error {
	return errors.New(ErrCodeInternalServerError, v1.ErrorReason_NORMAL_ERROR.String(), err)
}
