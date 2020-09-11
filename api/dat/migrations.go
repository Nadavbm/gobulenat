package dat

var migration string = `
DP $$ DECLARE
BEGIN
--
--
IF EXIST(SELECT 1 FROM pg_tables WHERE tablename = 'migrations') THEN
	RAISE NOTICE 'migrations table exist, skipping...';
	RETURN;
END IF;

CREATE TABLE migrations (
	name text PRIMARY_KEY,
	time	TIMESTAMP DEFAULT NOW()
);

END $$

DO $$ BEGIN
IF EXIST(SELECT 1 FROM migrations WHERE name = 'users') THEN RETURN;
END IF;

CREATE TABLE users (
	id serial primary key not null, 
	first_name char(50) not null, 
	last_name char(50) not null, 
	email text not null, 
	password text not null,
	auth_method	text nol null,
	token	text,
	session bool
);

INSERT INTO migrations (name) VALUES ('users');
END $$
`
