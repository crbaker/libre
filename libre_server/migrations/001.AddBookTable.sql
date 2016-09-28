CREATE TABLE `books` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(64),
    `sub_title` VARCHAR(64),
    `published_date` VARCHAR(64),
    `description` VARCHAR(64),    
    `created` DATE CURRENT_TIMESTAMP
);