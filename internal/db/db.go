package db

import (
	"fmt"
	"log"

	"github.com/killerekt/grpc-prac/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit(config config.Config) (err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
		config.PostgresPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return
	}

	log.Println("Successfully connected to the database")
	return
}
