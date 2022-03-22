package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseMysql() {
	var err error
	const MYSQL = "dhikrama:database@tcp(127.0.0.1:3306)/dhikrama?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database ")
	}

	fmt.Println("connect to database")

}
