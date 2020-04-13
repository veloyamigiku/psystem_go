
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table point_operators (
    id serial primary key,
    name varchar(255) unique,
    password varchar(255)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table point_operators;
