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
