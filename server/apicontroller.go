package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"
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
	r.Get("/api/logout", api.LogoutHandler)
	r.Get("/api/get/products", api.ProductHandler)               //list all products
	r.Post("/api/ticker", api.TickerHandler)                     //product details
	r.Post("/api/fav-toggle", CheckCookie(api.FavouriteHandler)) //fav and unfav products when fav icon is clicked

	r.Get("/api/fav-list", CheckCookie(api.FavouriteListHandler)) //list and update fav products

	r.Get("/api/confirm-email/{veriToken}", api.ConfirmEmailHandler)             //verify email after signup
	r.Post("/api/reset-password", CheckCookie(api.ResetPasswordHandler))         //handle reset password
	r.Post("/api/forgot-password", api.ForgotPasswordHandler)                    //handle forgot password and send email
	r.Get("/api/reset-password/{resetPassToken}", api.ResetPassWTokenHandler)    //handle reset password with reset pass token
	r.Post("/api/update-password/{resetPassToken}", api.UpdatePassWTokenHandler) //update password with reset pass token

	r.Get("/api/auth", api.ReadCookie) //fetch cookie once login
	fmt.Println("Successfully connected!")
	return r
}

//add cookie
func AddCookie(w http.ResponseWriter, name string, value string) http.ResponseWriter {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return w
}

//delete cookie
func (api *API) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie := http.Cookie{
		Name:     "CryptoTracker",
		MaxAge:   -1,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	passback := &ResponseResult{Success: true, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)
}

// read cookie after login
func (api *API) ReadCookie(w http.ResponseWriter, r *http.Request) {
	passback := &CookieSession{CheckedCookie: false}
	w.Header().Set("Content-Type", "application/json")

	_, err := r.Cookie("CryptoTracker")
	if err != nil {
		fmt.Println("read cookie in api controller:", err)
		json.NewEncoder(w).Encode(passback)
		return
	}
	passback = &CookieSession{CheckedCookie: true}
	json.NewEncoder(w).Encode(passback)

}

//check cookie when API is called
func CheckCookie(next CookieHandler) http.HandlerFunc {
	// passback := &SuccessWithID{ID: "", IsLoggedIn: true, ErrorMsg: ""}
	function := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("CryptoTracker")
		if err != nil {
			fmt.Println("No cookie here: ", err)
			return
		}
		userID := cookie.Value
		next(w, r, userID)
		fmt.Println("check cookie:", userID)
	}
	return function
}

//handle signup verification token
func (api *API) ConfirmEmailHandler(w http.ResponseWriter, r *http.Request) {
	veriToken := strings.ToLower(chi.URLParam(r, "veriToken"))
	if veriToken != "" {
		err := api.DB.VerifyUserAcc(veriToken)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("email verification successful")

	t, err := template.ParseFiles("template/verifiedEmail.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
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

	newW := AddCookie(w, "CryptoTracker", id)
	fmt.Println("cookie added:", newW)

	json.NewEncoder(newW).Encode(passback)
	return
}

//handle signup validation
func (api *API) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &SignupResult{IsSignup: false, ErrorMsg: "", IsVerified: false}

	var userSignUpData UserSignUpData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userSignUpData)

	if err != nil {
		fmt.Println(err)
	}

	err = validateEmail(userSignUpData.Email)
	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: err.Error(), IsVerified: false}
		json.NewEncoder(w).Encode(passback)
		return
	}
	err = validatePassword(userSignUpData.Password)
	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: err.Error(), IsVerified: false}
		json.NewEncoder(w).Encode(passback)
		return
	}

	err = api.DB.Signup(userSignUpData.Username, userSignUpData.Email, userSignUpData.Password)
	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: err.Error(), IsVerified: false}
		json.NewEncoder(w).Encode(passback)
		return
	}

	// verifToken, err := api.DB.GetVerifToken(userSignUpData.Username, userSignUpData.Email)
	// if err != nil {
	// 	passback = &SignupResult{IsSignup: result, ErrorMsg: err.Error(), IsVerified: false}
	// 	json.NewEncoder(w).Encode(passback)
	// 	return
	// }

	// err = api.EmailInfo.VerifyEmail(userSignUpData.Email, verifToken)
	// if err != nil {
	// 	fmt.Println(err)
	// 	passback = &SignupResult{IsSignup: result, ErrorMsg: err.Error(), IsVerified: false}
	// 	json.NewEncoder(w).Encode(passback)
	// 	return
	// }
	passback = &SignupResult{IsSignup: true, ErrorMsg: "", IsVerified: false}
	json.NewEncoder(w).Encode(passback)
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

