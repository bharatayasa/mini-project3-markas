package config

import (
	"fmt"
	"log"
	"os"

	"github.com/bharatayasa/mini-project3-markas/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDb() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	mysqlconn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: mysqlconn,
	}

	err = autoMigrate(mysqlconn)
	if err != nil {
		log.Fatal(err)
	}
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Books{},
	)

	if err != nil {
		return err
	}

	return nil
}
