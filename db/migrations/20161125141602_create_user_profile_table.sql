
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_profiles (
  	id serial PRIMARY KEY,
    user_id serial REFERENCES users (id),
    name varchar,
    nick_name varchar,
    avatar varchar,
    gender int,
    age int,
    description varchar,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_profiles;
