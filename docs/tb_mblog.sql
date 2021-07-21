/*
 Navicat Premium Data Transfer

 Source Server         : weibo_dev
 Source Server Type    : MySQL
 Source Server Version : 50616
 Source Host           : localhost
 Source Schema         : db_weibo

 Target Server Type    : MySQL
 Target Server Version : 50616
 File Encoding         : 65001

 Date: 21/07/2021 17:44:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_mblog
-- ----------------------------
DROP TABLE IF EXISTS `tb_mblog`;
CREATE TABLE `tb_mblog` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `blogId` varchar(200) NOT NULL DEFAULT '',
  `name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `text` varchar(3000) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `imgs` varchar(3000) NOT NULL DEFAULT '',
  `scheme` varchar(300) NOT NULL DEFAULT '',
  `time_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_blogId` (`blogId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=133 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
