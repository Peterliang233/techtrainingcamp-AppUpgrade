DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
use app;
create table device
(
    id        int auto_increment
        primary key,
    device_id varchar(33) not null comment '设备的id',
    constraint device_device_id_uindex
        unique (device_id)
)
    comment '设备的白名单列表';

create table info
(
    id                  int auto_increment
        primary key,
    version             varchar(33) not null comment '请求版本的api',
    device_platform     varchar(33) not null comment '设备平台',
    device_id           varchar(33) not null comment '设备的ID',
    os_api              int         not null comment '安卓的系统版本',
    channel_number      varchar(33) not null comment '渠道号',
    version_code        varchar(33) not null comment '应用大版本',
    update_version_code varchar(33) not null comment '应用小版本',
    aid                 int         not null comment 'app的唯一标识',
    apu_arch            int         not null comment '设备的cpu架构',
    constraint info_device_id_uindex
        unique (device_id)
)
    comment '客户端上报的参数信息';

create table rule
(
    id                      int auto_increment
        primary key,
    aid                     int          not null comment 'app唯一标识',
    platform                varchar(33)  not null comment '平台',
    download_url            varchar(100) not null comment '命中后包的下载链接',
    update_version_code     varchar(33)  not null comment '包的当前的版本号',
    md_5                    varchar(100) not null comment '包的MD5',
    max_update_version_code varchar(33)  not null comment '可升级的最大版本',
    min_update_version_code varchar(33)  not null comment '可升级的最小版本',
    max_os_api              int          not null comment '支持的最大操作系统版本',
    min_os_api              int          not null comment '支持的最小操作系统版本',
    cpu_arch                varchar(33)  not null comment 'cpu架构',
    channel_number          varchar(33)  not null comment '渠道号',
    title                   varchar(100) not null comment '弹窗标题',
    update_tips             text         not null comment '弹窗的更新文本',
    constraint rule_a_id_uindex
        unique (aid)
)
    comment '新版本的配置规则';

create table user
(
    id       int auto_increment
        primary key,
    username varchar(33) not null comment '用户名',
    password varchar(33) not null comment '密码'
)
    comment '后台用户表';

