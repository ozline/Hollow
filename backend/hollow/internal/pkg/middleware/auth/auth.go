package auth

import (
	"context"
	v1 "hollow/api/hollow/v1"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt"
)

var currentUserKey struct{}

type CurrentUser struct {
	ID       int64  //用户ID
	Username string //用户名
	Status   int64  //用户状态
}

var (
	ErrorCode       = 500 //JWT解析错误返回状态码
	ErrMissingToken = errors.New(ErrorCode, v1.ErrorReason_PARAMS_ILLEGAL.String(), "Token Missing")
	ErrInvaildToken = errors.New(ErrorCode, v1.ErrorReason_PARAMS_ILLEGAL.String(), "Token Invalid")
)

type JWTClaims struct {
	Id       int64  `json:"id" example:"-1"`    //用户ID
	Username string `json:"username"`           //用户名
	Status   int64  `json:"status" example:"0"` //用户状态
	jwt.StandardClaims
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var JWTToken string
			if md, ok := metadata.FromServerContext(ctx); ok {
				JWTToken = md.Get("x-md-global-token")
			} else if tr, ok := transport.FromServerContext(ctx); ok {
				JWTToken = tr.RequestHeader().Get("Authorization")
			} else {
				return nil, ErrMissingToken
			}

			if len(JWTToken) == 0 {
				return nil, ErrMissingToken
			}

			claims, err := JWTParse(JWTToken, secret)

			if err != nil {
				return nil, errors.New(ErrorCode, v1.ErrorReason_PARAMS_ILLEGAL.String(), err.Error())
			}

			if claims.ExpiresAt < time.Now().Unix() {
				return nil, ErrInvaildToken
			}

			ctx = WithContext(ctx, &CurrentUser{
				ID:       claims.Id,
				Username: claims.Username,
				Status:   claims.Status,
			})

			tr, ok := transport.FromServerContext(ctx)

			if !ok {
				return nil, ErrInvaildToken
			}

			// Token持久化
			newToken, err := GetAuthToken(claims.Id, claims.Username, claims.Status, secret)

			if err != nil {
				return nil, ErrInvaildToken
			}

			tr.ReplyHeader().Set("Authorization", newToken)

			return handler(ctx, req)
		}
	}
}

func JWTParse(token string, secret string) (*JWTClaims, error) {
	tmp, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if tmp != nil {
		claims, result := tmp.Claims.(*JWTClaims)
		if result && tmp.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWTGenerate(claims JWTClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GetAuthToken(id int64, username string, status int64, secret string) (string, error) {
	claims := JWTClaims{
		Id:       id,
		Username: username,
		Status:   status,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 3600), //签名生效时间 一小时前
			ExpiresAt: int64(time.Now().Unix() + 3600), //签名过期时间 一小时后
		},
	}
	token, err := JWTGenerate(claims, secret)
	return token, err
}

// 从Context提取用户信息
func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

// 给Context附带用户信息
func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}

func GetUserInfo(ctx context.Context) *CurrentUser {
	currentUser := FromContext(ctx)
	return currentUser
}
