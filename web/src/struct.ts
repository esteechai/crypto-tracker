export interface AuthenticateUserSignin {
    email: string 
    password: string 
    isSubmit: boolean 
}

export interface AuthenticateUserSignup {
    username: string 
    email: string 
    password: string 
}

export interface coinbaseProducts {
    id: string 
    base_currency: string 
    quote_currency: string 
    base_min_size: string 
    base_max_size: string 
    quote_increment: string 
    base_increment: string 
    display_name: string 
    min_market_funds: string 
    max_market_funds: string 
    margin_enabled: boolean
    post_only: boolean
    limit_only: boolean 
    cancel_only: boolean 
    status: string 
    status_message: string 
}

export interface selectedProductID {
    ticker_id: string 
}

export interface coinbaseTicker {
    ID: string 
	Price:   string 
	Size:  string 
	Time:    string
	Ask:   string 
	Volume:  string 
}

export interface UserFavList {
    ID: string 
	Price:   string 
	Size:  string 
    Time:    string
    Bid: string 
	Ask:   string 
	Volume:  string 
}

export interface FavToggle {
    user_id: string 
    product_id: string
}

export interface CurrentUserID {
    user_id: string 
}

export interface ResetPasword {
    user_id: string 
    current_password: string 
    new_password: string 
}

export interface ForgotPasword {
    email: string 
}




