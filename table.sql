-- 所有表使用 utf-8mb4 编码，utf8mb4_unicode_ci 排序规则

DROP DATABASE IF EXISTS cloud_disk;

CREATE DATABASE cloud_disk;

USE cloud_disk;

-- 用户信息
CREATE TABLE `user_basic` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`identity` varchar(36) DEFAULT NULL,

	`name` varchar(60) DEFAULT NULL,
	`password` varchar(32) DEFAULT NULL,
	`email` varchar(100) DEFAULT NULL,

	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	`deleted_at` datetime DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 公共文件存储池
CREATE TABLE `repository_pool` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`identity` varchar(36) DEFAULT NULL,

	`hash` varchar(32) DEFAULT NULL COMMENT '文件的唯一标识',
	`name` varchar(255) DEFAULT NULL COMMENT '文件名称',
	`ext` varchar(30) DEFAULT NULL COMMENT '文件扩展名',
	`size` int(11) DEFAULT NULL COMMENT '文件大小',
	`path` varchar(255) DEFAULT NULL COMMENT '文件路径',

	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	`deleted_at` datetime DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- 用户存储池
CREATE TABLE `user_repository` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`identity` varchar(36) DEFAULT NULL,

	`parent_id` int(11) DEFAULT NULL COMMENT '父级文件层级, 0-【文件夹】',
	`user_identity` varchar(36) DEFAULT NULL COMMENT '对应用户的唯一标识',
	`repository_identity` varchar(36) DEFAULT NULL COMMENT '公共池中文件的唯一标识',
	`ext` varchar(255) DEFAULT NULL COMMENT '文件或文件夹类型',
	`name` varchar(255) DEFAULT NULL COMMENT '用户定义的文件名',

	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	`deleted_at` datetime DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 文件分享
CREATE TABLE `share_basic` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`identity` varchar(36) DEFAULT NULL,

	`user_identity` varchar(36) DEFAULT NULL COMMENT '对应用户的唯一标识',
	`repository_identity` varchar(36) DEFAULT NULL COMMENT '公共池中文件的唯一标识',
	`user_repository_identity` varchar(36) DEFAULT NULL COMMENT '用户池子中的唯一标识',
	`expired_time` int(11) DEFAULT NULL COMMENT '失效时间，单位秒,【0-永不失效】',
	`click_num` int(11) DEFAULT '0' COMMENT '点击次数',

	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	`deleted_at` datetime DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- 插入测试账号
 insert into `user_basic` (name, password, email) values ('test','c06db68e819be6ec3d26c6038d8e8d1f','123@123.com');