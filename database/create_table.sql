create database if not exists book_store

use book_store;


CREATE TABLE IF NOT EXISTS book (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(100) unique NOT NULL,
    genre_id bigint DEFAULT 0 NOT NULL,
    price  DOUBLE NOT NULL,
    amount bigint DEFAULT 0 NOT NULL,
    primary key (book_id),
    FOREIGN KEY (genre_id)  REFERENCES genre (genre_id)
);
