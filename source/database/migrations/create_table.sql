create database if not exists book_store

use book_store;

CREATE TABLE IF NOT EXISTS book (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(100) unique NOT NULL,
    genre_id bigint DEFAULT 0 NOT NULL,
    price  float NOT NULL,
    amount bigint DEFAULT 0 NOT NULL,
    primary key (book_id),
    FOREIGN KEY (genre_id)  REFERENCES book_genres (genre_id)
);

INSERT INTO Book
VALUES (1, 'The Three Musketeers', 1, 10.44, 5);
