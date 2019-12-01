package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var products []*Products
var productTicker *ProductTicker

func (api *API) setupHttp() chi.Router {
	getProducts()
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Post("/api/login", api.LoginHandler)
	r.Post("/api/signup", api.SignupHandler)
	r.Get("/api/get/products", api.ProductHandler)
	r.Get("api/get/ticker", api.TickerHandler)
	fmt.Println("Successfully connected!")
	return r
}

//handle login validation
func (api *API) LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	passback := &SuccessWithID{ID: "", Success: true, ErrorMsg: ""}

	var loginDetails UserSignInData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginDetails)
	if err != nil {
		fmt.Println(err)
	}

	id, err := api.DB.Login(loginDetails.Email, loginDetails.Password)
	if err != nil {
		passback = &SuccessWithID{ID: "", Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	passback = &SuccessWithID{ID: id, Success: true, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)
	return
}

//handle signup validation
func (api *API) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &SignupResult{Success: true, ErrorMsg: ""}

	var userSignUpData UserSignUpData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userSignUpData)

	if err != nil {
		fmt.Println(err)
	}

	err = validateEmail(userSignUpData.Email)
	if err != nil {
		passback = &SignupResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	err = validatePassword(userSignUpData.Password)
	if err != nil {
		passback = &SignupResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	result, err := api.DB.Signup(userSignUpData.Username, userSignUpData.Email, userSignUpData.Password)
	if err != nil {
		result = false
		if IsUniqueConstraintError(err, UniqueConstraintUsername) {
			passback = &SignupResult{Success: result, ErrorMsg: err.Error()}
		}
		if IsUniqueConstraintError(err, UniqueConstraintEmail) {
			passback = &SignupResult{Success: result, ErrorMsg: err.Error()}
		}
		return
	}

	passback = &SignupResult{Success: true, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)
	return
}

//handle products list retrieved from coinbase API
func (api *API) ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

//retrieve products from coinbase api
func getProducts() []*Products {
	resp, err := http.Get("https://api.pro.coinbase.com/products")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&products)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return products
}

func (api *API) TickerHandler(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://api.pro.coinbase.com/products/" + id + "/ticker")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	var productTicker *ProductTicker

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&productTicker)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(productTicker)
	return
}
