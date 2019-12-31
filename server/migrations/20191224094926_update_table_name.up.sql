ALTER TABLE users 
RENAME COLUMN verification TO verified;


ALTER TABLE product_ticker
RENAME TO products_tickers; 


ALTER TABLE products_tickers
RENAME COLUMN ticker_id TO id; 


