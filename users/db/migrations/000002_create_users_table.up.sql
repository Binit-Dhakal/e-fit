CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    username text,
    password text,
    email text UNIQUE,
    email_verified boolean DEFAULT FALSE,
    image text,
    is_guest boolean DEFAULT TRUE,
    createdAt timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamptz(6) DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_users_updated_at ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updatedAt := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER users_update_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_users_updated_at ();

