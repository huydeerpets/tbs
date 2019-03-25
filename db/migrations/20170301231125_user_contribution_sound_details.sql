
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_sound_details` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'UserContributionID',
    `priority` INT COMMENT 'Priority',
    `talk_type` INT COMMENT 'TalkType',
    `body` VARCHAR(256) COMMENT 'Body',
    `body_sound` VARCHAR(256) COMMENT 'BodyReadAloud',
    `voice_type` INT COMMENT 'SoundType',
    `created_at` DATETIME NULL COMMENT 'Created At',
    `updated_at` DATETIME NULL COMMENT 'Updated At',
    `deleted_at` DATETIME NULL COMMENT 'Deleted At',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_priority_index` (user_contribution_id, priority)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_sound_details`;
