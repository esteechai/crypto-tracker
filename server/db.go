package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

//login
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

//signup
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

func (d *DBDriver) UpdateTicker(tickerData *TickerData) {
	var test string
	var query string

	err := d.Conn.Get(&test, `SELECT ticker_id FROM product_ticker WHERE ticker_id=$1`, tickerData.ID)
	if err != nil {
		query = `INSERT INTO product_ticker (ticker_id, price, size, time,  bid, ask, volume) VALUES(:ID, :Price, :Size, :Time, :Bid, :Ask, :Volume)`
	} else {
		query = `UPDATE product_ticker SET price=:Price, size=:Size, time=:Time, bid=:Bid, ask=:Ask, volume=:Volume WHERE ticker_id=:ID`
	}

	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"ID":     tickerData.ID,
		"Price":  tickerData.Price,
		"Size":   tickerData.Size,
		"Time":   tickerData.Time,
		"Bid":    tickerData.Bid,
		"Ask":    tickerData.Ask,
		"Volume": tickerData.Volume,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (d *DBDriver) SelectedProduct(id string) (*TickerData, error) {
	data := &TickerData{}
	err := d.Conn.Get(data, `SELECT ticker_id,price,size,time,bid,ask,volume FROM product_ticker WHERE ticker_id = $1`, id)

	fmt.Println("Selected Product: ", data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		fmt.Println("no rows here")
		return nil, err
	}

	fmt.Println(data)
	return data, nil
}
