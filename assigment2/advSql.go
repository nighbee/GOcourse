package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "72zv5u3xp"
	dbname   = "go"
)

// Connect to the PostgreSQL database
func connectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Connection pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db, nil
}

// Create a table with constraints
func createUsersTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE,
			age INTEGER NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Users table created successfully")
	return nil
}

// Insert data with transactions
func insertUsers(db *sql.DB, users []User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback if any error occurs

	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	for _, user := range users {
		_, err = tx.Exec(query, user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	fmt.Println("Users inserted successfully")
	return nil
}

// Query data with filtering and pagination
func queryUsers(db *sql.DB, ageFilter int, limit, offset int) ([]User, error) {
	var users []User
	var query string
	if ageFilter > 0 {
		query = `SELECT id, name, age FROM users WHERE age >= $1 
                                ORDER BY id LIMIT $2 OFFSET $3`
		rows, err := db.Query(query, ageFilter, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			err = rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
	} else {
		query = `SELECT id, name, age FROM users ORDER BY id LIMIT $1 OFFSET $2`
		rows, err := db.Query(query, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			err = rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
	}

	return users, nil
}

// Update user details
func updateUser(db *sql.DB, id int, newName string, newAge int) error {
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	_, err := db.Exec(query, newName, newAge, id)
	if err != nil {
		return err
	}
	fmt.Printf("User with ID %d updated successfully\n", id)
	return nil
}

// Delete user by ID
func deleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Printf("User with ID %d deleted successfully\n", id)
	return nil
}

// User struct
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the users table
	err = createUsersTable(db)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	// Insert users within a transaction
	users := []User{
		{Name: "John Doe", Age: 30},
		{Name: "Jane Smith", Age: 25},
		{Name: "Bob Johnson", Age: 40},
		{Name: "Alice Williams", Age: 35},
	}
	err = insertUsers(db, users)
	if err != nil {
		log.Fatalf("Failed to insert users: %v", err)
	}

	// Query users with filtering and pagination
	fmt.Println("\nUsers over 30 years old (Limit 2, Offset 0):")
	usersOverThirty, err := queryUsers(db, 30, 2, 0)
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}
	for _, user := range usersOverThirty {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	fmt.Println("\nAll users (Limit 2, Offset 1):")
	allUsers, err := queryUsers(db, 0, 2, 1)
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}
	for _, user := range allUsers {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Update user details
	err = updateUser(db, 1, "John Smith", 35)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	// Delete user by ID
	err = deleteUser(db, 2)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
}
