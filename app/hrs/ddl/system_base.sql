-- =============================================
-- 操作日志 - MySQL DDL脚本
-- 适用版本：MySQL 5.7+ / MySQL 8.0+
-- 字符集：utf8mb4（兼容emoji，适配所有中文场景）
-- 存储引擎：InnoDB（支持事务、外键、行级锁）
-- =============================================

-- 1. 操作日志表（系统审计）
CREATE TABLE IF NOT EXISTS `operation_log` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID（主键，BIGINT避免溢出）',
  `operator_id` INT(11) NOT NULL COMMENT '操作人ID（预留关联用户表）',
  `operator_name` VARCHAR(50) NOT NULL COMMENT '操作人姓名',
  `biz_module` VARCHAR(50) NOT NULL COMMENT '操作模块（如：房源管理、租户管理）',
  `operation` VARCHAR(50) NOT NULL COMMENT '操作行为（如：新增、修改、删除）',
  `content` TEXT DEFAULT NULL COMMENT '操作内容（JSON格式存储详细信息）',
  `ip` VARCHAR(50) NOT NULL COMMENT '操作IP地址',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  PRIMARY KEY (`id`),
  -- 索引优化：高频查询模块/操作人
  INDEX `idx_log_biz_module` (`biz_module`),
  INDEX `idx_log_operator_id` (`operator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';

-- =============================================
-- 数据字典 - MySQL DDL脚本
-- 适用版本：MySQL 5.7+ / MySQL 8.0+
-- 字符集：utf8mb4（兼容emoji，适配所有中文场景）
-- 存储引擎：InnoDB（支持事务、外键、行级锁）
-- =============================================

-- ----------------------------
-- 1. 字典类型表（sys_dict_type）：存储字典的分类信息，如"用户状态""订单类型"等
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_dict_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dict_code` varchar(50) NOT NULL COMMENT '字典类型编码（唯一标识，如USER_STATUS）',
  `dict_name` varchar(100) NOT NULL COMMENT '字典类型名称（如用户状态）',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态（0-禁用，1-启用）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序号（升序排列）',
  `remark` varchar(500) DEFAULT '' COMMENT '字典类型备注说明',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '软删除标记（0-未删除，1-已删除）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_dict_code` (`dict_code`) USING BTREE COMMENT '字典编码唯一索引',
  KEY `idx_dict_type_status` (`status`) USING BTREE COMMENT '状态索引，提升查询效率'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='字典类型表';

-- ----------------------------
-- 2. 字典项表（sys_dict_item）：存储具体的字典值，关联字典类型表的dict_code
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_dict_item` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dict_code` varchar(50) NOT NULL COMMENT '关联的字典类型编码',
  `item_value` varchar(100) NOT NULL COMMENT '字典项值（如0、1、PAY_ALIPAY）',
  `item_label` varchar(100) NOT NULL COMMENT '字典项标签（如禁用、启用、支付宝支付）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序号（同字典类型下升序排列）',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态（0-禁用，1-启用）',
  `remark` varchar(500) DEFAULT '' COMMENT '字典项备注说明',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '软删除标记（0-未删除，1-已删除）',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_dict_code` (`dict_code`) USING BTREE COMMENT '字典编码索引，关联查询核心索引',
  KEY `idx_dict_item_status` (`status`) USING BTREE COMMENT '状态索引',
  KEY `idx_dict_item_value` (`dict_code`,`item_value`) USING BTREE COMMENT '字典编码+值联合索引，唯一确定一个字典项'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='字典项表';

