package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:root@tcp(127.0.0.1:3306)/springboot?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDatabase() (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, nil
	}
	return db, nil
}
