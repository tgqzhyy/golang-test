CREATE TABLE `permission` (
  `permission_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `app_id` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '应用编号',
  `permission_key` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '权限Key',
  `permission_name` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '权限名称',
  `permission_desc` varchar(512) COLLATE utf8mb4_bin NOT NULL COMMENT '权限描述',
  `parent_permission_id` int(11) NOT NULL COMMENT '父级权限编号',
  `permission_type` tinyint(4) NOT NULL COMMENT '权限类型',
  `permission_detail` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '权限详情',
  `is_deleted` tinyint(4) NOT NULL COMMENT '是否删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`permission_id`),
  UNIQUE KEY `uk_permission_permission_key` (`permission_key`),
  UNIQUE KEY `uk_permission_app_id_permission_name` (`app_id`,`permission_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='权限';