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