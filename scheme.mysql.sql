-- DB scheme for the testrest, mysql-specific
-- MySQL 5.5 or newer required

-- DROP DATABASE IF EXISTS testrest;
-- CREATE DATABASE testrest;

-- USE testrest;

-- DROP TABLE IF EXISTS account;
CREATE TABLE account (
	id     BIGINT PRIMARY KEY AUTO_INCREMENT,
	name   VARCHAR(20) NOT NULL,
	amount BIGINT NOT NULL,
	ctime  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	mtime  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) COMMENT = "Accounts are stored here" ENGINE = InnoDB;

CREATE UNIQUE INDEX account_name ON account (name);

-- Disallow non-empty account deletion
DROP TRIGGER IF EXISTS account_delete;
--
DELIMITER //
CREATE TRIGGER account_delete 
BEFORE DELETE ON account FOR EACH ROW
BEGIN
  IF OLD.amount <> 0 THEN
    SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Account is not empty';
  END IF;
END;//
DELIMITER ;

-- Disallow negative amounts
DROP TRIGGER IF EXISTS account_update;
--
DELIMITER //
CREATE TRIGGER account_update 
BEFORE UPDATE ON account FOR EACH ROW
BEGIN
  IF NEW.amount < 0 THEN
    SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Negative amount not allowed';
  END IF;
END;//
DELIMITER ;

-- Disallow negative amounts
DROP TRIGGER IF EXISTS account_insert;
--
DELIMITER //
CREATE TRIGGER account_insert 
BEFORE INSERT ON account FOR EACH ROW
BEGIN
  IF NEW.amount < 0 THEN
    SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Negative amount not allowed';
  END IF;
END;//
DELIMITER ;
