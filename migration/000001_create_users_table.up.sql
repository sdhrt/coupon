CREATE TABLE IF NOT EXISTS users( 
    user_id CHAR(36) NOT NULL PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);
