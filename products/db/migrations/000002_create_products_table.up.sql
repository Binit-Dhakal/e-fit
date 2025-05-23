CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name text,
    slug text UNIQUE NOT NULL,
    category text,
    images text[],
    brand text,
    description text,
    stock integer,
    price numeric(12, 2),
    rating numeric(3, 2),
    num_reviews integer DEFAULT 0,
    is_featured boolean DEFAULT FALSE,
    created_at timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz(6) DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_products_updated_at ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER products_update_timestamp
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_products_updated_at ();

