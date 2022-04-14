
-- +migrate Up
CREATE TABLE `block` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `block_number` BIGINT NOT NULL,
    `parent_hash` VARCHAR(100) NOT NULL COMMENT '',
    `difficulty` BIGINT NOT NULL COMMENT '',
    `hash` VARCHAR(100) NOT NULL COMMENT '',
    `transactions_count` BIGINT NOT NULL COMMENT '',
    `timestamp` BIGINT NOT NULL COMMENT '',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    UNIQUE INDEX (`block_number`),
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='block 資料';

-- +migrate Down
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `block`;
