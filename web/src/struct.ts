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
    ID: string 
    BaseCurrency: string 
    QuoteCurrency: string 
    BaseMinSize: string 
    BaseMaxSize: string 
    QuoteIncrement: string 
    BaseIncrement: string 
    DisplayName: string 
    MinMarketFunds: string 
    MaxMarketFunds: string 
    MarginEnabled: boolean
    PostOnly: boolean
    LimitOnly: boolean 
    CancelOnly: boolean 
    Status: string 
    StatusMessage: string 
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
    product_id: string
}

export interface CurrentUserID {
    user_id: string 
}

export interface ResetPasword {
    current_password: string 
    new_password: string 
}

export interface ForgotPasword {
    email: string 
}




