
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE UNIQUE INDEX user_profiles_user_id_unique ON user_profiles (user_id);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX user_profiles_user_id_unique;
