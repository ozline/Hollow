// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"hollow/internal/biz"
	"hollow/internal/conf"
	"hollow/internal/data"
	"hollow/internal/server"
	"hollow/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	bucket := data.NewOSS(confData)
	conn := data.NewRedis(confData)
	aliyunShortMsg := data.NewShortMsg(confData)
	node := data.NewSnowflake(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, bucket, conn, aliyunShortMsg, node)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase)
	forestRepo := data.NewForestRepo(dataData, logger)
	forestUsecase := biz.NewForestUsecase(forestRepo, logger)
	forestService := service.NewForestService(forestUsecase)
	grpcServer := server.NewGRPCServer(confServer, auth, userService, forestService, logger)
	httpServer := server.NewHTTPServer(confServer, auth, userService, forestService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
