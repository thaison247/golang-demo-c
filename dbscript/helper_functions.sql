-- select employees wth extra department's information 
create or replace function get_employees_with_department (
  	p_limit int,
	p_offset int
) 
	returns table (
		employee_id int,
		full_name text,
		phone_number text,
		email text,
		gender boolean,
		job_title text,
		created_at timestamp with time zone,
		updated_at date,
		address text,
		department_id int,
		department_code varchar(5),
		department_name text
	) 
	language plpgsql
as $$
begin
	return query 
		select
			e.*, d.department_id, d.department_code, d.department_name
		from
			employees e join
			(emp_dep ed join departments d on ed.department_id = d.department_id)
			on e.employee_id = ed.employee_id
		where
			ed.id = (select emp_dep.id from emp_dep where emp_dep.employee_id = e.employee_id and emp_dep.effect_from <= now() 
					 and emp_dep.created_at = (select max(ed1.created_at)
											    from emp_dep ed1
											    where ed1.employee_id = e.employee_id))
		limit p_limit offset p_offset;
end;$$



-- select employees (return employees wth department ids)
create or replace function get_employees_with_departmentid (
  	p_limit int,
	p_offset int
) 
	returns table (
		employee_id int,
		full_name text,
		phone_number text,
		email text,
		gender boolean,
		job_title text,
		created_at date,
		updated_at date,
		address text,
		department_id int
	) 
	language plpgsql
as $$
begin
	return query 
		select
			e.*, ed.department_id
		from
			employees e join emp_dep ed on e.employee_id = ed.employee_id
		where
			ed.effect_from = (select max(effect_from) from emp_dep where emp_dep.employee_id = e.employee_id and emp_dep.effect_from <= now())
		limit p_limit offset p_offset;
end;$$


-- get employee by id (return employee with department id)
create or replace function get_employee_by_id (
	p_employee_id int
)
RETURNS table (
	employee_id int,
	full_name text,
	phone_number text,
	email text,
	gender boolean,
	job_title text,
	created_at date,
	updated_at date,
	address text,
	department_id int
)
LANGUAGE plpgsql
AS $$
BEGIN
	RETURN QUERY 
		SELECT e.*, ed.department_id
		FROM employees e, emp_dep ed
		WHERE 
            e.employee_id = p_employee_id
            AND e.employee_id = ed.employee_id 
            AND ed.effect_from = (SELECT max(effect_from) 
                                  FROM emp_dep 
                                  WHERE 
                                        emp_dep.employee_id = p_employee_id 
                                        AND emp_dep.effect_from <= now());
END;
$$

-- get one employee by id (including department infomation)
create or replace function get_one_employee_with_department (
  	p_employee_id int
) 
	returns table (
		employee_id int,
		full_name text,
		phone_number text,
		email text,
		gender boolean,
		job_title text,
		created_at timestamp with time zone,
		updated_at date,
		address text,
		department_id int,
		department_code varchar(5),
		department_name text,
		effect_from date
	) 
	language plpgsql
as $$
begin
	return query 
		select
			e.*, d.department_id, d.department_code, d.department_name, ed.effect_from
		from
			employees e join
			(emp_dep ed join departments d on ed.department_id = d.department_id)
			on e.employee_id = ed.employee_id
		where
			e.employee_id = p_employee_id and
			ed.id = (select emp_dep.id from emp_dep where emp_dep.employee_id = p_employee_id and emp_dep.effect_from <= now() 
					 and emp_dep.created_at = (select max(ed1.created_at)
											    from emp_dep ed1
											    where ed1.employee_id = p_employee_id));
end;$$

-- get employee's department id
create function get_employee_departmentid (
	p_employeeid int
)
returns integer
language plpgsql
as $$
declare department_id integer;
begin
	select ed1.department_id::integer
	into department_id
	from emp_dep ed1
	where ed1.employee_id = 1 
		  and ed1.effect_from = (select max(ed2.effect_from)
							 from emp_dep ed2
							 where ed2.employee_id = ed1.employee_id
							 	   and ed2.effect_from <= now());
							
	return department_id;
end;
$$