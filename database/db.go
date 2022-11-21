package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	strConnetion := "host=localhost user=dbpersons password=dbpersons dbname=dbpersons port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnetion))
	if err != nil {
		log.Panic("Database connection error", err)
	}
}
