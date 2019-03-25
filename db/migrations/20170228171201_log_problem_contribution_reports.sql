
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `log_problem_contribution_reports` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'PostID',
    `user_id` BIGINT NOT NULL COMMENT 'User ID',
    `type` INT COMMENT 'Type',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `log_problem_contribution_reports`;
