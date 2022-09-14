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

	ErrCodeNeedMFA = 417

	//SERVICE
	ErrMissingParams = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "缺失必要参数")
	ErrParamsIllegal = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "参数非法")

	//BIZ
	ErrUserBlocked         = errors.New(ErrCodeNormal, v1.ErrorReason_USER_BLOCKED.String(), "用户被封禁")
	ErrUserCheckFailed     = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "账号验证失败")
	ErrUserNotExist        = errors.New(ErrCodeNormal, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "用户没有找到")
	ErrUserExisted         = errors.New(ErrCodeNormal, v1.ErrorReason_PARAMS_ILLEGAL.String(), "用户名已经存在")
	ErrMsgCodeVerifyFailed = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "验证码验证失败")
	ErrMsgCodeInvalid      = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "验证码已过期或不存在")

	//DATA
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "未知错误")

	ErrUserInsufficientPermissions = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户权限不足")
	ErrUserNotExisted              = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不存在")
	ErrFatherNotExisted            = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "父楼不存在")
	ErrLeafNotExisted              = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "帖子不存在")

	ErrUserInvalid    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不可用")
	ErrUserNotMatched = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "用户不匹配")

	ErrCommentNotExist = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "评论不存在")
	ErrHaveLiked       = errors.New(ErrCodeNormal, v1.ErrorReason_REPEATED_OPERATION.String(), "你已经点过赞了")

	ErrNeedMFA          = errors.New(ErrCodeNeedMFA, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "需要MFA")
	ErrHaveActivatedMFA = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "已经激活了MFA")
	ErrNotEnabledMFA    = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "尚未激活MFA")
	ErrMFAVerifyFailed  = errors.New(ErrCodeNormal, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "动态码验证失败")

	ErrReportsNotExisted = errors.New(ErrCodeNormal, v1.ErrorReason_ITEM_EXISTED.String(), "投诉不存在")
	ErrReportExisted     = errors.New(ErrCodeNormal, v1.ErrorReason_ITEM_EXISTED.String(), "存在同一个ID且尚未通过的投诉,若需要再次投诉请等待审核通过")
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
