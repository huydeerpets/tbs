
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_masters` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL COMMENT 'UserName',
    `email` VARCHAR(100) NOT NULL COMMENT 'Email',
    `password` text NOT NULL COMMENT 'Password',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_masters`;
