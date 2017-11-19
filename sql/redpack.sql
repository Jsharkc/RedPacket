CREATE DATABASE IF NOT EXISTS `redpack`;

USE `redpack`;

CREATE TABLE IF NOT EXISTS `user` (

	`id` bigint(20) unsigned AUTO_INCREMENT,

	`balance` int(10) unsigned NOT NULL DEFAULT 0,

	PRIMARY KEY (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=1969 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


-- -----------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `redpack` (

	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,

	`uid` bigint(20) unsigned NOT NULL DEFAULT 0,

	`totalmoney` bigint(20) unsigned NOT NULL DEFAULT 0,

	`num` int(10) unsigned NOT NULL DEFAULT 0,

	`blessing` varchar(255) NOT NULL DEFAULT '',

	`password` varchar(255) NOT NULL DEFAULT '',

	`status` smallint(3) unsigned NOT NULL DEFAULT 0,

	`created` datetime,

	PRIMARY KEY (`id`),

	CONSTRAINT `redpack_user_uid` FOREIGN KEY (`uid`) REFERENCES `user` (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


-- -----------------------------------------------------------------------


CREATE TABLE IF NOT EXISTS `grabrecord` (

	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,

	`rpid` bigint(20) unsigned NOT NULL DEFAULT 0,

	`uid` bigint(20) unsigned NOT NULL DEFAULT 0,

	`money` int(10) unsigned NOT NULL DEFAULT 0,

	PRIMARY KEY (`id`),

	CONSTRAINT `packrecord_user_uid` FOREIGN KEY (`uid`) REFERENCES `user` (`id`),

	CONSTRAINT `packrecord_redpack_rpid` FOREIGN KEY (`rpid`) REFERENCES `redpack` (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


