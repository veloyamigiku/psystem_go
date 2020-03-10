
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists users (
    id serial primary key,
    name varchar(255) unique,
    username varchar(255),
    password varchar(255)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
