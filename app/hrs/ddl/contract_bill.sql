-- =============================================
-- 合同/账单 - MySQL DDL脚本
-- 适用版本：MySQL 5.7+ / MySQL 8.0+
-- 字符集：utf8mb4（兼容emoji，适配所有中文场景）
-- 存储引擎：InnoDB（支持事务、外键、行级锁）
-- =============================================

-- 1. 合同信息表（关联用户）
CREATE TABLE IF NOT EXISTS `base_contract` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '合同ID（主键）',
  `contract_no` VARCHAR(50) NOT NULL COMMENT '合同编号（唯一）',
  `ref_biz_id` INT(11) NOT NULL COMMENT '业务关联引用ID',
  `tenant_id` INT(11) NOT NULL COMMENT '关联租户ID',
  `rent_amount` DECIMAL(10,2) NOT NULL COMMENT '月租金（单位：元）',
  `start_date` DATE NOT NULL COMMENT '合同生效日期',
  `end_date` DATE NOT NULL COMMENT '合同到期日期',
  `sign_status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '签署状态（0：未签署，1：已签署）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_contract_no` (`contract_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同信息表';

-- 2. 账单信息表（关联合同）
CREATE TABLE IF NOT EXISTS `base_bill` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '账单ID（主键）',
  `bill_no` VARCHAR(50) NOT NULL COMMENT '账单编号（唯一）',
  `contract_id` INT(11) NOT NULL COMMENT '关联合同ID',
  `amount` DECIMAL(10,2) NOT NULL COMMENT '账单金额（单位：元）',
  `bill_type` TINYINT(1) NOT NULL COMMENT '账单类型（1：租金，2：水电费）',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '支付状态（0：未支付，1：已支付）',
  `due_date` DATE NOT NULL COMMENT '到期日期',
  `pay_time` DATETIME DEFAULT NULL COMMENT '支付时间（未支付时为NULL）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_bill_no` (`bill_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账单信息表';
