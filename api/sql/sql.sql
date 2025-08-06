CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key,
    username varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    passwd varchar(72) not null unique,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;