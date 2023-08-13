create table customer
(
	id   varchar(100) not null,
	name varchar(100) not null,
	primary key(id)
) engine = InnoDb;

delete from customer 

select * from customer

alter table customer 
	add column email varchar(100),
	add column balance integer default 0,
	add column rating double  default 0.0,
	add column created_at timestamp default CURRENT_TIMESTAMP,
	add column birth_date DATE,
	add column married boolean default false;

DESC customer 

INSERT INTO customer(id, name, email, balance, rating, birth_date, married) 
VALUES 
	('nanda', 'Nanda', 'nanda@email.com', 10000000, 90.0, '1999-10-10', true),
	('afif', 'Afif', 'afif@email.com', 80000000, 93.0, '2010-10-10', true),
	('joko', 'Joko', 'Joek@email.com', 53000000, 53.0, '2012-10-10', false),
	('joni', 'Joni', 'jonnnn@email.com', 23000000, 78.0, '1995-10-10', false);


