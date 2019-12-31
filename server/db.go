package main

import (
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

//Login func retrives user ID from login input credentials
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

	err = bcrypt.CompareHashAndPassword(data.PasswordHash, []byte(password))
	if err != nil {
		fmt.Println(err, data.PasswordHash, []byte(password))
		return "", IncorrectPassword
	}
	return data.ID, nil
}

//Signup func validates and inserts user data to db
func (d *DBDriver) Signup(username string, email string, password string) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return SignupError
	}

	//generates user ID
	userId, err := uuid.NewV4()
	if err != nil {
		return SignupError
	}

	idstr := userId.String()

	//generates verfication token
	verificationToken, err := uuid.NewV4()
	if err != nil {
		return SignupError
	}
	vToken := verificationToken.String()

	//generates reset password token
	resetPasswordToken, err := uuid.NewV4()
	if err != nil {
		return SignupError
	}

	passToken := resetPasswordToken.String()

	query := `INSERT INTO users (id, email, username, password_hash, verified, verification_token, reset_pass_token) VALUES (:ID, :Email, :Username, :PasswordHash, :Verification, :VerificationToken, :ResetPassToken)`
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
			return ViolateUNUsername
		}
		if IsUniqueConstraintError(err, UniqueConstraintEmail) {
			fmt.Println("violate email uq")
			return ViolateUNEmail
		}
		return SignupError
	}
	return nil
}

//GetVerifToken func retrieves user's verification token after signup
func (d *DBDriver) GetVerifToken(username string, email string) (string, error) {
	var verifToken string
	err := d.Conn.Get(&verifToken, `SELECT verification_token FROM users WHERE username=$1 AND email=$2`, username, email)

	if err != nil {
		return "", BadRequestError
	}
	return verifToken, nil
}

//AddProducts func adds all products retrieved from coinbase API intp db
func (d *DBDriver) AddProducts(products []Products) error {
	var query string
	var id string

	for _, product := range products {
		err := d.Conn.Get(&id, `SELECT id FROM products WHERE id=$1`, product.ID)
		// fmt.Println("id: ", product.ID)
		if err == sql.ErrNoRows {
			query = `INSERT INTO products (id, base_currency, quote_currency, base_min_size, base_max_size, quote_increment, base_increment, display_name, min_market_funds, max_market_funds, margin_enabled, post_only, limit_only, cancel_only, status, status_message) 
			VALUES(:ID, :BaseCurrency, :QuoteCurrency, :BaseMinSize, :BaseMaxSize, :QuoteIncrement, :BaseIncrement, :DisplayName, :MinMarketFunds, :MaxMarketFunds, :MarginEnabled, :PostOnly, :LimitOnly, :CancelOnly, :Status, :StatusMessage)`
		} else {
			query = `UPDATE products SET id=:ID, base_currency=:BaseCurrency, quote_currency=:QuoteCurrency, base_min_size=:BaseMinSize, base_max_size=:BaseMaxSize, quote_increment=:QuoteIncrement, base_increment=:BaseIncrement, display_name=:DisplayName,  min_market_funds=:MinMarketFunds, max_market_funds=:MaxMarketFunds, margin_enabled=:MarginEnabled, post_only=:PostOnly, limit_only=:LimitOnly, cancel_only=:CancelOnly, status=:Status, status_message=:StatusMessage
				WHERE id=:ID`

		}
		_, err = d.Conn.NamedExec(query, map[string]interface{}{
			"ID":             product.ID,
			"BaseCurrency":   product.BaseCurrency,
			"QuoteCurrency":  product.QuoteCurrency,
			"BaseMinSize":    product.BaseMinSize,
			"BaseMaxSize":    product.BaseMaxSize,
			"QuoteIncrement": product.QuoteIncrement,
			"BaseIncrement":  product.BaseIncrement,
			"DisplayName":    product.DisplayName,
			"MinMarketFunds": product.MinMarketFunds,
			"MaxMarketFunds": product.MaxMarketFunds,
			"MarginEnabled":  product.MarginEnabled,
			"PostOnly":       product.PostOnly,
			"LimitOnly":      product.LimitOnly,
			"CancelOnly":     product.CancelOnly,
			"Status":         product.Status,
			"StatusMessage":  product.StatusMessage,
		})
		if err != nil {
			return BadRequestError
		}
	}
	return nil
}

//get all products from db
func (d *DBDriver) GetProducts() ([]ProductsList, error) {
	data := []ProductsList{}
	err := d.Conn.Select(&data, `SELECT * FROM products`)

	if err != nil {
		return nil, DbQueryError
	}

	if err == sql.ErrNoRows {
		fmt.Println("no products")
		return nil, DbQueryError
	}
	return data, nil
}

//UpdateTicker func updates or inserts tickers into db
func (d *DBDriver) UpdateTicker(tickerData *TickerData) error {
	var test string
	var query string

	err := d.Conn.Get(&test, `SELECT id FROM products_tickers WHERE id=$1`, tickerData.ID)
	if err != nil {
		query = `INSERT INTO products_tickers (id, price, size, time,  bid, ask, volume) VALUES(:ID, :Price, :Size, :Time, :Bid, :Ask, :Volume)`
	} else {
		query = `UPDATE products_tickers SET price=:Price, size=:Size, time=:Time, bid=:Bid, ask=:Ask, volume=:Volume WHERE id=:ID`
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
		return DbQueryError
	}
	return nil
}

//SelectedProduct func retrieves product details from db based on product_id
func (d *DBDriver) SelectedProduct(id string) (*TickerData, error) {
	data := &TickerData{}
	err := d.Conn.Get(data, `SELECT id,price,size,time,bid,ask,volume FROM products_tickers WHERE id = $1`, id)

	if err != nil {
		return nil, DbQueryError
	}

	if err == sql.ErrNoRows {
		return nil, DbQueryError
	}
	return data, nil
}

