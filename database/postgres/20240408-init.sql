CREATE TABLE product_ (
  id SERIAL PRIMARY KEY,
  name_ VARCHAR(255) NOT NULL,
  price INT UNIQUE NOT NULL
);

CREATE INDEX user__phone_idx ON user_ (phone);