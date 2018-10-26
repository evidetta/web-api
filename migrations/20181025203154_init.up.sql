---Create two database users
CREATE USER readwrite LOGIN ENCRYPTED PASSWORD 'password';
CREATE USER readonly LOGIN ENCRYPTED PASSWORD 'password';

---Grant default privileges
ALTER DEFAULT PRIVILEGES
    FOR USER readwrite
    IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE ON TABLES TO readwrite;

ALTER DEFAULT PRIVILEGES
    FOR USER readonly
    IN SCHEMA public
    GRANT SELECT ON TABLES TO readonly;

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
