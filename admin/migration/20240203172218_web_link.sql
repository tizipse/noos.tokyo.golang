-- +goose Up
-- +goose StatementBegin
create table `web_link`
(
    `id`         int unsigned     not null auto_increment,
    `name`       varchar(120)     not null default '' comment '名称',
    `summary`    varchar(255)     not null default '' comment '简介',
    `url`        varchar(255)     not null default '' comment '链接',
    `order`      tinyint unsigned not null default 0 comment '序号：正序',
    `is_system`  tinyint unsigned not null default 0 comment '是否内置：1=是；2=否；',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='官网-链接表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_link`;
-- +goose StatementEnd
