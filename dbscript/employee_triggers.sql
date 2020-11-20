-- create trigger function
CREATE FUNCTION Log_update_trigger_function()
	RETURNS TRIGGER
	LANGUAGE PLPGSQL
AS 
$$
BEGIN 
	IF NEW <> OLD THEN
		UPDATE employees
		SET updated_at = now()
		WHERE employee_id = NEW.employee_id;
	END IF;
	
	RETURN NEW;
END;
$$;

-- create trigger, bind to employees table
CREATE TRIGGER Update_trigger 
AFTER UPDATE
ON employees
FOR EACH ROW
    EXECUTE PROCEDURE Log_update_trigger_function();
	