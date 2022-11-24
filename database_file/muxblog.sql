CREATE TABLE `users`(
  `id` int(3) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `firstname` VARCHAR(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `email` varchar(100) UNIQUE NOT NULL,
  `password` varchar(100) NOT NULL
)ENGINE=INNODB;

CREATE TABLE categories (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(200) NOT NULL
) ENGINE=INNODB;

CREATE TABLE `comments` (
  `id` int(3) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `id_post_comment` int(3) NOT NULL,
  `user_name_comment` varchar(200) NOT NULL,
  `comment` varchar(200) NOT NULL
)ENGINE=INNODB;

CREATE TABLE posts (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(200) NOT NULL,
  `slug` varchar(200) NOT NULL,
  `img` varchar(100) NOT NULL,
  `body` text NOT NULL,
  `category_id` int(3) NOT NULL,
  `user_id` int(3) NOT NULL,
  `user_name` varchar(200) NOT NULL,
  CONSTRAINT fk_category
    FOREIGN KEY (category_id) 
        REFERENCES categories(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
  CONSTRAINT fk_users
    FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE=INNODB;

SELECT posts.id AS post_id, posts.title, comments.id AS comment_id, comments.id_post_comment, comments.user_name_comment,comments.comment FROM comments JOIN posts ON posts.id = comments.id_post_comment WHERE posts.id = 2;
