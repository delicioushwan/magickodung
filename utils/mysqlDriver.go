package utils

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger : logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(entities.User{}, entities.Question{}, entities.Option{}, entities.Answer{})
}


func InitTtestDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.TEST_DB_Port,
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

func RunInTx(ctx context.Context, db *gorm.DB, opts *sql.TxOptions, f func(txDb *gorm.DB) error) error {
	tx := db.WithContext(ctx).Begin(opts)
	if tx.Error != nil {
		return fmt.Errorf("start tx: %v", tx.Error)
	}

	if err := f(tx); err != nil {
		if err1 := tx.Rollback().Error; err1 != nil {
			return fmt.Errorf("rollback tx: %v (original error: %v)", err1, err)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit tx: %v", err)
	}
	return nil
}
