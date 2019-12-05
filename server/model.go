package main

import "github.com/jmoiron/sqlx"

type DBDriver struct {
	Conn *sqlx.DB
}

type API struct {
	DB *DBDriver
}

//user input on login
type UserSignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//user input for signup
type UserSignUpData struct {
	Username string `json:"username"`
	Email    string `json: "email"`
	Password string `json:"password"`
}

//result for login
type SuccessWithID struct {
	ID         string `json:"id"`
	IsLoggedIn bool   `json:"is_login"`
	ErrorMsg   string `json:"error_msg"`
}

//user details retrieved from db after login
type SignInCreds struct {
	ID           string `db:"id"`
	PasswordHash []byte `db:"password_hash"`
}

type ResponseResult struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error_msg"`
}

//result for signup
type SignupResult struct {
	IsSignup bool   `json:"is_signup"`
	ErrorMsg string `json:"error_msg"`
}

const (
	UniqueConstraintUsername = "user_un_username"
	UniqueConstraintEmail    = "user_un_email"
)

type Products struct {
	ID             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
	BaseIncrement  string `json:"base_increment"`
	DisplayName    string `json:"display_name"`
	MinMarketFunds string `json:"min_market_funds"`
	MaxMarketFunds string `json:"max_maret_funds"`
	MarginEnabled  bool   `json:"margin_enabled"`
	PostOnly       bool   `json:"post_only"`
	LimitOnly      bool   `json:"limit_only"`
	CancelOnly     bool   `json:"cancel_only"`
	Status         string `json:"status"`
	StatusMessage  string `json:"status_message"`
}

type ProductTicker struct {
	TradeID string `json:"string,trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Time    string `json:"time"`
	Bid     string `json:"bid"`
	Ask     string `json:"ask"`
	Volume  string `json:"volume"`
}

type TickerData struct {
	ID     string `db:"ticker_id"`
	Price  string `db:"price"`
	Size   string `db:"size"`
	Time   string `db:"time"`
	Bid    string `db:"bid"`
	Ask    string `db:"ask"`
	Volume string `db:"volume"`
}

type SelectedTickerData struct {
	ID     string `"json:"id"`
	Price  string `json:"price"`
	Size   string `json:"size"`
	Time   string `json:"time"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Volume string `json:"volume"`
}

type UserFavDB struct {
	FavID    string `db:"fav_id"`
	UserID   string `db:"user_id"`
	TickerID string `db:"ticker_id"`
}

type UserFav struct {
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	isFav     bool   `json:"is_fav"`
}

type UserFavResult struct {
	QueryResult string `json:"result"`
	IsSuccess   bool   `json:"is_success"`
}