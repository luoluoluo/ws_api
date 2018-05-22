CREATE DATABASE IF NOT EXISTS `ws_task` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `ws_task`;

CREATE TABLE `user` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(80) NOT NULL DEFAULT '' COMMENT '微信openid',
  `name` VARCHAR(80) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别（1 男 2 女）',
  `session_key` VARCHAR(80) NOT NULL DEFAULT '' COMMENT '微信session_key',
  `update_time` INT(11) NOT NULL,
  `create_time` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `openid` (`openid`),
  KEY `name` (`name`)
) ENGINE=InnoDB COMMENT '用户表';

CREATE TABLE `task` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) NOT NULL,
  `text` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文字内容',
  `images` VARCHAR(2000) NOT NULL DEFAULT '' COMMENT '图片内容，多张用逗号隔开',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0 待审， 1 正常, 2 删除',
  `create_time` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB COMMENT '任务表';

CREATE TABLE `comment` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `task_id` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `text` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文字内容',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0 待审， 1 正常， 2 删除',
  `create_time` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `task_id` (`task_id`)
) ENGINE=InnoDB COMMENT '评论表';