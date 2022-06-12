/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 127.0.0.1:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 13/06/2022 00:15:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_info_id` bigint DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_videos_comments` (`video_id`),
  KEY `fk_user_infos_comments` (`user_info_id`),
  CONSTRAINT `fk_user_infos_comments` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_videos_comments` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_favor_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_favor_videos`;
CREATE TABLE `user_favor_videos` (
  `user_info_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  PRIMARY KEY (`user_info_id`,`video_id`),
  KEY `fk_user_favor_videos_video` (`video_id`),
  CONSTRAINT `fk_user_favor_videos_user_info` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_user_favor_videos_video` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_infos
-- ----------------------------
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `follow_count` bigint DEFAULT NULL,
  `follower_count` bigint DEFAULT NULL,
  `is_follow` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_logins
-- ----------------------------
DROP TABLE IF EXISTS `user_logins`;
CREATE TABLE `user_logins` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_info_id` bigint DEFAULT NULL,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`,`username`),
  KEY `fk_user_infos_user` (`user_info_id`),
  CONSTRAINT `fk_user_infos_user` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_relations
-- ----------------------------
DROP TABLE IF EXISTS `user_relations`;
CREATE TABLE `user_relations` (
  `user_info_id` bigint NOT NULL,
  `follow_id` bigint NOT NULL,
  PRIMARY KEY (`user_info_id`,`follow_id`),
  KEY `fk_user_relations_follows` (`follow_id`),
  CONSTRAINT `fk_user_relations_follows` FOREIGN KEY (`follow_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_user_relations_user_info` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_info_id` bigint DEFAULT NULL,
  `play_url` longtext COLLATE utf8mb4_unicode_ci,
  `cover_url` longtext COLLATE utf8mb4_unicode_ci,
  `favorite_count` bigint DEFAULT NULL,
  `comment_count` bigint DEFAULT NULL,
  `is_favorite` tinyint(1) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_infos_videos` (`user_info_id`),
  CONSTRAINT `fk_user_infos_videos` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
