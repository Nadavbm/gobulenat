package datbass

var migration string = `
CREATE TABLE users (
	id serial primary key not null, 
	first_name char(50) not null, 
	last_name char(50) not null, 
	email text not null, 
	password text not null, 
	session bool
)
`
