CREATE USER product_user WITH PASSWORD 'product_pass';

CREATE DATABASE product_db OWNER product_user;

GRANT ALL privileges ON DATABASE product_db TO product_user;

