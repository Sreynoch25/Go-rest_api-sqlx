-- Migration Up
CREATE TABLE IF NOT EXISTS tbl_users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    login_id VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_name VARCHAR(100),
    role_id INTEGER,
    is_admin BOOLEAN DEFAULT FALSE,
    login_session VARCHAR(255),
    last_login TIMESTAMP,
    currency_id INTEGER,
    language_id INTEGER,
    profile TEXT,
    parent_id INTEGER,
    level VARCHAR(50),
    status_id INTEGER NOT NULL DEFAULT 1,
    "order" INTEGER,
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    updated_at TIMESTAMP,
    deleted_by INTEGER,
    deleted_at TIMESTAMP,
);

