-- Create "db_versions" table
CREATE TABLE `db_versions` (
  `id` integer NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `version` integer NULL,
  PRIMARY KEY (`id`)
);
-- Create index "idx_db_versions_deleted_at" to table: "db_versions"
CREATE INDEX `idx_db_versions_deleted_at` ON `db_versions` (`deleted_at`);
