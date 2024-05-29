create table sys_user
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_name   varchar(50)                            not null comment '用户名',
    nick_name   varchar(150) default ''                not null comment '昵称',
    avatar      varchar(150) default ''                not null comment '头像',
    password    varchar(100)                           not null comment '密码',
    salt        varchar(40)                            not null comment '加密盐',
    email       varchar(100) default ''                not null comment '邮箱',
    mobile      varchar(100) default ''                not null comment '手机号',
    user_status tinyint                                not null comment '状态  0：禁用   1：正常',
    dept_id     bigint                                 not null comment '部门id',
    job_id      bigint                                 not null comment '岗位id',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    del_flag    tinyint      default 1                 not null comment '是否删除  0：已删除  1：正常',
    constraint name
        unique (user_name)
)
    comment '用户管理';

INSERT INTO sys_user (id, user_name, nick_name, avatar, password, salt, email, mobile, user_status, dept_id, create_by,
                      create_time, update_by, update_time, del_flag, job_id)
VALUES (1, 'admin', '超管管理员', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '123456', 'sfsdfsdfs', '1002219331@qq.com', '18613030352', 1, 4, 'admin',
        current_time, 'admin', current_time, 1, 1);
INSERT INTO sys_user (id, user_name, nick_name, avatar, password, salt, email, mobile, user_status, dept_id, create_by,
                      create_time, update_by, update_time, del_flag, job_id)
VALUES (2, 'developer', '开发人员', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '123456', 'sfsdfsdfs', '1002219331@qq.com', '18613030351', 1, 4, 'admin',
        current_time, 'admin', current_time, 1, 1);
INSERT INTO sys_user (id, user_name, nick_name, avatar, password, salt, email, mobile, user_status, dept_id, create_by,
                      create_time, update_by, update_time, del_flag, job_id)
VALUES (3, 'test', '测试人员', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '123456', 'sfsdfsdfs', '1002219331@qq.com', '18613030350', 1, 4, 'admin',
        current_time, 'admin', current_time, 1, 1);
