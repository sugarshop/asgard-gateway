use orders;

CREATE TABLE IF NOT EXISTS lemonsqueezy_orders
(
    `id`                bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'main key',
    `trade_order_id`    varchar(255)        NOT NULL DEFAULT '' COMMENT 'trade order id for a user payment transaction',
    `uid`               varchar(255)        NOT NULL DEFAULT '' COMMENT 'user id, format: clerk_xxxxx, auth0_xxxxx',
    `store_id`          bigint(20)          NOT NULL DEFAULT 0 COMMENT 'The ID of the store this order belongs to.',
    `identifier`        varchar(255)        NOT NULL DEFAULT 0 COMMENT 'The unique identifier (UUID) for this order',
    `status`            varchar(20)         NOT NULL DEFAULT 0 COMMENT 'The status of the order. One of pending, failed, paid, refunded',
    `order_id`          bigint(20)          NOT NULL DEFAULT 0 COMMENT 'The ID of the order',
    `product_id`        bigint(20)          NOT NULL DEFAULT 0 COMMENT 'The ID of the product',
    `variant_id`        bigint(20)          NOT NULL DEFAULT 0 COMMENT 'The ID of the product variant',
    `product_name`      varchar(20)         NOT NULL DEFAULT 0 COMMENT 'The name of the product',
    `variant_name`      varchar(20)         NOT NULL DEFAULT 0 COMMENT 'The name of the product variant',
    `checkout_link`     varchar(20)         NOT NULL DEFAULT 0 COMMENT 'checkout link for this user',
    `order_create_time` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'An ISO-8601 formatted date-time string indicating when the order was created.',
    `created_at`        datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_trade_order_id` (`trade_order_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;