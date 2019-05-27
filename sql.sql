CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'passport_Id',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户实名',
  `idcard` varchar(64) NOT NULL DEFAULT '' COMMENT '证件号码，加密',
  `mobile` varchar(64) NOT NULL DEFAULT '' COMMENT '用户开二类卡时手机号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uid` (`uid`),
  KEY `idx_idcard` (`idcard`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='个人信息主表'