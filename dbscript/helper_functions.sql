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


-- get one employee and current department (left join with 'current_empdep_view')
create or replace function get_one_employee_with_department_v2 (
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
		updated_at timestamp with time zone,
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
			e.*, cev.department_id, cev.department_code, cev.department_name, cev.effect_from
		from
			employees e left join current_empdep_view cev on e.employee_id = cev.employee_id
		where e.employee_id = p_employee_id;
end;$$

--
create or replace function get_employees_with_department_v3 (
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
		SELECT E.*, D.department_id, D.department_code, D.department_name
		FROM (employees E LEFT JOIN employee_department_view ED ON E.employee_id = ED.employee_id)
			 JOIN departments D ON ED.department_id = D.department_id
		LIMIT p_limit OFFSET p_offset;
end;$$


-- get one employee and current department Version 3
create or replace function get_one_employee_with_department_v3 (
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
		SELECT E.*, D.department_id, D.department_code, D.department_name, ED.effect_from
		FROM (employees E JOIN employee_department_view ED ON E.employee_id = ED.employee_id)
			 JOIN departments D ON ED.department_id = D.department_id
		WHERE E.employee_id = p_employee_id;
end;$$


--- SonNH

SELECT *, 'ACTIVE' AS status
FROM emp_dep JOIN (
	SELECT employee_id, MAX(effect_from) AS effect_from
	FROM emp_dep
	WHERE effect_from <= NOW()::DATE
	GROUP BY employee_id
) B ON emp_dep.employee_id = B.employee_id AND emp_dep.effect_from=B.effect_from
UNION
SELECT *, 'INACTIVE' AS status
FROM emp_dep JOIN (
	SELECT employee_id, MIN(effect_from) AS effect_from
	FROM emp_dep
	GROUP BY employee_id
	HAVING MIN(effect_from) > NOW()::DATE
) B ON emp_dep.employee_id = B.employee_id AND emp_dep.effect_from=B.effect_from
