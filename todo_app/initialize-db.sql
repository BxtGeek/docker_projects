-- Create database
CREATE DATABASE IF NOT EXISTS tododb;

-- Use the database
USE tododb;

-- Create tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed BOOLEAN DEFAULT FALSE
);

-- Create a user with proper permissions
CREATE USER IF NOT EXISTS 'todouser'@'%' IDENTIFIED BY 'todopassword';
GRANT ALL PRIVILEGES ON tododb.* TO 'todouser'@'%';
FLUSH PRIVILEGES;
