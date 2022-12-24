package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"wesley601/event-driver/pkg/bus"
)

type UserController struct {
	db       *sql.DB
	eventBus bus.EventBus
}

type User struct {
	ID     int64
	Name   string
	Status string
}

func New(db *sql.DB, eventBus bus.EventBus) *UserController {
	return &UserController{
		db:       db,
		eventBus: eventBus,
	}
}

func (uc *UserController) CreateUser(user User) error {
	stmt, err := uc.db.Prepare("INSERT INTO users(name, status) VALUES(?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Name, user.Status)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	payload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	uc.eventBus.Publish("user.created", payload)

	return nil
}

func (uc *UserController) ListUsers() ([]User, error) {
	fmt.Println("aqui?")
	var users []User
	rows, err := uc.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Status); err != nil {
			return users, err
		}
		fmt.Printf("user: %v\n", user)
		users = append(users, user)
	}

	return users, nil
}
