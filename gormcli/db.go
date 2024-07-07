package gormcli

import (
	"bubble/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func OpenDb() {
	dbConf := configs.GetGlobalConfig().DbConfig
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User, dbConf.PassWord, dbConf.Host, dbConf.Port, dbConf.Dbname)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Open database err")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Get DB err")
	}

	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbConf.MaxIdleTime * int(time.Second)))
}

func GetDB() *gorm.DB {
	once.Do(OpenDb)
	return db
}
