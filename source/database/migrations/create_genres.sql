use bookstore;


CREATE TABLE book_genres (
    name_genre VARCHAR (100) not null,
    genre_id INTEGER UNSIGNED NOT NULL, 
    primary key (genre_id)

);

insert into book_genres (name_genre, genre_id) values
    ('Adventure', 1),
    ('Classics', 2),
    ('Fantasy', 3);

