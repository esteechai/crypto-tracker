package main

import "github.com/jmoiron/sqlx"

type DBDriver struct {
	Conn *sqlx.DB
}

type API struct {
	DB *DBDriver
}

type UserSignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpData struct {
	Username string `json:"username`
	Email    string `json: "email`
	Password string `json:"password"`
}

type SuccessWithID struct {
	ID       string `json:"id"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error_msg"`
}

type SignInCreds struct {
	ID           string `db:"id"`
	PasswordHash []byte `db:"password_hash"`
}

type ResponseResult struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error_msg"`
}

type SignupResult struct {
	Success  bool   `json:"success"`
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
	TradeID string `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Time    string `json:"time"`
	Bid     string `json:"bid"`
	Ask     string `json:"ask"`
	Volume  string `json:"volume"`
}
