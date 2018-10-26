---Create two database users
CREATE ROLE readwrite LOGIN ENCRYPTED PASSWORD 'password';
CREATE ROLE readonly LOGIN ENCRYPTED PASSWORD 'password';

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

---Create the users table
CREATE TABLE users(
  id              SERIAL,
  name            VARCHAR(100) NOT NULL,
  address         VARCHAR(200) NOT NULL,
  date_of_birth   DATE NOT NULL,
  created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at      TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at      TIMESTAMP WITH TIME ZONE
);
