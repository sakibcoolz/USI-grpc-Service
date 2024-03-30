package service

import (
	"go.uber.org/zap"

	"USI-grpc-Service/config"
	"USI-grpc-Service/domain"
)

type IService interface {
}

type Service struct {
	store domain.IDomain
	log   *zap.Logger
}

func NewService(store domain.IDomain, log *zap.Logger, config config.IConfig) IService {
	return &Service{
		store: store,
		log:   log,
	}
}
