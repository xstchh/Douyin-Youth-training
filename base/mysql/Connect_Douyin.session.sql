DROP TABLE if exists 'user_table';
CREATE TABLE 'user_table' (
    id BIGINT primary key auto_increment comment '主键',
    name varchar(255) not null comment '昵称',
    follow_count INT default 0 comment '关注数',
    follower_count INT default 0 comment '粉丝数',
    is_follow BOOLEAN comment '是否关注',
    avatar varchar(255) comment '头像',
    background_image varchar(255) comment '个人资料背景',
    signature_ varchar(255) comment '个人简介',
    total_favorited varchar(255) comment '获赞数',
    work_count INT default 0 comment '作品数',
    favorited_count INT default 0 comment '喜欢数'
) comment = '用户表';
DROP TABLE if exists 'video_table';
CREATE TABLE 'video_table'(
    id BIGINT primary key auto_increment comment '主键',
    user_id BIGINT not null comment '视频发布者',
    play_url varchar(255) not null comment '视频地址',
    cover_url varchar(255) not null comment '封面地址',
    favorited_count BIGINT default 0 comment '点赞数',
    comment_count BIGINT default 0 comment '评论数',
    is_favorite BOOLEAN default false comment '是否点赞',
    title varchar(255) not null comment '标题'
) comment = '视频表';
DROP TABLE if exists 'comment_table';
CREATE TABLE 'comment_table'(
    id BIGINT primary key auto_increment comment '主键',
    user_id BIGINT not null comment '评论发布者',
    content varchar(255) not null comment '评论内容',
    create_time DATE not null comment '发布时间'
) comment = '评论表';
DROP TABLE if exists 'follower_table';
CREATE TABLE 'follower_table'(
    id BIGINT primary key auto_increment comment '主键',
    user_id BIGINT not null comment '被关注者',
    fans_id BIGINT not null comment '粉丝id'
) comment = '粉丝表';
DROP TABLE if exists 'likes_table';
CREATE TABLE 'likes_table'(
    id BIGINT primary key auto_increment comment '主键',
    user_id BIGINT not null comment '被关注者',
    video_id BIGINT not null comment '粉丝id'
) comment = '点赞表';