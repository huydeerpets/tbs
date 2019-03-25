
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_work_histories` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT 'User ID',
    `work_id` BIGINT NOT NULL COMMENT 'WorkID',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_work_histories`;
