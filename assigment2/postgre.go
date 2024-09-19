package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "72zv5u3xp"
	dbname   = "postgres"
)

func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func createTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name varchar(50),age INT)"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table created successfully!")
	return nil
}

func insertData(db *sql.DB, name string, age int) error {
	query := "insert into users(name, age) values($1,$2)"
	_, err := db.Exec(query, name, age)
	if err != nil {
		return err
	}
	fmt.Printf("inserted user %s with age %d\n", name, age)
	return nil
}

func printAllData(db *sql.DB) {
	query := "select * from users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Printf("users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
}

type users struct {
	name string
	age  int
}

func main() {

	db, err := connect()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	user := users{name: "John", age: 25}
	createTable(db)
	insertData(db, user.name, user.age)
	printAllData(db)

}
