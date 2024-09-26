package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "72zv5u3xp"
	dbname   = "go"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	db.AutoMigrate(&User{})

	users := []User{
		{Name: "John", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 35},
		{Name: "Charlie", Age: 40},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal("Error inserting user: ", result.Error)
		}
		fmt.Printf("inserted user %s with age %d\n", user.Name, user.Age)
	}
	var allUsers []User
	db.Find(&allUsers)

	fmt.Println("Users:")
	for _, user := range allUsers {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

}
