package main

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "John", age: 25})
	var user User
	db.First(&user, 1)
	db.Model(&user).Update("Name", "David")
	db.Delete(&user, 1)

}
