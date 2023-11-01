CREATE TABLE IF NOT EXISTS `comments` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `feed_id` INT NOT NULL,
  `comment` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (id),
   CONSTRAINT `comments_ibfk_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE CASCADE,
   CONSTRAINT `comments_ibfk_feed_id` FOREIGN KEY (`feed_id`)
        REFERENCES `feeds` (`id`) ON DELETE CASCADE
        
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;