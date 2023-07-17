-- テーブルスキーマー定義
CREATE TABLE `users` (
    `id`                      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `create_at`               DATETIME(6) NOT NULL COMMENT 'ユーザ作成日時',
    `update_at`               DATETIME(6) NOT NULL COMMENT 'ユーザ修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザーテーブル';
