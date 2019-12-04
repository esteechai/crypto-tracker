CREATE TABLE product_ticker (
	ticker_id TEXT PRIMARY KEY NOT NULL,
	price TEXT NOT NULL,
	size TEXT NOT NULL,
	time TEXT NOT NULL,
	bid TEXT NOT NULL,
	ask TEXT NOT NULL,
	volume TEXT NOT NULL
);

CREATE TABLE users (
	id TEXT PRIMARY KEY NOT NULL,
	email TEXT NOT NULL,
	username TEXT NOT NULL,
	password_hash bytea NOT NULL,
	verification bool NOT NULL,
	verification_token TEXT NOT NULL,
	reset_pass_token TEXT NOT NULL,
	CONSTRAINT user_un_email UNIQUE (email),
	CONSTRAINT user_un_username UNIQUE (username)
);