package server

import (
	"USI-grpc-Service/config"
	"USI-grpc-Service/controller"
	dbm "USI-grpc-Service/db"
	"USI-grpc-Service/domain"
	"USI-grpc-Service/model/modelpb/login"
	"USI-grpc-Service/service"
	"fmt"
	"net"
	"os"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Server(config config.IConfig, log *zap.Logger) {
	// connect to database
	db, err := dbm.Connect(config, log)
	if err != nil {
		log.Fatal("error while connecting to database", zap.Error(err))
	}

	validates := validator.New(validator.WithRequiredStructEnabled())

	store := domain.NewDomain(db, log)

	srv := service.NewService(store, log, config)

	gRPCServer := grpc.NewServer()

	controller := controller.NewController(&srv, config, log, validates)

	login.RegisterUSIServer(gRPCServer, controller)

	serviceHost := fmt.Sprintf("%s:%s", config.GetService().Host, config.GetService().Port)

	con, err := net.Listen("tcp", serviceHost)
	if err != nil {
		log.Fatal("error while configure listen service ", zap.String("hostAt", serviceHost), zap.Error(err))
	}

	log.Info("serving service", zap.String("hostAt", serviceHost))

	err = gRPCServer.Serve(con)
	if err != nil {
		log.Fatal("error while serving service", zap.Error(err))
	}

	log.Info("serving service terminated", zap.String("hostAt", serviceHost))

}
