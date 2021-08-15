package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Dsn string
}

func (p Postgres) DBInit() *gorm.DB {
	db, err := gorm.Open(postgres.Open(p.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connected to database")
	}

	log.Print("Connected to database")

	return db
}
