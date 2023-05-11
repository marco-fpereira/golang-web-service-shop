BEGIN;

CREATE TABLE products (
    id serial PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    description varchar NOT NULL,
    price decimal NOT NULL,
    amount integer NOT NULL
);

INSERT INTO products (name, description, price, amount) VALUES
('Shirt', 'White oversized t-shirt', 29.9, 11),
('Notebook', 'Acer Aspire 3', 29.9, 11),
('Sneaker', 'Nike Air force 1', 499.99, 7),
('Headphone', 'Edifier G2', 199.90, 15);
/*
COPY products (name, description, price, amount) FROM stdin;
Shirt   White oversized t-shirt 29.9    11
Notebook    Acer Aspire 3   29.9    11
Sneaker Nike Air force 1    499.99  7
Headphone   Edifier G2  199.90  15
\.
*/

COMMIT;

ANALYZE products;

