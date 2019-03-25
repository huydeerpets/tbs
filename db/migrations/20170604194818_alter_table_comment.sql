
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table contribution_total_follows comment 'Contribution Total Follows';
alter table goose_db_version comment 'DBVersion';
alter table log_bug_reports comment 'LogBugReport';
alter table log_contribution_images comment 'LogContributionImage';
alter table log_problem_contribution_reports comment 'InapproriateContributionReportLog';
alter table log_questions comment 'QuestionLog';
alter table log_user_contributions comment 'ContributionLog';
alter table user_character_images comment 'CharacterImage';
alter table user_contribution_details comment 'ContributionDetails';
alter table user_contribution_follows comment 'ContributionFollows';
alter table user_contribution_movies comment 'ContributionMovie';
alter table user_contribution_searches comment 'ContributionSearch';
alter table user_contribution_sounds comment 'ContributionSound';
alter table user_contribution_sound_details comment 'ContributionSoundDetails';
alter table user_contribution_tags comment 'ContributionTags';
alter table user_contribution_uploads comment 'Contributionupload';
alter table user_contributions comment 'Contribution';
alter table user_forget_passwords comment 'ForgotPassword';
alter table user_masters comment 'Users';
alter table user_profile_images comment 'ProfileImages';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table contribution_total_follows comment '';
alter table goose_db_version comment '';
alter table log_bug_reports comment '';
alter table log_contribution_images comment '';
alter table log_problem_contribution_reports comment '';
alter table log_questions comment '';
alter table log_user_contributions comment '';
alter table user_character_images comment '';
alter table user_contribution_details comment '';
alter table user_contribution_follows comment '';
alter table user_contribution_movies comment '';
alter table user_contribution_searches comment '';
alter table user_contribution_sounds comment '';
alter table user_contribution_sound_details comment '';
alter table user_contribution_tags comment '';
alter table user_contribution_uploads comment '';
alter table user_contributions comment '';
alter table user_forget_passwords comment '';
alter table user_masters comment '';
alter table user_profile_images comment '';
