
CREATE TABLE users_favourites (
    fav_id TEXT PRIMARY KEY NOT NULL, 
    user_id TEXT NOT NULL,
    ticker_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(ticker_id) REFERENCES product_ticker(ticker_id),
	archived BOOLEAN NOT NULL DEFAULT FALSE,
	archived_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW() 
);
