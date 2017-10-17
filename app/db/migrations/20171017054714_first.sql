
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- CREATE USER 'developer'@'%' IDENTIFIED BY 'password';
-- grant all on media.* to 'developer'@'%' identified by 'password';

CREATE TABLE `article_categories` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `name` int(11) NOT NULL,
      `media_id` int(11) NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `articles` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `title` varchar(100) NOT NULL DEFAULT '',
      `content` text NOT NULL,
      `author_id` int(11) NOT NULL,
      `media_id` int(11) NOT NULL,
      `article_category_id` int(11) NOT NULL,
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

CREATE TABLE `medias` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `name` int(11) NOT NULL,
      `publisher_id` int(11) NOT NULL,
      `domain` int(11) DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `authors` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `email` varchar(50) NOT NULL DEFAULT '',
      `first_name` varchar(50) NOT NULL DEFAULT '',
      `last_name` varchar(50) NOT NULL DEFAULT '',
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `email` varchar(100) NOT NULL DEFAULT '',
      `first_name` varchar(50) NOT NULL DEFAULT '',
      `last_name` varchar(50) NOT NULL DEFAULT '',
      `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0:unknown, 1:men, 2:women',
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP USER 'developer'
DROP TABLE `article_categories`
DROP TABLE `articles`
DROP TABLE `medias`
DROP TABLE `authors`
DROP TABLE `users`
