CREATE TABLE `migrations` (
    `scriptname` VARCHAR(64) PRIMARY KEY,
    `created` DATE CURRENT_TIMESTAMP
);

CREATE TABLE `books` (
    `title` VARCHAR(64),
    `created` DATE CURRENT_TIMESTAMP
);