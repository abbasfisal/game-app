-- +migrate Up
-- in mysql 8 role will set to user for old users by default
ALTER TABLE `users`
    ADD COLUMN `role` ENUM('user','admin') NOT NULL;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;