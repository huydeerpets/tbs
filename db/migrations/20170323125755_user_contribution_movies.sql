
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_movies` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'UserContributionID',
    `movie_type` INT(3) DEFAULT 1 COMMENT 'MovieType',
    `movie_id` VARCHAR(100) COMMENT 'MovieID',
    `movie_status` INT(3) DEFAULT 1 COMMENT 'MovieStatus',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_movie_status_index` (user_contribution_id, movie_status)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_movies`;
