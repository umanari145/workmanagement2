
-- +migrate Up
CREATE TABLE `rooms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room_name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `type` tinyint(1) DEFAULT '10',
  `status` tinyint(1) DEFAULT '10',
  `note` text,
  `created` timestamp NULL DEFAULT NULL,
  `modified` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
DROP TABLE `rooms`;