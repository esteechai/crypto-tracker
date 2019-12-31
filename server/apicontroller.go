package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (api *API) setupHttp() chi.Router {
	go api.GetProductTicker()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

//AddCookie func adds cookie once user is logged in
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

//LogoutHandler func deletes cookie once user is logged out
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

//ReadCookie func reads and returns user's cookie
func (api *API) ReadCookie(w http.ResponseWriter, r *http.Request) {
	passback := &CookieSession{CheckedCookie: false}
	w.Header().Set("Content-Type", "application/json")

	_, err := r.Cookie("CryptoTracker")

	if err != nil {
		json.NewEncoder(w).Encode(passback)
		return
	}

	passback = &CookieSession{CheckedCookie: true}
	json.NewEncoder(w).Encode(passback)
}

//CheckCookie func checks user's cookie
func CheckCookie(next CookieHandler) http.HandlerFunc {
	function := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("CryptoTracker")

		if err != nil {
			fmt.Println("Check cookie error: ", err)
			return
		}
		userID := cookie.Value
		next(w, r, userID)
	}
	return function
}

//ConfirmEmailHandler func verifies user's email and set account status to verified
func (api *API) ConfirmEmailHandler(w http.ResponseWriter, r *http.Request) {
	veriToken := strings.ToLower(chi.URLParam(r, "veriToken"))
	if veriToken != "" {
		err := api.DB.VerifyUserAcc(veriToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	t, err := template.ParseFiles("template/verifiedEmail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

//LoginHandler func handles user login
func (api *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &SuccessWithID{ID: "", IsLoggedIn: false, ErrorMsg: ""}

	var loginDetails UserSignInData

	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		passback := &SuccessWithID{ID: "", IsLoggedIn: false, ErrorMsg: JSONParseError.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	id, err := api.DB.Login(loginDetails.Email, loginDetails.Password)
	if err != nil {
		passback = &SuccessWithID{ID: "", IsLoggedIn: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	passback = &SuccessWithID{ID: id, IsLoggedIn: true, ErrorMsg: ""}
	newW := AddCookie(w, "CryptoTracker", id)
	json.NewEncoder(newW).Encode(passback)
	return
}

//SignupHandler func handles user signup
func (api *API) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &SignupResult{IsSignup: false, ErrorMsg: "", IsVerified: false}

	var userSignUpData UserSignUpData
	err := json.NewDecoder(r.Body).Decode(&userSignUpData)

	if err != nil {
		passback = &SignupResult{IsSignup: false, ErrorMsg: JSONParseError.Error(), IsVerified: false}
		return
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

//ProductHandler func handles products list retrieved from coinbase API
func (api *API) ProductHandler(w http.ResponseWriter, r *http.Request) {
	passback := &ResponseResult{Success: false, ErrorMsg: ""}

	result, err := api.DB.GetProducts()
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

//GetProductsTicker func updates tickers every second in db
func (api *API) GetProductTicker() {
	var products []Products

	if products == nil {
		resp, err := http.Get("https://api.pro.coinbase.com/products")

		if err != nil {
			fmt.Println("Update Ticker Error: ", err)
			return
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = api.DB.AddProducts(products)
		if err != nil {
			fmt.Println("List Products Error: ", err)
			return
		}
	}

	n := 0
	c := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-c.C:
			err := api.FetchTicker(products[n].ID)
			if err != nil {
				fmt.Println("Update Ticker Error: ", err)
				continue
			}
			if n >= len(products)-1 {
				n = 0
			} else {
				n++
			}
		}
	}
}

//FetchTicker func gets product details
func (api *API) FetchTicker(id string) error {
	resp, err := http.Get("https://api.pro.coinbase.com/products/" + id + "/ticker")

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var productTicker *ProductTicker

	err = json.NewDecoder(resp.Body).Decode(&productTicker)

	if err != nil {
		return err
	}

	tickerData := &TickerData{ID: id, Price: productTicker.Price, Size: productTicker.Size, Time: productTicker.Time, Bid: productTicker.Bid, Ask: productTicker.Ask, Volume: productTicker.Volume}
	err = api.DB.UpdateTicker(tickerData)
	if err != nil {
		return err
	}
	return nil
}

type CoinID struct {
	ID string `json:"ticker_id"`
}

//TickerHandler func handles toggles favourites in db
func (api *API) TickerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id CoinID
	passback := &ResponseResult{Success: false, ErrorMsg: ""}

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: JSONParseError.Error()}
		json.NewEncoder(w).Encode(passback)
	}

	result, err := api.DB.SelectedProduct(id.ID)
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}

//FavouriteHandler func checks user's favourites in db
func (api *API) FavouriteHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResponseResult{Success: false, ErrorMsg: ""}

	var userFav UserFav
	err := json.NewDecoder(r.Body).Decode(&userFav)
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: JSONParseError.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	result, err := api.DB.CheckFav(userID, userFav.ProductID)
	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

//FavouriteListHandler func gets user's products list
func (api *API) FavouriteListHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResponseResult{Success: false, ErrorMsg: ""}

	result, err := api.DB.GetFavProducts(userID)

	if err != nil {
		passback = &ResponseResult{Success: false, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(passback)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

//ResetPasswordHandler func handles password reset without reset pass token
func (api *API) ResetPasswordHandler(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResponseResult{Success: true, ErrorMsg: ""}
	var resetPass ResetPass

	err := json.NewDecoder(r.Body).Decode(&resetPass)
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

//ForgotPasswordHandler func handles Forgot Password
func (api *API) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	passback := &ResetPassTokenResult{ResetPassToken: "", ErrorMsg: ""}
	var email ForgotPass

	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		passback = &ResetPassTokenResult{ResetPassToken: "", ErrorMsg: JSONParseError.Error()}
		json.NewEncoder(w).Encode(passback)
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

//ResetPassWTokenHandler func handles password reset with token
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

//UpdatePassWTokenHandler func validates new password input on forgot password
func (api *API) UpdatePassWTokenHandler(w http.ResponseWriter, r *http.Request) {
	resetPassToken := strings.ToLower(chi.URLParam(r, "resetPassToken"))
	newPassword := r.FormValue("newPw")
	confirmPassword := r.FormValue("confirmPw")
	t, err := template.ParseFiles("template/resetPasswordError.html")
	passback := ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: ""}

	if newPassword == "" || confirmPassword == "" {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Please enter both new and confirm password."}
		err = t.Execute(w, passback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	//checks whether both password input matches
	if newPassword != confirmPassword {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Your new password does not match with confirm password."}
		err = t.Execute(w, passback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	//checks password strength
	err = validatePassword(confirmPassword)
	if err != nil {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "Weak password."}
		err = t.Execute(w, passback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	//inserts into db
	err = api.DB.UpdatePassWToken(resetPassToken, confirmPassword)
	if err != nil {
		passback = ResetPassTokenResult{ResetPassToken: resetPassToken, ErrorMsg: "This link has expired."}
		err = t.Execute(w, passback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	//shows success
	t, err = template.ParseFiles("template/resetPasswordSuccess.html")
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
