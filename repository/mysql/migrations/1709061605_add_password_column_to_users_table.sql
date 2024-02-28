-- +migrate Up
ALTER TABLE users
    ADD COLUMN password varchar(255) not null;

-- +migrate Down
ALTER TABLE users DROP COLUMN password;