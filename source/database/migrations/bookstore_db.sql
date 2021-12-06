create database if not exists bookstore

use bookstore;

CREATE TABLE book_genres (
    name_genre VARCHAR (100) not null,
    genre_id INTEGER UNSIGNED NOT NULL, 
    primary key (genre_id)
);

CREATE TABLE IF NOT EXISTS book (
    book_id bigint(20)   UNSIGNED NOT NULL,
    `name` varchar(100) unique NOT NULL,
    genre_id bigint NOT NULL,
    price  float NOT NULL,
    amount bigint NOT NULL,
    FOREIGN KEY (genre_id)  REFERENCES book_genres (genre_id)
);

INSERT INTO Book
VALUES (1, 'The Three Musketeers', 1, 10.44, 5);

insert into book_genres (name_genre, genre_id) values
    ('Adventure', 1),
    ('Classics', 2),
    ('Fantasy', 3);

