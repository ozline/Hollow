package service

import (
	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrMissingParams = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "missing params data")
	ErrParamsIllegal = errors.New(int(v1.ErrorReason_PARAMS_ILLEGAL.Number()), v1.ErrorReason_PARAMS_ILLEGAL.String(), "params invalid")
)
