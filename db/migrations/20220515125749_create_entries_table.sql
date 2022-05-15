-- migrate:up
CREATE TABLE `entries` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `content` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
-- migrate:down
DROP TABLE `entries`;