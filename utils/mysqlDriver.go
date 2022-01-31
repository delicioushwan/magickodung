package utils

import (
	"fmt"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)
	// "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(entities.User{})
}
