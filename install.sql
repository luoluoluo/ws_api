CREATE DATABASE IF NOT EXISTS `ws_task` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `ws_task`;

CREATE TABLE `user` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(80) NOT NULL DEFAULT '' COMMENT '微信openid',
  `name` VARCHAR(80) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别（0 女，1 男）',
  `update_time` INT(11) NOT NULL,
  `create_time` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `openid` (`openid`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户表';

CREATE TABLE `task` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) NOT NULL,
  `text` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文字内容',
  `images` VARCHAR(2000) NOT NULL DEFAULT '' COMMENT '图片内容，多张用逗号隔开',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0 隐藏， 1 正常',
  `create_time` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '任务表';