---Grant default privileges
ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE ON TABLES TO readwrite;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    GRANT SELECT ON TABLES TO readonly;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    GRANT USAGE, SELECT ON SEQUENCES TO readwrite;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    GRANT SELECT ON SEQUENCES TO readonly;

--Create random key generative function.
CREATE OR REPLACE FUNCTION generateTag ()
RETURNS TEXT AS $tag$
DECLARE
  tag TEXT;
BEGIN
  SELECT md5(concat(cast(now() AS TEXT), 'PrJOXyyNeX')) INTO tag;
  RETURN tag;
END;
$tag$ LANGUAGE PLPGSQL;

---Create the users table
CREATE TABLE users(
  id              SERIAL,
  tag             VARCHAR(32) UNIQUE DEFAULT generateTag(),
  name            VARCHAR(100) NOT NULL,
  address         VARCHAR(200) NOT NULL,
  date_of_birth   DATE NOT NULL,
  created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at      TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at      TIMESTAMP WITH TIME ZONE
);

--Create index on updated_at in the user table
CREATE INDEX updated_at_idx ON users (updated_at);
