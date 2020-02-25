
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table points (
    id serial primary key,
    point integer,
    user_id serial references users(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table points;
