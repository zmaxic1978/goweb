CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE authors
(
    id          serial PRIMARY KEY,
    firstname   varchar(255) not null,
    lastname    varchar(255) not null,
    description varchar(255) not null,
    birthday    date null
);

CREATE TABLE books
(
    id       serial PRIMARY KEY,
    name     varchar(255) not null,
    authorid int REFERENCES authors(id) ON DELETE CASCADE NOT NULL,
    year     int          not null,
    isbn     varchar(255) not null
);

INSERT INTO authors (firstname, lastname, description, birthday) VALUES ('Агата', 'Кристи', 'детективщица!', '1890-09-15');
INSERT INTO authors (firstname, lastname, description, birthday) VALUES ('Артур', 'Конан Дойл', 'детективщик', '1859-05-22');
INSERT INTO authors (firstname, lastname, description, birthday) VALUES ('Стивен', 'Кинг', 'автор ужастиков', '1947-09-21');
INSERT INTO authors (firstname, lastname, description, birthday) VALUES ('Айзек', 'Азимов', 'современный писатель-фантаст', '1919-04-06');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Убийство в доме викария', 1, 1954, '123-343-5465');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Смерть на Ниле', 1, 1951, '123-343-5466');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Собака Баскервиллей', 2, 1904, '56-343-5465');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Этюд в багровых тонах', 2, 1905, '56-343-5466');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Оно', 3, 1986, '985-13-6269-7');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Мертвая зона', 3, 1991, '991-15-4544-2');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Я робот', 4, 1950, '985-13-6269-7');
INSERT INTO books (name, authorid, year, isbn) VALUES ('Сами боги', 4, 1972, '978-5-04-172716-1');