
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_sound_lengths` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'UserContributionID',
    `second` INT(10) DEFAULT 0 COMMENT 'RegenerationTime',
    `length` INT(10) DEFAULT 0 COMMENT 'WordCount',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_index` (user_contribution_id)
) COMMENT = 'ContributionVoiceLength' ENGINE = InnoDB ;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_sound_lengths`;
