CREATE TABLE sys.currency_info (
	id BIGINT auto_increment NOT NULL COMMENT '主键',
	currency_code varchar(100) NOT NULL COMMENT '货币代号',
	currency_type INT NOT NULL COMMENT '货币类型，1-纸币，2-硬币',
	currency_name varchar(100) NOT NULL COMMENT '货币名称',
	nominal_value DECIMAL(12,0) NOT NULL COMMENT '货币面值',
	weight_in_gram DECIMAL(8,2) NOT NULL COMMENT '货币重量',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '更新时间',
	delete_flag TINYINT UNSIGNED DEFAULT 0 NOT NULL COMMENT '删除标记，0-正常，1-已删除',
	CONSTRAINT currency_info_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='货币信息';

CREATE TABLE `forex_price` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `src_currency_code` varchar(100) NOT NULL COMMENT '持有货币代号',
  `dest_currency_code` varchar(100) NOT NULL COMMENT '兑换目标货币代号',
  `base_price` decimal(12,2) NOT NULL COMMENT '交易汇率',
  `exchange_date` date NOT NULL COMMENT '兑换日期',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_flag` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '删除标记，0-正常，1-已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='外汇牌价信息';