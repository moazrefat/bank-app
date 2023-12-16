SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

create database bankapp;
create table bankapp.user (id int not null auto_increment primary key, name varchar(255) not null,mail varchar(255),age int not null,passwd varchar(255) not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into bankapp.user (name,mail,age,passwd) values ("Moaaz Noaman", "moaz.refat@hotmail.com",35,"password");
create table bankapp.sessions (uid int,sessionid varchar(128));
create table bankapp.userdetails (uid int not null primary key, userimage varchar(64), address varchar(64), animal varchar(32), word varchar(64));
insert bankapp.userdetails(uid,userimage,address,animal,word) values (1,"moaaz.png","mainanustrasse 60, Munich","Lion","bank app admin");
commit; 