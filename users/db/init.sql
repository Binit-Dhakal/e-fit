CREATE USER users_service WITH PASSWORD 'users_pass';

CREATE DATABASE users_db OWNER users_service;

GRANT ALL privileges ON DATABASE users_db TO users_service;

