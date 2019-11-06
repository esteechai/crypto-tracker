package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)

func (api *API) setupHttp() chi.Router {
	r := chi.NewRouter()
	//r.Post("/login", TestHandler)
	r.Post("/login", api.LoginHandler)
	r.Post("/signup", api.SignupHandler)
	fmt.Println("Successfully connected!")
	return r
}

func (api *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginDetails userData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginDetails)
	if err != nil {
		fmt.Println(err)
	}
	passback, err := api.DB.Login(loginDetails.Email, loginDetails.Password)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passback)
	return
}

func (api *API) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupDetails userSignup
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signupDetails)
	if err != nil {
		fmt.Println(err)
	}
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	idstr := id.String()
	passback, err := api.DB.Signup(idstr, signupDetails.Username, signupDetails.Email, signupDetails.Password)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passback)
	return
}
