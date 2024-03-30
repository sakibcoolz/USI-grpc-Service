package domain

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IDomain interface {
}

type Domain struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDomain(db *gorm.DB, log *zap.Logger) IDomain {
	return &Domain{
		db:  db,
		log: log,
	}
}
