create database if not exists officerent;
use officerent;

create table if not exists country 
(
    id int not null primary key auto_increment,
    country varchar(100)
);

create table if not exists state 
(
    id int not null primary key auto_increment,
    country_id int foreign key references country(id),
    state varchar(100),
);

create table if not exists city 
(
    id int not null primary key auto_increment,
    state_id int foreign key references state(id),
    city varchar(100)
);

create table if not exists district 
(
    id int not null primary key auto_increment,
    city_id int foreign key references city(id),
    district varchar(100)
);

create table if not exists office
(
    id binary(16) not null primary key,
    title varchar(100),
    description varchar(255),
    people integer,
    price decimal(10, 2),
    country_id int foreign key references country(id),
    state_id int foreign key references state(id),
    city_id int foreign key references city(id),
    district_id int foreign key references district(id),
    created_at datetime,
    updated_at datetime
);

insert into country values (1, 'Brasil');

insert into state values (1, 1, 'São Paulo');

insert into city values (1, 1, 'São Paulo');

insert into district values (1, 1, 'Jardim Patente');