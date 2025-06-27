package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my friend!")
}
func anotherPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye friend!")
}

func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		reqId := r.URL.Query().Get("id")
		user1 := User{
			ID:   1,
			Name: "Alex",
			Age:  25,
		}
		user2 := User{
			ID:   2,
			Name: "David",
			Age:  42,
		}
		w.Header().Set("Content-Type", "application/json")
		if reqId != "" {
			if reqIdInt, err := strconv.Atoi(reqId); err == nil {
				if reqIdInt == 1 {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user1)
				} else if reqIdInt == 2 {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user2)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			}
		} else {
			json.NewEncoder(w).Encode([]User{user1, user2})
		}
		return
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Your user is create")
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}

}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/anPage", anotherPage)
	http.HandleFunc("/user", user)

	log.Println("Srever is working on localhost:7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatalf("Server connection lost %v", err)
	}

}
