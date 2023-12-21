-- +goose Up
-- +goose StatementBegin
create table `web_member`
(
    `id`          varchar(64)      not null,
    `title_id`    int unsigned     not null default 0 comment '职位ID',
    `name`        varchar(120)     not null default '' comment '名称',
    `nickname`    varchar(120)     not null default '' comment '别称',
    `thumb`       varchar(255)     not null default '' comment '头像',
    `ins`         varchar(255)     not null default '' comment 'INS',
    `order`       tinyint unsigned not null default 0 comment '序号：正序',
    `is_delegate` tinyint unsigned not null default 0 comment '是否代表：1=是；2=否',
    `is_enable`   tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at`  timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp                 default null,
    primary key (`id`),
    key (`title_id`)
) collate = utf8mb4_unicode_ci comment ='官网-成员表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `web_member`;
-- +goose StatementEnd
