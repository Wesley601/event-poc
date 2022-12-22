package main

import (
	"encoding/json"
	"log"
	"net/http"
	"wesley601/event-driver/internal/api"
	"wesley601/event-driver/pkg/bus"
	"wesley601/event-driver/pkg/db"
)

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	eBus, err := bus.New()
	if err != nil {
		panic(err)
	}
	defer eBus.Close()

	userController := api.New(dbConnection, eBus)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var user api.User

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				panic(err)
			}

			userController.CreateUser(user)
			w.WriteHeader(http.StatusCreated)
		}

		if r.Method == "GET" {
			users, err := userController.ListUsers()
			if err != nil {
				panic(err)
			}

			json.NewEncoder(w).Encode(users)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
