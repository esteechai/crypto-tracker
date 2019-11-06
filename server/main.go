package main

import (
	"encoding/json"
	"net/http"

	"fmt"

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
	http.ListenAndServe(":3000", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var loginDetails userData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginDetails)
	if err != nil {
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginDetails)
}

// func (db *DBDriver) LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var loginDetails userData
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&loginDetails)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	passback, err := db.login(loginDetails.ID)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(passback)
// 	return
// }

// func (db *DBDriver) SignupHandler(w http.ResponseWriter, r *http.Request) {
// 	var signupDetails userSignup
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&signupDetails)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	id, err := uuid.NewV4()
// 	if err != nil {
// 		log.Fatalf("failed to generate UUID: %v", err)
// 	}

// 	idstr := id.String()
// 	var passback []userSignup
// 	err = db.Conn.Select(&passback, `INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)`, idstr, signupDetails.Username, signupDetails.Email, signupDetails.Password)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(passback)
// 	return
// }
