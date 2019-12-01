package main

import (
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
	api := &API{DB: DBDriver}

	r := api.setupHttp()
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
