package controller

import (
	"USI-grpc-Service/config"
	"USI-grpc-Service/model/modelpb/login"
	"USI-grpc-Service/service"
	"context"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Controller struct {
	service  service.IService
	config   config.IConfig
	log      *zap.Logger
	validate *validator.Validate
	login.UnimplementedUSIServer
}

func NewController(service service.IService, config config.IConfig, log *zap.Logger, validate *validator.Validate) login.USIServer {
	return &Controller{
		service:  service,
		config:   config,
		log:      log,
		validate: validate,
	}
}

func (c *Controller) Login(ctx context.Context, in *login.LoginRequest) (*login.LoginResponse, error) {
	c.log.Info("Login received ", zap.String("username", in.Username))
	return &login.LoginResponse{}, nil
}
