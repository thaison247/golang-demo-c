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
		address text,
		job_title text,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
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
											   where ed1.employee_id = emp_dep.employee_id
											  		  and ed1.effect_from <= now()))
		limit p_limit offset p_offset;
end;$$

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
		address text,
		job_title text,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
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
											   where ed1.employee_id = emp_dep.employee_id
											  		  and ed1.effect_from <= now()));
end;$$