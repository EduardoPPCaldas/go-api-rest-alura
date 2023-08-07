package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func GetDatabase() *gorm.DB {
	dbName := "personalitydb"
	createDBDsn := "user:user@tcp(localhost:3306)/"
	db, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	_ = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	dsn := "user:user@tcp(localhost:3306)/personalitydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	return db
}
