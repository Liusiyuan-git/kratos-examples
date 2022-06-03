// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"test/internal/biz"
	"test/internal/conf"
	"test/internal/server"
	"test/internal/service"
)

// Injectors from wire.go:

func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	greeterUsecase := biz.NewGreeterUsecase(logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
	}, nil
}
