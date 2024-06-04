create table ums_member_tag
(
    id                  bigint auto_increment
        primary key,
    tag_name            varchar(100)   not null comment '标签名称',
    finish_order_count  int            not null comment '自动打标签完成订单数量',
    finish_order_amount bigint not null comment '自动打标签完成订单金额'
)
    comment '用户标签表';

