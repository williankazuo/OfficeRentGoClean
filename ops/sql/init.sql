create database if not exists officerent;
use officerent;

create table if not exists office
(
    id binary(16) not null primary key,
    title varchar(100),
    description varchar(255),
    people integer,
    price decimal(10, 2),
    country varchar(50),
    state varchar(50),
    city varchar(50),
    district varchar(50),
    created_at datetime,
    updated_at datetime
);