package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var products []*Products
var productTicker *ProductTicker

func (api *API) setupHttp() chi.Router {
	GetProducts()
	go api.GetProductTicker()
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Post("/api/login", api.LoginHandler)
	r.Post("/api/signup", api.SignupHandler)
	r.Get("/api/get/products", api.ProductHandler)
	r.Post("/api/ticker", api.TickerHandler)
	r.Post("/api/favToggle", api.FavouriteHandler)
	fmt.Println("Successfully connected!")
	return r
}

//handle login validation
func (api *API) LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	passback := &SuccessWithID{ID: "", IsLoggedIn: true, ErrorMsg: ""}

	var loginDetails UserSignInData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginDetails)
	if err != nil {
		fmt.Println(err)
	}

	id, err := api.DB.Login(loginDetails.Email, loginDetails.Password)
	if err != nil {
		passback = &SuccessWithID{ID: "", IsLoggedIn: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	passback = &SuccessWithID{ID: id, IsLoggedIn: true, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)
	return
}

//handle signup validation
func (api *API) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &SignupResult{IsSignup: true, ErrorMsg: ""}

	var userSignUpData UserSignUpData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userSignUpData)

	if err != nil {
		fmt.Println(err)
	}

	err = validateEmail(userSignUpData.Email)
	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	err = validatePassword(userSignUpData.Password)
	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	result, err := api.DB.Signup(userSignUpData.Username, userSignUpData.Email, userSignUpData.Password)
	if err != nil {
		result = false
		if IsUniqueConstraintError(err, UniqueConstraintUsername) {
			passback = &SignupResult{IsSignup: result, ErrorMsg: err.Error()}
		}
		if IsUniqueConstraintError(err, UniqueConstraintEmail) {
			passback = &SignupResult{IsSignup: result, ErrorMsg: err.Error()}
		}
		return
	}

	passback = &SignupResult{IsSignup: true, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)

	fmt.Println("passback: ", passback)

	return
}

//handle products list retrieved from coinbase API
func (api *API) ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

//retrieve products from coinbase api
func GetProducts() []*Products {
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

func (api *API) GetProductTicker() {
	if products == nil {
		GetProducts()
	}

	n := 0
	for {
		go api.FetchTicker(products[n].ID)

		//fmt.Println(products[n].ID)

		if n >= len(products)-1 {
			n = 0
		} else {
			n++
		}
		time.Sleep(1 * time.Second)
	}
}

// get product details
func (api *API) FetchTicker(id string) {
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

	tickerData := &TickerData{ID: id, Price: productTicker.Price, Size: productTicker.Size, Time: productTicker.Time, Bid: productTicker.Bid, Ask: productTicker.Ask, Volume: productTicker.Volume}
	api.DB.UpdateTicker(tickerData)
}

type CoinID struct {
	ID string `json:"ticker_id"`
}

func (api *API) TickerHandler(w http.ResponseWriter, r *http.Request) {
	var id CoinID
	fmt.Println("api controller: ", id)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	result, err := api.DB.SelectedProduct(id.ID)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (api *API) FavouriteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &UserFavResult{QueryResult: "", IsSuccess: false}

	var userFav *UserFav
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userFav)
	if err != nil {
		fmt.Println(err)
	}

	if userFav.isFav {
		err := api.DB.RemoveFav(userFav.UserID, userFav.ProductID)
		if err != nil {
			fmt.Println(err)
			passback = &UserFavResult{QueryResult: err.Error(), IsSuccess: false}
			return
		}
	} else {
		err := api.DB.AddFav(userFav.UserID, userFav.ProductID)
		if err != nil {
			fmt.Println(err)
			passback = &UserFavResult{QueryResult: err.Error(), IsSuccess: false}
			return
		}
	}
	passback = &UserFavResult{QueryResult: "", IsSuccess: true}
	json.NewEncoder(w).Encode(passback)
	return
}
