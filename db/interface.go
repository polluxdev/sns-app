package db

import "gorm.io/gorm"

type DBInterface interface {
	DBInit() *gorm.DB
}
