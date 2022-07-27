package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Username string
}

func (u User) String() string {
	return fmt.Sprintf("User: id=%d, username=%s", u.Id, u.Username)
}

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// NOTE: my username and password is "ellie"; if yours isn't - please change those parts!
	connectionDeets := "ellie:ellie@tcp(127.0.0.1:3306)/gobase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionDeets), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&User{Username: "Watermelon"})

	// Read
	var user User
	db.First(&user) // find first user
	fmt.Println("before update:")
	fmt.Println(user)

	// Update
	db.Model(&user).Update("Username", "Pineapple")
	fmt.Println("after update:")
	fmt.Println(user)

	// Delete - delete user
	// Delete with additional conditions
	db.Where("username = ?", "Pineapple").Delete(&user)
	fmt.Println("hi there!")
}
