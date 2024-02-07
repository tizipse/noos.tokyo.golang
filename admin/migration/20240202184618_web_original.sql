-- +goose Up
-- +goose StatementBegin
create table `web_original`
(
    `id`         varchar(64)      not null,
    `name`       varchar(120)     not null default '' comment '名称',
    `thumb`      varchar(255)     not null default '' comment '头像',
    `ins`        varchar(255)     not null default '' comment 'INS',
    `summary`    varchar(255)     not null default '' comment '简介',
    `order`      tinyint unsigned not null default 0 comment '序号：正序',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`)
) collate = utf8mb4_unicode_ci comment ='官网-产品表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_original`;
-- +goose StatementEnd
