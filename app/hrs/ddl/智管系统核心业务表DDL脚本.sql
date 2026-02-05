-- =============================================
-- 数据字典模块 - MySQL DDL脚本
-- 适用版本：MySQL 5.7+ / MySQL 8.0+
-- 字符集：utf8mb4（兼容emoji，适配所有中文场景）
-- 存储引擎：InnoDB（支持事务、外键、行级锁）
-- =============================================

-- 设置脚本执行环境
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 1. 字典类型表（sys_dict_type）
-- 存储字典的分类信息，如"用户状态""订单类型"等
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
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
-- 2. 字典项表（sys_dict_item）
-- 存储具体的字典值，关联字典类型表的dict_code
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_item`;
CREATE TABLE `sys_dict_item` (
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

-- 恢复外键检查（如需添加外键约束，可取消下方注释）
-- ALTER TABLE `sys_dict_item` ADD CONSTRAINT `fk_dict_item_code` FOREIGN KEY (`dict_code`) REFERENCES `sys_dict_type` (`dict_code`) ON DELETE RESTRICT ON UPDATE CASCADE;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- 初始化示例数据（可选）
-- ----------------------------
INSERT INTO `sys_dict_type` (`dict_code`, `dict_name`, `sort`, `remark`) VALUES
('USER_STATUS', '用户状态', 1, '系统用户的启用/禁用状态'),
('ORDER_TYPE', '订单类型', 2, '订单的业务类型分类');

INSERT INTO `sys_dict_item` (`dict_code`, `item_value`, `item_label`, `sort`, `remark`) VALUES
('USER_STATUS', '0', '禁用', 1, '用户账号禁用，无法登录'),
('USER_STATUS', '1', '启用', 2, '用户账号正常，可登录'),
('ORDER_TYPE', '01', '普通订单', 1, '常规商品购买订单'),
('ORDER_TYPE', '02', '秒杀订单', 2, '限时秒杀活动订单'),
('ORDER_TYPE', '03', '拼团订单', 3, '多人拼团订单');



/*
 * 智管系统核心业务表DDL脚本
 * 适用版本：MySQL 8.0+
 * 字符集：utf8mb4（支持中文/Emoji）
 * 存储引擎：InnoDB（支持事务/外键）
 * 执行顺序：property → tenant → contract → bill → operation_log
 */

-- 1. 房源信息表
CREATE TABLE `property` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '房源ID（主键）',
  `property_code` VARCHAR(30) NOT NULL COMMENT '房源编号（唯一标识）',
  `area` DECIMAL(10,2) NOT NULL COMMENT '房源面积（单位：㎡）',
  `room_count` INT(2) NOT NULL COMMENT '房间数量',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '房源状态（0：空闲，1：已出租）',
  `address` VARCHAR(255) NOT NULL COMMENT '房源地址',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_property_code` (`property_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源信息表';

-- 2. 租户信息表
CREATE TABLE `tenant` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '租户ID（主键）',
  `name` VARCHAR(50) NOT NULL COMMENT '租户姓名',
  `id_card` VARCHAR(18) NOT NULL COMMENT '身份证号（唯一）',
  `phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '电子邮箱（可选）',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '租户状态（0：禁用，1：正常）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tenant_id_card` (`id_card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户信息表';

-- 3. 合同信息表（关联房源/租户）
CREATE TABLE `contract` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '合同ID（主键）',
  `contract_no` VARCHAR(50) NOT NULL COMMENT '合同编号（唯一）',
  `property_id` INT(11) NOT NULL COMMENT '关联房源ID',
  `tenant_id` INT(11) NOT NULL COMMENT '关联租户ID',
  `rent_amount` DECIMAL(10,2) NOT NULL COMMENT '月租金（单位：元）',
  `start_date` DATE NOT NULL COMMENT '合同生效日期',
  `end_date` DATE NOT NULL COMMENT '合同到期日期',
  `sign_status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '签署状态（0：未签署，1：已签署）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_contract_no` (`contract_no`),
  -- 外键约束：确保数据引用完整性
  FOREIGN KEY (`property_id`) REFERENCES `property` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同信息表';

-- 4. 账单信息表（关联合同）
CREATE TABLE `bill` (
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
  UNIQUE KEY `uk_bill_no` (`bill_no`),
  FOREIGN KEY (`contract_id`) REFERENCES `contract` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账单信息表';

-- 5. 操作日志表（系统审计）
CREATE TABLE `operation_log` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID（主键，BIGINT避免溢出）',
  `operator_id` INT(11) NOT NULL COMMENT '操作人ID（预留关联用户表）',
  `operator_name` VARCHAR(50) NOT NULL COMMENT '操作人姓名',
  `module` VARCHAR(50) NOT NULL COMMENT '操作模块（如：房源管理、租户管理）',
  `action` VARCHAR(50) NOT NULL COMMENT '操作行为（如：新增、修改、删除）',
  `content` TEXT DEFAULT NULL COMMENT '操作内容（JSON格式存储详细信息）',
  `ip` VARCHAR(50) NOT NULL COMMENT '操作IP地址',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  PRIMARY KEY (`id`),
  -- 索引优化：高频查询模块/操作人
  INDEX `idx_log_module` (`module`),
  INDEX `idx_log_operator_id` (`operator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';

-- 脚本执行完成提示
SELECT '智管系统表结构创建成功' AS `result`;
