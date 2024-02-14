CREATE
OR REPLACE FUNCTION check_table_size()
RETURNS TRIGGER AS $$
DECLARE
table_size INT;
    max_table_size
INT := 5;
BEGIN
-- Execute the query to get the table size and store it in table_size variable
EXECUTE 'SELECT COUNT(*) FROM ' || TG_RELNAME::regclass INTO table_size;
IF
table_size >= max_table_size THEN
        RAISE EXCEPTION 'Table size limit reached';
END IF;
RETURN NEW;
END;
$$
LANGUAGE plpgsql;
