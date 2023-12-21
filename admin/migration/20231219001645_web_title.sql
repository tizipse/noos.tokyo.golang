-- +goose Up
-- +goose StatementBegin
create table `web_title`
(
    `id`         int unsigned     not null auto_increment,
    `name`       varchar(120)     not null default '' comment '名称',
    `order`      tinyint unsigned not null default 0 comment '序号：正序',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='官网-职位表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_title`;
-- +goose StatementEnd
