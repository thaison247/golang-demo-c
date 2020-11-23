-- Table Definition
CREATE TABLE "public"."users" (
    "id" varchar(20) NOT NULL,
    "name" varchar(50),
    "email" varchar(255),
    "phone" varchar(20),
    PRIMARY KEY ("id")
);

-- employee management DB
-- create employee table
CREATE TABLE IF NOT EXISTS employees (
  	employee_id SERIAL NOT NULL,
  	full_name TEXT NOT NULL,
  	phone_number TEXT,
	email TEXT NOT NULL,
	gender BOOLEAN,
	job_title TEXT,
	created_at DATE,
	updated_at DATE,
  PRIMARY KEY (employee_id)
);

INSERT INTO employees(full_name, phone_number, email, address, gender, job_title, created_at)
VALUES ('Sơn Anh Thanh', '(072)665-0203', 'dong.vo@mymailcr.com', '2697, Ấp Thập Thúc Vỹ, Xã Xuân Phi, Quận Huynh Vĩnh Phúc', TRUE, 'Developer', now())
        ('Kim Hải Đạo', '(0168)520-0445', 'tran.nguy@tapiitudulu.com', '9 Phố Đới Quế Tuyết, Xã Hà Uyển Vĩnh, Quận Trân Chiến Hồ Chí Minh', TRUE, 'Developer', now()),
		('Chử Liên Phong', '0186 474 0976', 'psu@civoo.com', '670 Phố Yên, Xã 7, Quận Dương Thanh Bình Định', TRUE, 'Tester', now()),
		('Dư Phượng Sáng', '(0169)776-4470', 'ndang@triumphlotto.com', '9 Phố Tăng, Phường 1, Huyện Bảo Lâm Đồng', FALSE, 'Developer', now())

-- create department table
CREATE TABLE IF NOT EXISTS departments (
	department_id SERIAL NOT NULL,
	department_code VARCHAR(5),
	dapartment_name TEXT NOT NULL,
	created_at DATE,
	updated_at DATE,
	active BOOLEAN,
	PRIMARY KEY (department_id)
)

INSERT INTO departments(department_code, dapartment_name, created_at, active)
VALUES ('HR', 'Human Resource', now(), TRUE),
		('MAR', 'Marketing', now(), TRUE),
		('FIN', 'Finance', now(), TRUE),
		('PROD', 'Production', now(), TRUE);

-- create employee-department table
CREATE TABLE IF NOT EXISTS emp_dep (
	id SERIAL NOT NULL,
	employee_id INT NOT NULL,
	department_id INT NOT NULL,
	effect_from DATE NOT NULL,
	created_at DATE NOT NULL,
	PRIMARY KEY (id)
)

ALTER TABLE emp_dep
ADD CONSTRAINT empdep_employee_fk FOREIGN KEY (employee_id) REFERENCES employees(employee_id) ON DELETE CASCADE

ALTER TABLE emp_dep
ADD CONSTRAINT empdep_department_fk FOREIGN KEY (department_id) REFERENCES departments(department_id) ON DELETE CASCADE

INSERT INTO emp_dep (employee_id, department_id, effect_from, created_at)
VALUES (1, 3, now(), now()),
		(2, 1, now(), now()),
		(3, 4, now(), now()),
		(4, 3, now(), now());
