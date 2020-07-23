DROP DATABASE IF EXISTS `fx_datastore`;
CREATE DATABASE `fx_datastore` default character set utf8mb4;
GRANT ALL ON fx_datastore.* TO 'apluser'@'%' IDENTIFIED BY 'apluser';
FLUSH PRIVILEGES;
