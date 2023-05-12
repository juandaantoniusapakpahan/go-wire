create database belajar_golang_result_api;

create table belajar_golang_result_api.category (
    id integer primary key auto_increment,
    name varchar(200) not null
)engine=InnoDB;

alter table belajar_golang_result_api.category auto_increment=1000;

insert into belajar_golang_result_api.category(name) values("Juanda");

select * from belajar_golang_result_api.category;

