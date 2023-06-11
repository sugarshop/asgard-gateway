use rights; -- rights db, store user rights belongs to many product.

CREATE TABLE IF NOT EXISTS chattyai_rights
(
    `id`                       bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'main key',
    `uid`                      varchar(255)        NOT NULL DEFAULT '' COMMENT 'user id, format: clerk_xxxxx, auth0_xxxxx',
    `token_quota`              bigint(20)          NOT NULL DEFAULT 0 COMMENT 'authorized token quota in this subscription period',
    `token_used`               bigint(20)          NOT NULL DEFAULT 0 COMMENT 'used token_quota in this subscription period',
    `token_used_total`         bigint(20)          NOT NULL DEFAULT 0 COMMENT 'total used token_quota in the user life time',
    `conversation_quota`       bigint(20)          NOT NULL DEFAULT 0 COMMENT 'authorized conversation quota',
    `conversation_used`        bigint(20)          NOT NULL DEFAULT 0 COMMENT 'authorized conversation quota in this subscription period',
    `conversation_used_total`  bigint(20)          NOT NULL DEFAULT 0 COMMENT 'used authorized conversation quota in the user life time',
    `assistant_quota`          bigint(20)          NOT NULL DEFAULT 0 COMMENT 'authorized assistant_quota quota',
    `assistant_used`           bigint(20)          NOT NULL DEFAULT 0 COMMENT 'used assistant_quota quota in this subscription period',
    `gpt_4_access`             boolean             NOT NULL DEFAULT FALSE COMMENT 'if true, has gpt-4 access, or false',
    `api_access`               boolean             NOT NULL DEFAULT FALSE COMMENT 'if true, has external chat api access, or false',
    `subscription_date`        datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'subscription start date',
    `subscription_update_date` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'subscription update date',
    `subscription_end_date`    datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'subscription end date',
    `created_at`               datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`               datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;