package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// server port
	ServerPort := flag.String("server-port", ":8084", "server port")

	//database connection
	DBHost := flag.String("db-host", "localhost", "db host")
	DBPort := flag.Int("db-port", 5432, "db port")
	DBUser := flag.String("db-user", "esteecoinbase", "db user")
	DBPassword := flag.String("db-pw", "dev", "db password")
	DBName := flag.String("db-name", "esteecoinbase", "db name")

	//email
	Scheme := flag.String("scheme", "http", "scheme for http or https")
	ServerDomain := flag.String("server-domain", "cryptotracker.interns.theninja.life", "server domain name")
	EmailAPIKey := flag.String("email-api-key", "1f1faefbc512e382465fe0c5031334a1-f8b3d330-a01f2943", "api key for sending email")
	EmailDomain := flag.String("email-domain", "sandbox780aae9321c242b999909e946d7d9f5b.mailgun.org", "domain name for sending email")

	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		*DBHost, *DBPort, *DBUser, *DBPassword, *DBName)

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
	err = http.ListenAndServe(*ServerPort, r)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
