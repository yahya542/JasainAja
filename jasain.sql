CREATE DATABASE JasainAja_db;

CREATE TABLE users (
    user_id  SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    time_balance INT DEFAULT 0
);

CREATE TABLE providers (
    provider_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL,
    skills TEXT,
    bio TEXT,
    rating FLOAT
);

CREATE TABLE REQUEST (
    request_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    provider_id INT,
    description TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (provider_id) REFERENCES providers(provider_id)
);

CREATE TABLE SERVICES (
    service_id SERIAL PRIMARY KEY,
    provider_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    duration_minutes INT NOT NULL,
    price_time INT NOT NULL
   
);

CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    request_id INT NOT NULL REFERENCES REQUEST(request_id) ON DELETE CASCADE,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP,
    status TEXT CHECK (status IN ('in_progress','completed','cancelled')) DEFAULT 'in_progress'
);

CREATE TABLE admin_reports (
    report_id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL REFERENCES transactions(transaction_id) ON DELETE CASCADE,
    report_type TEXT CHECK (report_type IN ('transaction_confirmed','completed','dispute')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status TEXT CHECK (status IN ('unread','read')) DEFAULT 'unread'
);