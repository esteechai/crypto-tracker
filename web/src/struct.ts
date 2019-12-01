export interface authenticateData {
    email: string 
    password: string 
    isSubmit: boolean 
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

export interface coinbaseTicker {
    trade_id: string 
	price:   string 
	size:  string 
	time:    string
	ask:   string 
	volume:  string 
}