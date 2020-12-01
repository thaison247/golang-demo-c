-- add employee procedure
create or replace procedure add_employee(
	p_full_name IN employees.full_name%TYPE,
	p_phone_number IN employees.phone_number%TYPE,
	p_email IN employees.email%TYPE,
	p_gender IN employees.gender%TYPE,
	p_job_title IN employees.job_title%TYPE,
	p_address IN employees.address%TYPE,
	p_department_id IN departments.department_id%TYPE,
	p_effect_from IN emp_dep.effect_from%TYPE
)
language plpgsql    
as $$
declare
	new_employee_id int;
begin
    -- insert employee 
    insert into employees(full_name, phone_number, email, gender, job_title, address)
	values (p_full_name, p_phone_number, p_email, p_gender, p_job_title, p_address)
	returning employee_id into new_employee_id;
	
	-- insert employee-department record
	insert into emp_dep(employee_id, department_id, effect_from)
	values (new_employee_id, p_department_id, p_effect_from);    

    commit;
end;$$