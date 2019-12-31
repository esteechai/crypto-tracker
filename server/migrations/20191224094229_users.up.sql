CREATE TABLE users (
	id TEXT PRIMARY KEY NOT NULL,
	email TEXT NOT NULL,
	username TEXT NOT NULL,
	password_hash bytea NOT NULL,
	verification boolean NOT NULL,
	verification_token TEXT NOT NULL,
	reset_pass_token TEXT NOT NULL,
	CONSTRAINT user_un_email UNIQUE (email),
	CONSTRAINT user_un_username UNIQUE (username),
	archived BOOLEAN NOT NULL DEFAULT FALSE,
	archived_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW() 
);
