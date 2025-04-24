package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "goapi/docs" // Swagger 文件
	"goapi/model"
	"goapi/router"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

// @title Go API (MC 架構)
// @version 1.0
// @description 使用 Gin 搭建的 MVC 架構 API 專案
// @host localhost:8080
// @BasePath /
func ensureDatabaseExists() {
	err := godotenv.Load()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, name)
	masterDSN := dsn
	db, err := gorm.Open(sqlserver.Open(masterDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("連接 master 失敗：%v", err)
	}

	var exists int64
	db.Raw("SELECT COUNT(*) FROM sys.databases WHERE name = ?", "gobase").Scan(&exists)
	if exists == 0 {
		db.Exec("CREATE DATABASE gobase")
		log.Println("✅ gobase 資料庫已建立")
	} else {
		log.Println("🔍 gobase 資料庫已存在")
	}
}

func main() {
	// 先檢查資料庫是否存在，若無則建立
	ensureDatabaseExists()
	err := godotenv.Load()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("連接資料庫失敗：" + err.Error())
	}

	// 自動建表
	err = db.AutoMigrate(&model.Form{})
	if err != nil {
		panic("資料表建立失敗：" + err.Error())
	}

	r := router.SetupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
