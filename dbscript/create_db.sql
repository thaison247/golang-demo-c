-- employee management DB
-- create employee table
CREATE TABLE IF NOT EXISTS employees (
  	employee_id SERIAL NOT NULL,
  	full_name TEXT NOT NULL,
  	phone_number TEXT,
	email TEXT NOT NULL UNIQUE,
	gender BOOLEAN DEFAULT FALSE,
	address TEXT,
	job_title TEXT,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone,
  PRIMARY KEY (employee_id)
);

INSERT INTO employees(full_name, phone_number, email, address, gender, job_title)
VALUES ('Sơn Anh Thanh', '(072)665-0203', 'dong.vo@mymailcr.com', '2697, Ấp Thập Thúc Vỹ, Xã Xuân Phi, Quận Huynh Vĩnh Phúc', TRUE, 'Developer'),
        ('Kim Hải Đạo', '(0168)520-0445', 'tran.nguy@tapiitudulu.com', '9 Phố Đới Quế Tuyết, Xã Hà Uyển Vĩnh, Quận Trân Chiến Hồ Chí Minh', TRUE, 'Developer'),
		('Chử Liên Phong', '0186 474 0976', 'psu@civoo.com', '670 Phố Yên, Xã 7, Quận Dương Thanh Bình Định', TRUE, 'Tester'),
		('Dư Phượng Sáng', '(0169)776-4470', 'ndang@triumphlotto.com', '9 Phố Tăng, Phường 1, Huyện Bảo Lâm Đồng', FALSE, 'Developer');

-- create department table
CREATE TABLE IF NOT EXISTS departments (
	department_id SERIAL NOT NULL,
	department_code VARCHAR(5) NOT NULL UNIQUE,
	department_name TEXT NOT NULL UNIQUE,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone,
	active BOOLEAN DEFAULT TRUE,
	PRIMARY KEY (department_id)
);

INSERT INTO departments(department_code, dapartment_name, active)
VALUES ('HR', 'Human Resource', TRUE),
		('MAR', 'Marketing', TRUE),
		('FIN', 'Finance', TRUE),
		('PROD', 'Production', TRUE);

-- create employee-department table
CREATE TABLE IF NOT EXISTS emp_dep (
	id SERIAL NOT NULL,
	employee_id INT NOT NULL,
	department_id INT NOT NULL,
	effect_from DATE NOT NULL DEFAULT now(),
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	PRIMARY KEY (id)
);

ALTER TABLE emp_dep
ADD CONSTRAINT empdep_employee_fk FOREIGN KEY (employee_id) REFERENCES employees(employee_id) ON DELETE CASCADE;

ALTER TABLE emp_dep
ADD CONSTRAINT empdep_department_fk FOREIGN KEY (department_id) REFERENCES departments(department_id) ON DELETE CASCADE;

INSERT INTO emp_dep (employee_id, department_id, effect_from, created_at)
VALUES (1, 3, now(), now()),
		(2, 1, now(), now()),
		(3, 4, now(), now()),
		(4, 3, now(), now());
