-- +goose Up
-- +goose StatementBegin
create table `sys_role`
(
    `id`         int unsigned not null auto_increment,
    `name`       varchar(32)  not null default '' comment '名称',
    `summary`    varchar(255) not null default '' comment '简介',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '系统-角色表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sys_role`;
-- +goose StatementEnd
