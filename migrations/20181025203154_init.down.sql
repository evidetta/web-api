--Drop Index
DROP INDEX updated_at_idx;

---Drop the users table
DROP TABLE IF EXISTS users;

DROP FUNCTION generateTag ();

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    REVOKE ALL ON TABLES FROM readwrite;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    REVOKE ALL ON TABLES FROM readonly;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    REVOKE ALL ON SEQUENCES FROM readwrite;

ALTER DEFAULT PRIVILEGES
    FOR ROLE postgres
    IN SCHEMA public
    REVOKE ALL ON SEQUENCES FROM readonly;

DROP USER readwrite;
DROP USER readonly;
