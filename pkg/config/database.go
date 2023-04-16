package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database(dns string) *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil
	}
	return db
}
