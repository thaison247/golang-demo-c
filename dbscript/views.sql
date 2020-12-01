-- employees with their's current department infomation
create view current_empdep_view
as 
select
	ed.employee_id, ed.department_id, d.department_code, d.department_name
from
	employees e join
	(emp_dep ed join departments d on ed.department_id = d.department_id)
	on e.employee_id = ed.employee_id
where
	ed.id = (select emp_dep.id from emp_dep where emp_dep.employee_id = e.employee_id and emp_dep.effect_from <= now() 
			 and emp_dep.created_at = (select max(ed1.created_at)
									   from emp_dep ed1
									   where ed1.employee_id = emp_dep.employee_id
											  and ed1.effect_from <= now()));



----
CREATE VIEW employee_department_view
AS
(SELECT emp_dep.*, 'ACTIVE' AS status
FROM emp_dep JOIN (
	SELECT employee_id, MAX(effect_from) AS effect_from
	FROM emp_dep
	WHERE effect_from <= NOW()::DATE
	GROUP BY employee_id
) B ON emp_dep.employee_id = B.employee_id AND emp_dep.effect_from=B.effect_from
UNION
SELECT emp_dep.*, 'INACTIVE' AS status
FROM emp_dep JOIN (
	SELECT employee_id, MIN(effect_from) AS effect_from
	FROM emp_dep
	GROUP BY employee_id
	HAVING MIN(effect_from) > NOW()::DATE
) B ON emp_dep.employee_id = B.employee_id AND emp_dep.effect_from=B.effect_from)
