package dat

var migration string = `

DO $$ DECLARE
BEGIN
--
-- stage: if exist do not create migrations table
--
IF EXISTS(SELECT * FROM pg_tables WHERE tablename = 'migrations') THEN
	RAISE NOTICE 'migrations table exist, skipping...';
	RETURN;
END IF;

--
-- stage: migraitons table creation
--
CREATE TABLE migrations (
	name text PRIMARY KEY,
	time TIMESTAMP DEFAULT NOW()
);

END $$;

--
-- stage: users table creation unless exist
--
DO $$ BEGIN
IF EXISTS(SELECT * FROM migrations WHERE name = 'users') THEN RETURN;
END IF;

CREATE TABLE users (
	id serial primary key not null, 
	first_name char(50) not null, 
	last_name char(50) not null, 
	email text not null, 
	password text not null,
	auth_method	text,
	token	text,
	session bool
);

INSERT INTO migrations (name) VALUES ('users');
END $$;
`
