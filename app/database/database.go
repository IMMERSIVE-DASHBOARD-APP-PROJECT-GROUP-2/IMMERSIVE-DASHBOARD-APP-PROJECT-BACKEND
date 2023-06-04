package database

import (
	"fmt"

	"github.com/DASHBOARDAPP/app/config"
	classdata "github.com/DASHBOARDAPP/features/class/data"
	logdata "github.com/DASHBOARDAPP/features/log/data"
	menteedata "github.com/DASHBOARDAPP/features/mentee/data"
	userdata "github.com/DASHBOARDAPP/features/user/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitialMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userdata.User{},
		&classdata.Class{},
		&menteedata.Mentee{},
		&logdata.Log{},
	)
	if err != nil {
		return err
	}
	return nil
}
