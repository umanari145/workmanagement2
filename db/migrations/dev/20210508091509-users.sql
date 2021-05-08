
-- +migrate Up
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chatgirl_id` int(11) DEFAULT NULL,
  `username` varchar(200) DEFAULT NULL,
  `japanese_name` varchar(20) DEFAULT NULL,
  `email` varchar(120) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `transfer_account` varchar(100) DEFAULT NULL,
  `user_type` smallint(6) DEFAULT '1',
  `working_type` tinyint(4) DEFAULT NULL,
  `area_type` tinyint(4) DEFAULT NULL,
  `created` timestamp NULL DEFAULT NULL,
  `modified` timestamp NULL DEFAULT NULL,
  `ip` varchar(100) DEFAULT NULL,
  `last_announce_check_time` timestamp NULL DEFAULT NULL,
  `last_access_time` timestamp NULL DEFAULT NULL,
  `is_delete` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
DROP TABLE `users`;