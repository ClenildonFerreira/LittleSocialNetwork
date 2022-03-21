CREATE DATABASE IF NOT EXISTS devsocialnetwork;
USE devsocialnetwork;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name_user varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password_user varchar(100) not null,
    createdIn timestamp default current_timestamp()
) ENGINE=INNODB;