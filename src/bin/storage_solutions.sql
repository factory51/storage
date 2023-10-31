

CREATE DATABASE storage_solution;


CREATE TABLE storage_bucket
(
    `created_at` DATETIME NOT NULL,
    client_ident VARCHAR(36) NOT NULL,
    `key` VARCHAR(255) NOT NULL,
    `value` TEXT CHARSET utf8 NOT NULL,
    PRIMARY KEY(client_ident, `key`),
    INDEX(`key`)

);

