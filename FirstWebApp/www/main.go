package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avgGrades, happiness float64
}

func main() {
	handleRequest()
}

// Handle Request
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

// Home Page
func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, "<b>Main Text</b>")
	template, _ := template.ParseFiles("templates/homePage.html")
	template.Execute(w, bob)
}

// Contacts Page
func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page")
}
