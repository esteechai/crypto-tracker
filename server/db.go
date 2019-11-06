package main

import "fmt"

func (d *DBDriver) Login(email string, password string) ([]userData, error) {
	var passback []userData

	err := d.Conn.Select(&passback, `SELECT id, username, password from users WHERE email=$1 AND password=$2 `, email, password)
	if err != nil {
		return nil, err
	}
	return passback, nil
}

func (d *DBDriver) Signup(id string, username string, email string, password string) ([]userSignup, error) {
	var passback []userSignup
	err := d.Conn.Select(&passback, `INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)`, id, username, email, password)
	if err != nil {
		fmt.Println(err)
	}
	return passback, nil
}
