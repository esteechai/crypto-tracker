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
	err := d.Conn.Get(&data, `SELECT id, password_hash FROM users WHERE email=$1`, email)

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

	//generate user ID
	userId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	idstr := userId.String()

	//generate verfication token
	verificationToken, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate verfication_token: %v", err)
	}
	vToken := verificationToken.String()

	//generate reset password token
	resetPasswordToken, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Failed to generate reset_password_token: %v", err)
	}
	passToken := resetPasswordToken.String()

	query := `INSERT INTO users (id, email, username, password_hash, verification, verification_token, reset_pass_token) VALUES (:ID, :Email, :Username, :PasswordHash, :Verification, :VerificationToken, :ResetPassToken)`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"ID":                idstr,
		"Email":             email,
		"Username":          username,
		"PasswordHash":      passwordHash,
		"Verification":      false,
		"VerificationToken": vToken,
		"ResetPassToken":    passToken,
	})

	fmt.Println(err)

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

func (d *DBDriver) AddFav(userId string, productID string) error {
	favId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate fav_id: %v", err)
	}
	fav_id := favId.String()

	query := `INSERT INTO users_favourites (fav_id, user_id, ticker_id) VALUES (:FavID, :UserID, :ProductID)`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"FavID":     fav_id,
		"UserID":    userId,
		"ProductID": productID,
	})

	if err != nil {
		fmt.Println(err)
		return AddFavProductError
	}
	fmt.Println("successfully added: ", productID)
	return nil
}

func (d *DBDriver) RemoveFav(userId string, productID string) error {
	query := `DELETE FROM users_favourites WHERE user_id=$1 AND ticker_id=$2`
	_, err := d.Conn.Exec(query, userId, productID)

	if err != nil {
		fmt.Println(err)
		return RemoveFavProductError
	}
	fmt.Println("successfully remove: ", productID)
	return nil
}

func (d *DBDriver) CheckFav(userID, productID string) (*[]FavProducts, error) {
	var favID string
	err := d.Conn.Get(&favID, `SELECT fav_id FROM users_favourites WHERE user_id=$1 AND ticker_id=$2`, userID, productID)

	fmt.Println("check prod: ", productID)

	if err != nil {
		if err == sql.ErrNoRows {
			d.AddFav(userID, productID)
		} else {
			fmt.Println(err)
			return nil, BadRequestError
		}
	} else {
		d.RemoveFav(userID, productID)
	}

	data, err := d.GetFavProducts(userID)
	if err != nil {
		fmt.Println(err)
		return nil, BadRequestError
	}
	return data, nil
}

func (d *DBDriver) GetFavProducts(id string) (*[]FavProducts, error) {
	data := &[]FavProducts{}

	err := d.Conn.Select(data, `
	SELECT p.ticker_id, p.price, p.size, p.time, p.bid, p.ask, p.volume 
	FROM users_favourites uf, product_ticker p 
	WHERE uf.ticker_id = p.ticker_id
	AND uf.user_id=$1`, id)

	if err == sql.ErrNoRows {
		fmt.Println("fav list is empty")
		return nil, EmptyFavProductList
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
