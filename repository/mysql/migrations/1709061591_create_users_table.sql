-- +migrate Up
create table users
(
    id           int primary key AUTO_INCREMENT,
    name         varchar(255) not null,
    phone_number varchar(255) not null unique,
    created_at   datetime default current_timestamp
);

-- +migrate Down
Drop Table users;