-- +migrate Up
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) NOT NULL,
  `email` varchar(100) UNIQUE NOT NULL,
  `birthday` datetime NULL DEFAULT NULL,
  `pass_code` int(11) DEFAULT NULL,
  `weight` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `first_name`, `last_name`, `email`, `birthday`, `pass_code`, `weight`)
VALUES
	(1,'2017-09-19 12:11:55','2017-09-19 12:11:55',NULL,'Bruce','Wayne','bwayne@wayne-enterprice','1988-02-20 00:00:00',NULL,NULL);

-- +migrate Down
DROP TABLE `users`;
