create database if not exists book_store

use book_store;


CREATE TABLE IF NOT EXISTS book (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    genre_id INT NOT NULL,
    price DECIMAL(10),
    amount INT DEFAULT 0 NOT NULL,
    primary key (book_id),
    FOREIGN KEY (genre_id)  REFERENCES genre (genre_id)
);

CREATE TABLE IF NOT EXISTS genre(
    genre_id INT PRIMARY KEY AUTO_INCREMENT,
    name_genre VARCHAR (30)
)