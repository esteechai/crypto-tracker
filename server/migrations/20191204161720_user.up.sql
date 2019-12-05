CREATE TABLE users (
	id TEXT PRIMARY KEY NOT NULL,
	email TEXT NOT NULL,
	username TEXT NOT NULL,
	password_hash bytea NOT NULL,
	verification boolean NOT NULL,
	verification_token TEXT NOT NULL,
	reset_pass_token TEXT NOT NULL,
	CONSTRAINT user_un_email UNIQUE (email),
	CONSTRAINT user_un_username UNIQUE (username)
);

CREATE TABLE users_favourites (
    fav_id TEXT PRIMARY KEY NOT NULL, 
    user_id TEXT NOT NULL,
    ticker_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(ticker_id) REFERENCES product_ticker(ticker_id)
);
