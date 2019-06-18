create table book
(
	id varchar(36) not null
		primary key,
	title varchar(255) not null,
	synopsis text null,
	isbn10 varchar(15) null,
	isbn13 varchar(15) null,
	category varchar(100) not null,
	language varchar(30) not null,
	publisher varchar(255) not null,
	edition varchar(50) null,
	soldamount int null,
	currentamount int null,
	paperbackprice float null,
	ebookprice float null,
	createdtime datetime not null,
	modifiedtime datetime not null,
	version int not null,
	constraint book_isbn10_uindex
		unique (isbn10),
	constraint book_isbn13_uindex
		unique (isbn13)
);

create index book_id_index
	on book (id);

create index book_publisher_index
	on book (publisher);

create index book_title_index
	on book (title);

create table review
(
	id varchar(36) not null,
	score int not null,
	description text null,
	book_id varchar(36) null,
	createdtime datetime not null,
	modifiedtime datetime not null,
	version int not null,
	constraint review_pk
		primary key (id),
	constraint review_book_id_fk
		foreign key (book_id) references book (id)
			on delete cascade
);
