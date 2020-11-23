-- create trigger function 
CREATE FUNCTION Department_update_trigger_function ()
	RETURNS TRIGGER
	LANGUAGE PLPGSQL
AS 
$$
BEGIN 
	IF NEW <> OLD THEN
		UPDATE departments
		SET updated_at = now()
		WHERE department_id = NEW.department_id;
	END IF;
	
	RETURN NEW;
END;
$$;

-- create trigger after update on departments table
CREATE TRIGGER Department_update_trigger
AFTER UPDATE
ON departments
FOR EACH ROW
	EXECUTE FUNCTION Department_update_trigger_function();