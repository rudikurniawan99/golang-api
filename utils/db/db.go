package db

import (
	"api-2/src/model"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbGorm struct {
	master *gorm.DB
}

func NewDBGorm() *dbGorm {
	var loadOnce sync.Once
	db := &dbGorm{}

	loadOnce.Do(func() {
		db.master = Connect()
	})

	return db
}

func (d *dbGorm) Master() *gorm.DB {
	return d.master
}

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Println("===")
		log.Println("= DB connection failed")
		log.Println(err)
		log.Println("===")
	} else {
		log.Println("===")
		log.Println("= DB connection success")
		log.Println("===")
	}

	return db
}
