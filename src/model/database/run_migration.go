package database

import (
	"fmt"
	"log"
	"os/user"
)

func RunMigration() {
	err := DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("database migrate")
}
