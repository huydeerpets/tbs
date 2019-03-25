
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `log_contribution_images` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'PostID',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `log_contribution_images`;
