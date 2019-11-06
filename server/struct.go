package main

import "github.com/jmoiron/sqlx"

type DBDriver struct {
	Conn *sqlx.DB
}

type API struct {
	DB *DBDriver
}

type userData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json: "email`
}

type userSignup struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
