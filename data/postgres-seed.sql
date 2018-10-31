create table pgUser (
	id serial not null primary key,
	username varchar,
	email varchar,
	phone varchar,
	password text
)


insert into pgUser(username, email, phone, password) values ('test', 'test@pg.com', '123455678', 'password')