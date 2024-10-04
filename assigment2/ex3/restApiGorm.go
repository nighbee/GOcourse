package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "72zv5u3xp"
	dbname   = "go"
)

// Connect to the PostgreSQL database
func connectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// User model
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate the schema: %v", err)
	}

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/user/", updateUser)
	http.HandleFunc("/user/", deleteUser)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
