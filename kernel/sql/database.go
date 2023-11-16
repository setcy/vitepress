package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string

	db *gorm.DB
)

func init() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
}

func InitDatabase() {
	var err error
	dsn := "host=" + dbHost +
		" port=" + dbPort +
		" user=" + dbUser +
		" password=" + dbPass +
		" dbname=" + dbName +
		" sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
