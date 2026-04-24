package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db_name, db_user, db_password, db_host, db_port string) {
	db, err := gorm.Open(mysql.Open(db_user+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_name+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
	fmt.Println("Database MySQL terkoneksi")
}
