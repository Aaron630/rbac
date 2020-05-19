/*
 Navicat Premium Data Transfer

 Source Server         : hyper-V
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : 192.168.218.175:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 19/05/2020 18:17:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_module
-- ----------------------------
DROP TABLE IF EXISTS `admin_module`;
CREATE TABLE `admin_module`  (
  `module_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` mediumint(8) UNSIGNED NOT NULL COMMENT '父权限id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '功能名称',
  `action` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作',
  PRIMARY KEY (`module_id`) USING BTREE,
  INDEX `module_id`(`module_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台管理模块表（模块表，一个模块代表一个权限）' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `role_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台用户角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '超级管理员');
INSERT INTO `admin_role` VALUES (2, '秘书');

-- ----------------------------
-- Table structure for admin_role_module
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_module`;
CREATE TABLE `admin_role_module`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `module_id` int(10) UNSIGNED NOT NULL COMMENT '权限模块id',
  `role_id` int(10) UNSIGNED NOT NULL COMMENT '角色Id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `module_and_role`(`module_id`, `role_id`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '管理员和权限关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `admin_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `admin_account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登陆账号',
  `admin_passwd` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户登陆密码',
  `admin_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `admin_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机号码',
  `admin_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户头像地址',
  `admin_gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '性别(1男，2女, 0 保密)',
  `admin_mail` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `admin_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 2 COMMENT '用户状态(0关闭，1启用)',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0表示不删除，1表示删除',
  `create_date` timestamp(0) NOT NULL COMMENT '创建时间',
  `create_by` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建人',
  `last_update_date` timestamp(0) NULL DEFAULT NULL COMMENT '最后更新时间',
  `update_by` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '最后更新人',
  PRIMARY KEY (`admin_id`) USING BTREE,
  INDEX `admin_account`(`admin_account`) USING BTREE,
  INDEX `admin_name`(`admin_name`) USING BTREE,
  INDEX `admin_phone`(`admin_phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台管理用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'superman', '7D3DE28DBEF46ED567B7F9B179FB1362', 'superman', '13145215216', NULL, 1, 'xxxx@qq.com', 2, 0, '2020-05-19 11:18:00', 'superman', NULL, NULL);
INSERT INTO `admin_user` VALUES (2, 'superwoman', '7D3DE28DBEF46ED567B7F9B179FB1362', 'superwoman', '13963152456', NULL, 1, 'ggg@qq.com', 2, 0, '2020-05-19 11:22:00', 'superman', NULL, NULL);

-- ----------------------------
-- Table structure for admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_role`;
CREATE TABLE `admin_user_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_id` int(10) UNSIGNED NOT NULL COMMENT '角色Id',
  `admin_id` int(10) UNSIGNED NOT NULL COMMENT '管理员id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `role_admin`(`role_id`, `admin_id`) USING BTREE,
  INDEX `admin_id`(`admin_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '管理员和角色关系表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