//AddFav func adds product into user_favourites in db
func (d *DBDriver) AddFav(userId string, productID string) error {
	favId, err := uuid.NewV4()

	if err != nil {
		return AddFavProductError
	}

	fav_id := favId.String()

	query := `INSERT INTO users_favourites (fav_id, user_id, ticker_id) VALUES (:FavID, :UserID, :ProductID)`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"FavID":     fav_id,
		"UserID":    userId,
		"ProductID": productID,
	})

	if err != nil {
		return AddFavProductError
	}

	return nil
}

//RemoveFav func removes product from users_favourites in db
func (d *DBDriver) RemoveFav(userId string, productID string) error {
	query := `	UPDATE users_favourites 
				SET archived = true, archived_at = current_timestamp
				WHERE user_id = :UserID 
				AND ticker_id = :TickerID`
	_, err := d.Conn.NamedExec(query, map[string]interface{}{
		"UserID":   userId,
		"TickerID": productID,
	})

	if err != nil {
		return RemoveFavProductError
	}
	return nil
}

//CheckFav func checks whether product has been favourited by user in users_favourtites in db
func (d *DBDriver) CheckFav(userID string, productID string) (*[]FavProducts, error) {
	var favID string
	var isArchived bool

	err := d.Conn.Get(&favID, `SELECT fav_id FROM users_favourites WHERE user_id=$1 AND ticker_id=$2`, userID, productID)

	if err != nil {
		if err == sql.ErrNoRows {
			d.AddFav(userID, productID)
		} else {
			return nil, BadRequestError
		}
	}

	err = d.Conn.Get(&isArchived, `SELECT archived FROM users_favourites WHERE user_id=$1 AND ticker_id=$2`, userID, productID)
	if err != nil {
		return nil, BadRequestError
	} else if isArchived == true {
		d.UnarchivedProduct(userID, productID)
	} else {
		err = d.RemoveFav(userID, productID)
		if err != nil {
			return nil, RemoveFavProductError
		}
	}

	data, err := d.GetFavProducts(userID)
	if err != nil {
		fmt.Println(err)
		return nil, DbQueryError
	}
	return data, nil
}

//UnarchivedProduct func unarchives products in users_favourites table in db
func (d *DBDriver) UnarchivedProduct(userID string, productID string) error {
	query := `	UPDATE users_favourites 
	SET archived = false, archived_at = current_timestamp
	WHERE user_id = :UserID 
	AND ticker_id = :TickerID`
	_, err := d.Conn.NamedExec(query, map[string]interface{}{
		"UserID":   userID,
		"TickerID": productID,
	})

	if err != nil {
		return DbQueryError
	}
	return nil
}

//GetFavProducts func retrieves all user's favourite products in db
func (d *DBDriver) GetFavProducts(id string) (*[]FavProducts, error) {
	data := &[]FavProducts{}

	err := d.Conn.Select(data, `
	SELECT p.id, p.price, p.size, p.time, p.bid, p.ask, p.volume 
	FROM users_favourites uf, products_tickers p 
	WHERE uf.ticker_id = p.id
	AND uf.user_id=$1
	AND uf.archived = false`, id)

	if err == sql.ErrNoRows {
		return nil, EmptyFavProductList
	}

	if err != nil {
		return nil, DbQueryError
	}
	return data, nil
}

//VerifyUserAcc funs sets user's verification bool to true
func (d *DBDriver) VerifyUserAcc(veriToken string) error {
	query := `UPDATE users SET verification = true WHERE verification_token = :VerificationToken`
	_, err := d.Conn.NamedExec(query, map[string]interface{}{
		"VerificationToken": veriToken,
	})

	if err != nil {
		return UserVerificationError
	}
	return nil
}

//ResetPassword func resets user's password
func (d *DBDriver) ResetPassword(userID string, currentPw string, newPw string) error {
	var pw []byte

	err := d.Conn.Get(&pw, `SELECT password_hash FROM users WHERE id=$1`, userID)
	if err != nil {
		fmt.Println(err)
		return DbQueryError
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
		return ResetPasswordError
	}
	return nil
}

//GetResetPassToken func handles password reset on With  Reset Password Token
func (d *DBDriver) GetResetPassToken(email string) (string, error) {
	var token string
	err := d.Conn.Get(&token, `SELECT reset_pass_token FROM users WHERE email=$1`, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", RequestResetPassTokenError
		}
		return "", DbQueryError
	}
	return token, nil
}

//UpdatePassToken func resets Reset Password Token once user has reset password
func (d *DBDriver) UpdatePassWToken(resetPassToken string, password string) error {
	var userID string

	//gets user_id
	err := d.Conn.Get(&userID, `SELECT id FROM users WHERE reset_pass_token=$1`, resetPassToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return DbQueryError
		}
		return BadRequestError
	}

	//hashes new password
	newPwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ResetPasswordError
	}

	//generates new reset password token
	newToken, err := uuid.NewV4()
	if err != nil {
		return ResetPasswordError
	}

	new_token := newToken.String()

	//updates new password hash and new reset password token
	query := `UPDATE users SET password_hash = :NewPassword, reset_pass_token = :NewToken WHERE id = :UserID`
	_, err = d.Conn.NamedExec(query, map[string]interface{}{
		"NewPassword": newPwHash,
		"NewToken":    new_token,
		"UserID":      userID,
	})
	if err != nil {
		return ResetPasswordError
	}
	return nil
}
