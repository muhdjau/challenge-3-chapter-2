create table books (
	id SERIAL primary key,
	title varchar(50) not null,
	author varchar(50) not null,
	description text not null
);

truncate table books restart identity;

select * from books
order by id asc;



