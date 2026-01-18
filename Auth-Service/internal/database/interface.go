package database

import (
	"gorm.io/gorm"
)

type ISqlDB interface {
	GetDB() (*gorm.DB, error)
}
