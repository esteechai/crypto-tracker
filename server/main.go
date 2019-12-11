package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 6432
	user     = "coinbase"
	password = "dev"
	dbname   = "coinbase"
)

func main() {
	Scheme := flag.String("scheme", "http", "scheme for http or https")
	ServerDomain := flag.String("server-domain", "localhost:8080", "server domain name")
	EmailAPIKey := flag.String("email-api-key", "9afc1023d18b6f5ffd4f9a6da6a7cce5-5645b1f9-7004f18a", "api key for sending email")
	EmailDomain := flag.String("email-domain", "sandbox46b550a5b410473e985b1c64e7525ea1.mailgun.org", "domain name for sending email")

	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	DBDriver := &DBDriver{Conn: conn}
	emailDomain := &SendEmailInfo{EmailAPIKey: *EmailAPIKey, ServerDomain: *ServerDomain, EmailDomain: *EmailDomain, Scheme: *Scheme}
	api := &API{DB: DBDriver, EmailInfo: emailDomain}

	r := api.setupHttp()
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
