-- +migrate Up
CREATE TABLE roles (
    id int NOT NULL PRIMARY KEY,
    name text
);

-- +migrate Down
DROP TABLE roles;
