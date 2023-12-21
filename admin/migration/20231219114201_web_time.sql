-- +goose Up
-- +goose StatementBegin
create table `web_time`
(
    `id`         int unsigned     not null auto_increment,
    `name`       varchar(120)     not null default '' comment '名称',
    `content`    varchar(120)     not null default '' comment '内容',
    `status`     varchar(5)       not null default '' comment '状态：open=开放；close=关闭',
    `order`      tinyint unsigned not null default 0 comment '序号：正序',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='官网-营业时间表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_time`;
-- +goose StatementEnd
