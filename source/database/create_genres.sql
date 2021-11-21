use book_store;

CREATE TABLE genre(
    genre_id INT PRIMARY KEY,
    name_genre VARCHAR (30) 
)

insert into genre (name_genre, genre_id) values
    ('Adventure', 1),
    ('Classics', 2),
    ('Fantasy', 3);
