CREATE TRIGGER enforce_table_size_limit
    BEFORE INSERT ON clientes
    FOR EACH ROW
    EXECUTE FUNCTION check_table_size();
