package config

import (
	"fmt"
	"sns-app/db"
	"os"

	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("HOST_URL"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"), os.Getenv("DB_TIME_ZONE"),
	)

	var dbi db.DBInterface = db.Postgres{Dsn: dsn}

	return dbi.DBInit()
}
