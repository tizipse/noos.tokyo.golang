-- +goose Up
-- +goose StatementBegin
create table `web_seo`
(
    `id`          int unsigned not null auto_increment,
    `channel`     varchar(20)  not null default '' comment '渠道：member=成员',
    `channel_id`  varchar(64)  not null comment '渠道 ID',
    `title`       varchar(255) not null default '' comment '标题',
    `keyword`     varchar(255) not null default '' comment '关键词',
    `description` varchar(255) not null default '' comment '描述',
    `created_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp             default null,
    primary key (`id`),
    key (`channel_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='官网-SEO 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_seo`;
-- +goose StatementEnd
