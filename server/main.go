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
	EmailAPIKey := flag.String("email-api-key", "1f1faefbc512e382465fe0c5031334a1-f8b3d330-a01f2943", "api key for sending email")
	EmailDomain := flag.String("email-domain", "sandbox780aae9321c242b999909e946d7d9f5b.mailgun.org", "domain name for sending email")

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
