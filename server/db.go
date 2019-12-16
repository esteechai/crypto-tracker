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
		return "", InvalidEmailOrPassword
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
		return "", IncorrectPassword
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

	if err != nil {
		if IsUniqueConstraintError(err, UniqueConstraintUsername) {
			fmt.Println("violate username uq")
			return false, ViolateUNUsername
		}
		if IsUniqueConstraintError(err, UniqueConstraintEmail) {
			fmt.Println("violate email uq")
			return false, ViolateUNEmail
		}
		return false, err
	}
	return true, nil
}

//get user's verification token after signup
func (d *DBDriver) GetVerifToken(username string, email string) (string, error) {
	var verifToken string
	err := d.Conn.Get(&verifToken, `SELECT verification_token FROM users WHERE username=$1 AND email=$2`, username, email)

	if err != nil {
		return "", BadRequestError
	}
	return verifToken, nil
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

//set user's verification bool to true
func (d *DBDriver) VerifyUserAcc(veriToken string) error {
	query := `UPDATE users SET verification = true WHERE verification_token = :VerificationToken`
	_, err := d.Conn.NamedExec(query, map[string]interface{}{
		"VerificationToken": veriToken,
	})

	if err != nil {
		return err
	}
	return nil
}

func (d *DBDriver) ResetPassword(userID string, currentPw string, newPw string) error {
	var pw []byte
	err := d.Conn.Get(&pw, `SELECT password_hash FROM users WHERE id=$1`, userID)
	if err != nil {
		fmt.Println(err)
		return BadRequestError
	}

	err = bcrypt.CompareHashAndPassword(pw, []byte(currentPw))
	if err != nil {
		fmt.Println(err)
		return PasswordMatchingIssue
	}

	newPwHash, err := bcrypt.GenerateFromPassword([]byte(newPw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("this error: ", err)
		return ResetPasswordError
	}

	query := `UPDATE users SET password_hash = :NewPassword WHERE id = :UserID`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"NewPassword": newPwHash,
		"UserID":      userID,
	})
	if err != nil {
		fmt.Println("this : ", err)

		return ResetPasswordError
	}
	return nil
}

func (d *DBDriver) GetResetPassToken(email string) (string, error) {
	var token string
	err := d.Conn.Get(&token, `SELECT reset_pass_token FROM users WHERE email=$1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", RequestResetPassTokenError
		}
		fmt.Println(err)
		return "", BadRequestError
	}
	return token, nil
}

func (d *DBDriver) UpdatePassWToken(resetPassToken string, password string) error {
	var userID string

	//get user_id
	err := d.Conn.Get(&userID, `SELECT id FROM users WHERE reset_pass_token=$1`, resetPassToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return DbQueryError
		}
		fmt.Println(err)
		return BadRequestError
	}

	//hash new password
	newPwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("this error: ", err)
		return ResetPasswordError
	}

	//generate new reset password token
	newToken, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate new reset password token: %v", err)
	}
	new_token := newToken.String()

	//update new password hash and new reset password token
	query := `UPDATE users SET password_hash = :NewPassword, reset_pass_token = :NewToken WHERE id = :UserID`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"NewPassword": newPwHash,
		"NewToken":    new_token,
		"UserID":      userID,
	})
	if err != nil {
		fmt.Println("this : ", err)

		return ResetPasswordError
	}
	return nil
}
