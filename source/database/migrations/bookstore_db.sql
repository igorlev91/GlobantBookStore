create database if not exists bookstore
GRANT ALL PRIVILEGES ON bookstore.* TO 'root'@'%' IDENTIFIED BY 'mysql';

use bookstore;

CREATE TABLE book_genres (
    name_genre VARCHAR (100) unique not null,
    genre_id INTEGER UNSIGNED NOT NULL, 
    primary key (genre_id)
);

CREATE TABLE IF NOT EXISTS book (
    book_id int  UNSIGNED NOT NULL,
    `name` varchar(100) unique NOT NULL,
    genre_id bigint NOT NULL,
    price float  default 0 NOT NULL,
    amount int default 0 NOT NULL,
    primary key (id),
    FOREIGN KEY (genre_id)  REFERENCES book_genres (genre_id) on delete cascade
);

INSERT INTO book
VALUES (1, 'The Three Musketeers', 1, 10.44, 5);

insert into book_genres (name_genre, genre_id) values
    ('Adventure', 1),
    ('Classics', 2),
    ('Fantasy', 3);
