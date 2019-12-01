package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (d *DBDriver) Login(email string, password string) (string, error) {
	var data SignInCreds
	err := d.Conn.Get(&data, `SELECT id, password_hash FROM public."user" WHERE email=$1`, email)

	if err != nil {
		fmt.Println(err)
		return "", BadRequestError
	}

	if err == sql.ErrNoRows {
		fmt.Println("no rows here")
		return "", EmptyRows

	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(passwordHash)

	err = bcrypt.CompareHashAndPassword(data.PasswordHash, []byte(password))
	if err != nil {
		fmt.Println(err, data.PasswordHash, []byte(password))
		return "", IncorrectPasswordFormat
	}

	return data.ID, nil
}

func (d *DBDriver) Signup(username string, email string, password string) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	userId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	idstr := userId.String()

	_, err = d.Conn.NamedExec(`INSERT INTO public."user" (id, username, email, password_hash) VALUES (:ID, :Username, :Email, :PasswordHash)`,
		map[string]interface{}{
			"ID":           idstr,
			"Username":     username,
			"Email":        email,
			"PasswordHash": passwordHash,
		})

	if err != nil {
		return false, err
	}
	return true, nil
}
