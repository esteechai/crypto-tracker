package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type DBDriver struct {
	Conn *sqlx.DB
}

type API struct {
	DB        *DBDriver
	EmailInfo *SendEmailInfo
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
	IsSignup   bool   `json:"is_signup"`
	ErrorMsg   string `json:"error_msg"`
	IsVerified bool   `json:"is_verified"`
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
	ProductID string `json:"product_id"`
	IsFav     bool   `json:"is_fav"`
}

type FavProducts struct {
	ID     string `db:"ticker_id"`
	Price  string `db:"price"`
	Size   string `db:"size"`
	Time   string `db:"time"`
	Bid    string `db:"bid"`
	Ask    string `db:"ask"`
	Volume string `db:"volume"`
}

type SendEmailInfo struct {
	EmailAPIKey  string
	EmailDomain  string
	Scheme       string
	ServerDomain string
}

type ResetPass struct {
	CurrentPw string `json:"current_password"`
	NewPw     string `json:"new_password"`
}

type ForgotPass struct {
	Email string `json:"email"`
}

type ResetPassTokenResult struct {
	ResetPassToken string `db:"reset_pass_token"`
	ErrorMsg       string `json:"error_msg"`
}

type ForgotPassReset struct {
	ResetPassToken string
	NewPassword    string
}

//check cookie
type CookieHandler func(w http.ResponseWriter, r *http.Request, userID string)

type CookieSession struct {
	CheckedCookie bool `json:"checked_cookie"`
}
