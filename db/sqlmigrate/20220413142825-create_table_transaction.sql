

-- +migrate Up
CREATE TABLE `transaction_record` (
                         `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `block_number` VARCHAR(36) NOT NULL,
                         `tx_hash` VARCHAR(100) NOT NULL COMMENT '',
                         `value` TEXT NOT NULL COMMENT '',
                         `gas` BIGINT NOT NULL COMMENT '',
                         `gas_price` BIGINT NOT NULL COMMENT '',
                         `nonce` BIGINT NOT NULL COMMENT '',
                         `from` VARCHAR(100) NOT NULL COMMENT '',
                         `to` VARCHAR(100) NOT NULL COMMENT '',
                         `data` TEXT NOT NULL COMMENT '',
                         `pending` boolean NOT NULL COMMENT '',
                         `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
                         `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
                         PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='transaction Record資料';

-- +migrate Down
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `transaction_record`;

