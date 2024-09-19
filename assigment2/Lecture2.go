package main

import "fmt"

type University interface {
	Print()
}
type User struct {
	FullName string
	Email    string
	password string // private field
	University
}

type Login interface {
	login(email, password string) bool
}

func (u User) GetPassword() string {
	return u.password
}
func (u User) SetPassword(password string) {
	u.password = password
}

type Manager struct {
	User
	ManagerId string
}

type Student struct {
	User
	StudentId string
}

func Print(u University) {
	switch v := u.(type) {
	case Student:
		fmt.Printf("Student Details:\nName: %s\nEmail: %s\nStudent ID: %s\n\n", v.FullName, v.Email, v.StudentId)
	case Manager:
		fmt.Printf("Manager Details:\nName: %s\nEmail: %s\nManager ID: %s\n\n", v.FullName, v.Email, v.ManagerId)
	default:
		fmt.Println("Unknown type")
	}
}

func (u User) Login(email, password string) bool {
	return u.Email == email && u.password == password

}

func main() {
	s := Student{
		User: User{
			FullName: " Almaz",
			Email:    "a@kbtu",
		},
		StudentId: "22B031193",
	}
	s.SetPassword("123456")

	manager := Manager{
		User: User{
			FullName: "Aigerim",
			Email:    "aigera@kbtu",
			//manager.SetPassword("123456"),
		},
		ManagerId: "123456",
	}
	manager.SetPassword("123456")

	Print(s)
	Print(manager)
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	if s.Login(email, password) {
		fmt.Println("Student login successful")
	}
	if manager.Login(email, password) {
		fmt.Println("Manager login successful")
	} else {
		fmt.Println("Login failed")
	}
}
