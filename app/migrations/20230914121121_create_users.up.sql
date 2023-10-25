CREATE TABLE "user" (
    id bigserial not null primary key,
    name varchar not null,
    surname varchar not null,
    email varchar not null,
    encrypted_password varchar not null,
    patronymic varchar,
    age int not null,
    gender varchar not null,
    nationality varchar not null
);