func (api *API) FavouriteHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")

	var userFav UserFav
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userFav)
	if err != nil {
		fmt.Println(err)
	}
	// result, err := api.DB.CheckFav(userFav.UserID, userFav.ProductID)
	result, err := api.DB.CheckFav(userID, userFav.ProductID)
	json.NewEncoder(w).Encode(result)
	return
}

// type UserID struct {
// 	ID string `json:"user_id"`
// }

func (api *API) FavouriteListHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("current user id:", userID)
	// var id UserID
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&id)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// result, err := api.DB.GetFavProducts(id.ID)
	fmt.Println("api controller: userid ", userID)

	result, err := api.DB.GetFavProducts(userID)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}

//handle reset password
func (api *API) ResetPasswordHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResponseResult{Success: true, ErrorMsg: ""}
	var resetPass ResetPass
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resetPass)
	if err != nil {
		fmt.Println(err)
		passback = &ResponseResult{Success: false, ErrorMsg: BadRequestError.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	if resetPass.CurrentPw == resetPass.NewPw {
		passback = &ResponseResult{Success: false, ErrorMsg: SameResetPwInput.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	err = validatePassword(resetPass.NewPw)
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: WeakPassword.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	err = api.DB.ResetPassword(userID, resetPass.CurrentPw, resetPass.NewPw)
	if err != nil {
		passback := &ResponseResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	json.NewEncoder(w).Encode(passback)
	return
}

func (api *API) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResetPassTokenResult{ResetPassToken: "", ErrorMsg: ""}
	var email ForgotPass

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&email)

	if err != nil {
		fmt.Println(err)
		return
	}

	token, err := api.DB.GetResetPassToken(email.Email)
	if err != nil {
		passback = &ResetPassTokenResult{ResetPassToken: "", ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	err = api.EmailInfo.ResetPassword(email.Email, token)
	if err != nil {
		fmt.Println(err)
		passback = &ResetPassTokenResult{ResetPassToken: "", ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	passback = &ResetPassTokenResult{ResetPassToken: token, ErrorMsg: ""}
	json.NewEncoder(w).Encode(passback)
	return
}

//handle password reset with token
func (api *API) ResetPassWTokenHandler(w http.ResponseWriter, r *http.Request) {
	resetPassToken := strings.ToLower(chi.URLParam(r, "resetPassToken"))
	passback := &ForgotPassReset{ResetPassToken: resetPassToken, NewPassword: ""}
	t, err := template.ParseFiles("template/resetPassword.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, passback)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//validate password input on forgot password
func (api *API) UpdatePassWTokenHandler(w http.ResponseWriter, r *http.Request) {
	resetPassToken := strings.ToLower(chi.URLParam(r, "resetPassToken"))
	newPassword := r.FormValue("newPw")
	confirmPassword := r.FormValue("confirmPw")
	t, err := template.ParseFiles("template/resetPasswordError.html")
	passback := ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: ""}
	// fmt.Println("forgot pw: ", newPassword, confirmPassword)

	if newPassword == "" || confirmPassword == "" {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Please enter both new and confirm password."}
		err = t.Execute(w, passback)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	//check whether both password input matches
	if newPassword != confirmPassword {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Your new password does not match with confirm password."}
		err = t.Execute(w, passback)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	//check password strength
	err = validatePassword(confirmPassword)
	if err != nil {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Weak password."}
		err = t.Execute(w, passback)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	//insert into db
	err = api.DB.UpdatePassWToken(resetPassToken, confirmPassword)
	if err != nil {
		// fmt.Println(err)
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "This link has expired."}
		err = t.Execute(w, passback)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	//show success
	t, err = template.ParseFiles("template/resetPasswordSuccess.html")
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
