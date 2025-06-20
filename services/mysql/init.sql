USE example_db;

-- drop
DROP TABLE IF EXISTS `users`;

-- create
CREATE TABLE IF NOT EXISTS `users` (
                                        `id`    INT AUTO_INCREMENT,
                                        `name`  VARCHAR(30) NOT NULL,
                                        `age`   INT NOT NULL,
                                        `email` VARCHAR(30),
                                        PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;