package db

import (
	"USI-grpc-Service/config"
	dbmodel "USI-grpc-Service/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect to the mysql database and return gorm connection
func Connect(config config.IConfig, log *zap.Logger) (db *gorm.DB, err error) {
	// connect to the database
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetDatabase().Username,
		config.GetDatabase().Password,
		config.GetDatabase().Host,
		config.GetDatabase().Port,
		config.GetDatabase().Scheme,
	)
	db, err = gorm.Open(
		mysql.Open(dns),
		&gorm.Config{},
	)
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	log.Info("database connected", zap.String("dns", dns))

	return db, err
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&dbmodel.RegisterUser{})
}
