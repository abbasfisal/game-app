-- +migrate Up
CREATE TABLE `access_controls`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `actor_id`   VARCHAR(191) NOT NULL,
    `actor_type` ENUM('role' , 'user') NOT NULL,
    `permission_id` INT NOT NULL ,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`)
);

-- +migrate Down
Drop Table `access_controls`;