-- Create users table
CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    created_by VARCHAR,
    updated_by VARCHAR,
    user_id_core VARCHAR NOT NULL,
    username VARCHAR NOT NULL UNIQUE,
    citizen_id VARCHAR NOT NULL,
    citizen_name VARCHAR NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    email VARCHAR NOT NULL,
    public_key TEXT NOT NULL,
    aptos_address TEXT NOT NULL,
    encrypted_hash TEXT NOT NULL,
    nonce TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create candidates table
CREATE TABLE candidates (
    id VARCHAR(50) PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    created_by VARCHAR,
    updated_by VARCHAR,
    candidate_name VARCHAR NOT NULL,
    citizen_id VARCHAR NOT NULL,
    avatar_url TEXT,
    server_id VARCHAR NOT NULL,
    candidate_index BIGINT NOT NULL
);

-- Create voting_servers table
CREATE TABLE voting_servers (
    id VARCHAR(50) PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    created_by VARCHAR,
    updated_by VARCHAR,
    admin_id VARCHAR NOT NULL,
    number_of_candidates BIGINT NOT NULL,
    maximum_number_of_voters BIGINT NOT NULL,
    server_name VARCHAR NOT NULL,
    server_id VARCHAR NOT NULL UNIQUE,
    opened_vote BOOLEAN NOT NULL DEFAULT FALSE,
    results TEXT,
    contract_address TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    exp_time BIGINT NOT NULL
);

-- Create configs table
CREATE TABLE configs (
    key VARCHAR PRIMARY KEY,
    value TEXT NOT NULL
);