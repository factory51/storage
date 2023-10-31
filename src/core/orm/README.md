# Core Functionality for application logic

## Testing Setup
```
DROP DATABASE IF EXISTS gocore;
CREATE DATABASE IF NOT EXISTS gocore ; -- CREATE A TEST DB
CREATE USER 'gocore_user'@'localhost' IDENTIFIED BY 'password'; -- create a test user
GRANT ALL PRIVILEGES ON gocore.* TO 'gocore_user'@'localhost'; -- assign to test database only
FLUSH PRIVILEGES; -- reload the priviledges

USE gocore ; -- switch to using new database

DROP TABLE IF EXISTS test_example ; 
CREATE TABLE IF NOT EXISTS test_example
(
    example_id INT NOT NULL AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL,
    public_ident VARCHAR(12) NOT NULL,
    status TINYINT NOT NULL,
    PRIMARY KEY(example_id)
);

INSERT INTO test_example VALUES(NULL, 'Example Description #1', 'a1b2c3f4d5e6', 1);
INSERT INTO test_example VALUES(NULL, 'Example Description #2', 'a2b3c4f5d6e7', 1);
```