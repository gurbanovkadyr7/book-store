create table users (
    id serial primary key,
    first_name varchar(10),
    last_name varchar(15),
    email varchar(50),
    phone varchar(15),
    hash_pass varchar(300)
);

create table books(
    id serial primary key,
    title varchar(50),
    author varchar(20),
    price float not null,
    about varchar(200),
    category varchar(200) references categories (id),
    book_language varchar(10) references languages(id)
);

create table languages(
    id serial primary key,
    code varchar(5) not null
);

create table categories(
    id serial primary key,
    title varchar(255) not null
);

create table categories_books(
	category_id int,
	book_id int,

	constraint book_id 
      foreign key (book_id) 
        references books (id),		

	constraint category_id 
      foreign key (category_id) 
        references categories (id)
);




