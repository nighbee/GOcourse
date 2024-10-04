package main

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"net/http"
)

func connectToDB() (*sql.DB, error){
	psqlInfo := "host=localhost port=5432 user=postgres password=72zv5u3xp dbname=go sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type User struct {
	id int 'json:"id"'
	name string 'json:"name"'
	age int 'json:"age"'
}

var db *sql.DB

func main() {
	var err error
	db, err = connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/users", updateUser)
	http.HandleFunc("/users", deleteUser)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
