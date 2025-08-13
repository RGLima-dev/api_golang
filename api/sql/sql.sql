CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key,
    username varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    passwd varchar(72) not null,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE followers (

    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE


    follower_id int not null,
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE 

    primary key(user_id,follower_id)
)ENGINE=INNODB;