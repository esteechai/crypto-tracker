CREATE TABLE product_ticker (
	ticker_id TEXT PRIMARY KEY NOT NULL,
	price TEXT NOT NULL,
	size TEXT NOT NULL,
	time TEXT NOT NULL,
	bid TEXT NOT NULL,
	ask TEXT NOT NULL,
	volume TEXT NOT NULL
);