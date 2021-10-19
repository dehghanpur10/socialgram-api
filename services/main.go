package services

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"socialgram/lib"
)

func Connect() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", lib.DB_USER, lib.DB_PASSWORD, lib.DB_HOST, lib.DB_PORT, lib.DB_NAME)
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migration(database)
	return database, nil
}

func migration(db *gorm.DB) {
	//db.AutoMigrate(&models.User{})
}
