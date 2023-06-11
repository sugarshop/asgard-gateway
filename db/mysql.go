package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlInit 数据库初始化连接
func MysqlInit() {
	dsn := "94p2ws6zyeusmkkmiw94:pscale_pw_fQQyJtbCQZbp224O7ARJ5FYKxo4BWLcuTm1uURzr6rf@tcp(aws.connect.psdb.cloud)/sugarshop?tls=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 连接池设置
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("connect to db APIGATEWAYDB db success")
	sugarshopDB = db
}

// sugar shop
var sugarshopDB *gorm.DB

// SugarShopDB sugarshop db
func SugarShopDB() *gorm.DB {
	return sugarshopDB
}
