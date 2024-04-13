CREATE TABLE user_ (
  id CHAR(36) PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  auth0_id CHAR(30)
);

INSERT INTO user_ (id, username, created_at, auth0_id) VALUES ('0630d891-d3b9-484e-8486-59c21bdf8f38', 'user1', '2024-04-08 13:02:00.000', 'auth0|661a0e9b2875554618dd9b91');

CREATE TABLE product_ (
  id CHAR(36) PRIMARY KEY,
  name_ VARCHAR(255) UNIQUE NOT NULL,
  price INT NOT NULL
);
