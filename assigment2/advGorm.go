package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name    string
	Age     int
	Profile Profile
}

// Profile model
type Profile struct {
	gorm.Model
	UserID            uint
	Bio               string
	ProfilePictureURL string
}

func main() {
	// Connect to PostgreSQL database
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Migrate the schema
	db.AutoMigrate(&User{}, &Profile{})

	// Insert data with associations
	user := User{Name: "John Doe", Age: 30}
	profile := Profile{Bio: "Software Engineer", ProfilePictureURL: "https://example.com/profile.jpg"}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		profile.UserID = user.ID
		if err := tx.Create(&profile).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Query data with associations
	var users []User
	err = db.Preload("Profile").Find(&users).Error
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		fmt.Printf("User: %s (%d), Profile: %s (%s)\n", user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}

	// Update data
	updateUser := User{Model: gorm.Model{ID: user.ID}}
	updateProfile := Profile{Model: gorm.Model{ID: profile.ID}, Bio: "Senior Software Engineer"}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&updateUser).Updates(User{Name: "John Doe Jr."}).Error; err != nil {
			return err
		}
		if err := tx.Model(&updateProfile).Updates(Profile{Bio: "Senior Software Engineer"}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Delete data
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", user.ID).Delete(&User{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
