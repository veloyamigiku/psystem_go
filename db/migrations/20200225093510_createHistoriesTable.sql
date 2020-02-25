
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table point_histories (
    id serial primary key,
    user_id serial references users(id),
    date timestamp with time zone,
    detail text,
    point integer
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table point_histories;
