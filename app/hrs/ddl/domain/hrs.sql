-- ----------------------------
-- 基础数据字典
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

-- ----------------------------
-- 智管系统核心业务表DDL脚本
-- 适用版本：MySQL 8.0+
-- 字符集：utf8mb4（支持中文/Emoji）
-- 存储引擎：InnoDB（支持事务/外键）
-- 执行顺序：property → tenant → contract → bill → operation_log
-- ----------------------------

-- 1. 房源信息表
CREATE TABLE IF NOT EXISTS `hrs_property` (
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
CREATE TABLE IF NOT EXISTS `hrs_tenant` (
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
