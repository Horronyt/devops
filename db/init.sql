CREATE TABLE mytable (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

INSERT INTO mytable (name) VALUES ('First'),('Second'),('Third');
