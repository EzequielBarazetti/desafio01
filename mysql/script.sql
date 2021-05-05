create database db_mysql;

use db_mysql;

create table cursos (id int AUTO_INCREMENT, nome varchar(200), PRIMARY KEY(id));

insert into cursos (nome) values ('Docker');
insert into cursos (nome) values ('Golang');
insert into cursos (nome) values ('Reack');
