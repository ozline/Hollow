package service

import (
	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrMissingParams = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "Missing Params Data")
	ErrParamsIllegal = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "params invalid")
)
