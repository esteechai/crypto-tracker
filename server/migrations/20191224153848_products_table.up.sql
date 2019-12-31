CREATE TABLE products (
	id TEXT PRIMARY KEY NOT NULL,
	base_currency TEXT NOT NULL,
	quote_currency TEXT NOT NULL,
	base_min_size TEXT NOT NULL,
	base_max_size TEXT NOT NULL,
	quote_increment TEXT NOT NULL,
	base_increment TEXT NOT NULL,
    display_name TEXT NOT NULL, 
    min_market_funds TEXT NOT NULL, 
    max_market_funds TEXT NOT NULL,
    margin_enabled BOOLEAN NOT NULL,
    post_only BOOLEAN NOT NULL, 
    limit_only BOOLEAN NOT NULL, 
    cancel_only BOOLEAN NOT NULL, 
    status TEXT NOT NULL, 
    status_message TEXT NOT NULL
);