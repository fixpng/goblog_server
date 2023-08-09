-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: gvb_db
-- ------------------------------------------------------
-- Server version	5.7.41

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `advert_models`
--

DROP TABLE IF EXISTS `advert_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `advert_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `href` longtext,
  `images` longtext,
  `is_show` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `banner_models`
--

DROP TABLE IF EXISTS `banner_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `banner_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `path` longtext,
  `hash` longtext,
  `name` varchar(38) DEFAULT NULL,
  `image_type` bigint(20) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chat_models`
--

DROP TABLE IF EXISTS `chat_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chat_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `nick_name` varchar(15) DEFAULT NULL,
  `avatar` varchar(128) DEFAULT NULL,
  `content` varchar(256) DEFAULT NULL,
  `ip` varchar(32) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `msg_type` tinyint(4) DEFAULT NULL,
  `is_group` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=205 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comment_models`
--

DROP TABLE IF EXISTS `comment_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `parent_comment_id` bigint(20) unsigned DEFAULT NULL,
  `content` varchar(256) DEFAULT NULL,
  `digg_count` tinyint(4) DEFAULT '0',
  `comment_count` tinyint(4) DEFAULT '0',
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_comment_models_sub_comments` (`parent_comment_id`),
  KEY `fk_comment_models_user` (`user_id`),
  CONSTRAINT `fk_comment_models_sub_comments` FOREIGN KEY (`parent_comment_id`) REFERENCES `comment_models` (`id`),
  CONSTRAINT `fk_comment_models_user` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `fade_back_models`
--

DROP TABLE IF EXISTS `fade_back_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fade_back_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `content` varchar(128) DEFAULT NULL,
  `apply_content` varchar(128) DEFAULT NULL,
  `is_apply` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `log_stash_models`
--

DROP TABLE IF EXISTS `log_stash_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_stash_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `ip` varchar(32) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `level` tinyint(4) DEFAULT NULL,
  `content` varchar(128) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `login_data_models`
--

DROP TABLE IF EXISTS `login_data_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_data_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `ip` varchar(20) DEFAULT NULL,
  `nick_name` varchar(42) DEFAULT NULL,
  `token` varchar(256) DEFAULT NULL,
  `device` varchar(256) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `login_type` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_login_data_models_user_model` (`user_id`),
  CONSTRAINT `fk_login_data_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `menu_banner_models`
--

DROP TABLE IF EXISTS `menu_banner_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_banner_models` (
  `menu_id` bigint(20) unsigned DEFAULT NULL,
  `banner_id` bigint(20) unsigned DEFAULT NULL,
  `sort` smallint(6) DEFAULT NULL,
  KEY `fk_menu_banner_models_menu_model` (`menu_id`),
  KEY `fk_menu_banner_models_banner_model` (`banner_id`),
  CONSTRAINT `fk_menu_banner_models_banner_model` FOREIGN KEY (`banner_id`) REFERENCES `banner_models` (`id`),
  CONSTRAINT `fk_menu_banner_models_menu_model` FOREIGN KEY (`menu_id`) REFERENCES `menu_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `menu_models`
--

DROP TABLE IF EXISTS `menu_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `path` varchar(256) DEFAULT NULL,
  `slogan` varchar(64) DEFAULT NULL,
  `abstract` longtext,
  `abstract_time` bigint(20) DEFAULT NULL,
  `banner_time` bigint(20) DEFAULT NULL,
  `sort` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message_models`
--

DROP TABLE IF EXISTS `message_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint(20) unsigned NOT NULL,
  `send_user_nick_name` varchar(42) DEFAULT NULL,
  `send_user_avatar` longtext,
  `rev_user_id` bigint(20) unsigned NOT NULL,
  `rev_user_nick_name` varchar(42) DEFAULT NULL,
  `rev_user_avatar` longtext,
  `is_read` tinyint(1) DEFAULT '0',
  `content` longtext,
  PRIMARY KEY (`id`,`send_user_id`,`rev_user_id`),
  KEY `fk_message_models_send_user_model` (`send_user_id`),
  KEY `fk_message_models_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_message_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_message_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tag_models`
--

DROP TABLE IF EXISTS `tag_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_collect_models`
--

DROP TABLE IF EXISTS `user_collect_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_collect_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` varchar(32) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `nick_name` varchar(36) DEFAULT NULL,
  `user_name` varchar(36) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `avatar` text,
  `email` varchar(128) DEFAULT NULL,
  `tel` varchar(18) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `token` varchar(64) DEFAULT NULL,
  `ip` varchar(20) DEFAULT NULL,
  `role` tinyint(4) DEFAULT '1',
  `sign_status` bigint(20) DEFAULT NULL,
  `integral` bigint(20) DEFAULT '0',
  `sign` varchar(128) DEFAULT NULL,
  `link` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-08  0:06:21
