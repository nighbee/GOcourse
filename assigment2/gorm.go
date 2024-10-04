package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	Name string `gorm:"unique;not null"`
	Age  int    `gorm:"not null"`
}

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate the schema: %v", err)
	}
	fmt.Println("Users table created successfully")

	// Insert data
	users := []User{
		{Name: "John Doe", Age: 30},
		{Name: "Jane Smith", Age: 25},
		{Name: "Bob Johnson", Age: 40},
		{Name: "Alice Williams", Age: 35},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatalf("Failed to insert users: %v", result.Error)
	}
	fmt.Println("Users inserted successfully")

	// Query data
	var allUsers []User
	result = db.Find(&allUsers)
	if result.Error != nil {
		log.Fatalf("Failed to query users: %v", result.Error)
	}
	fmt.Println("\nAll users:")
	for _, user := range allUsers {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Query data with filtering
	var usersOverThirty []User
	result = db.Where("age >= ?", 30).Find(&usersOverThirty)
	if result.Error != nil {
		log.Fatalf("Failed to query users: %v", result.Error)
	}
	fmt.Println("\nUsers over 30 years old:")
	for _, user := range usersOverThirty {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Update data
	db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "John Smith", Age: 35})
	fmt.Println("\nUser with ID 1 updated successfully")

	// Delete data
	db.Delete(&User{}, 2)
	fmt.Println("User with ID 2 deleted successfully")
}
