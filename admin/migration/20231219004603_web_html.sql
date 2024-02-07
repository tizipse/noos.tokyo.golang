-- +goose Up
-- +goose StatementBegin
create table `web_html`
(
    `id`         int unsigned not null auto_increment,
    `channel`    varchar(20)  not null default '' comment '渠道：member=成员；original=产品',
    `channel_id` varchar(64)  not null comment '渠道 ID',
    `content`    text comment '内容',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key (`channel_id`)
) auto_increment = 1000
  collate = utf8mb4_unicode_ci comment ='官网-HTML 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_html`;
-- +goose StatementEnd
