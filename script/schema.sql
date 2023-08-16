CREATE TABLE `employee` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL,
  `name` varchar(128) NOT NULL,
  `mobile` varchar(128) DEFAULT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(256) DEFAULT NULL,
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `country` varchar(128) DEFAULT NULL,
  `tax_country` varchar(128) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `passport` varchar(128) DEFAULT NULL,
  `id_card` varchar(128) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,

  `postal_code` varchar(128) DEFAULT NULL,
  `verified` tinyint DEFAULT '0',
  `bank_accounts` json DEFAULT NULL,
  `attachments` json DEFAULT NULL,
  `seed` varchar(16) NOT NULL,
  `created_by` bigint NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `modified_by` bigint NOT NULL DEFAULT '0',
  `modified_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `enabled` tinyint DEFAULT '1',
  `deleted` tinyint DEFAULT '0',
  `vsn` bigint DEFAULT '0',
  `account_id` bigint DEFAULT '0',
  `card_type` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email` (`email`),
  UNIQUE KEY `uk_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='雇员';

CREATE TABLE IF NOT EXISTS `verification_decision` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `body` text NOT NULL COMMENT '内容体',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='验证判定';

CREATE TABLE IF NOT EXISTS `verification_session` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `session_id` varchar(64) NOT NULL COMMENT '会话id',
  `employee_id` bigint NOT NULL DEFAULT 0 COMMENT '雇员id',
  `status` varchar(64) NOT NULL COMMENT '状态',
  `reason` varchar(256) COMMENT '原因',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_session_id`(`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='验证会话';

insert into verification_session(session_id,employee_id,status)
select cast(uid as char), uid, 'approved'
from employee
where verified = 1;
