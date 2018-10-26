---Drop the users table
DROP TABLE IF EXISTS users;

ALTER DEFAULT PRIVILEGES
    FOR USER readwrite
    IN SCHEMA public
    REVOKE ALL ON TABLES FROM readwrite;

ALTER DEFAULT PRIVILEGES
    FOR USER readonly
    IN SCHEMA public
    REVOKE ALL ON TABLES FROM readonly;

DROP USER readwrite;
DROP USER readonly;
