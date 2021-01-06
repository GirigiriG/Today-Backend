DROP TABLE app_user;
CREATE TABLE app_user (
	Id VARCHAR(255) PRIMARY KEY,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL,
	email VARCHAR NOT NULL
);
CREATE TABLE sprint(
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	sprint_name TEXT NOT NULL,
	start_date TIMESTAMP NOT NULL,
	end_date TIMESTAMP NOT NULL,
	percentage_of_completed_task INTEGER
);
CREATE TABLE project(
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	project_name TEXT NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	status VARCHAR(255) NOT NULL,
	created_date TIMESTAMP NOT NULL,
	last_modified_date TIMESTAMP NOT NULL,
	sprint_id VARCHAR(255) REFERENCES sprint(id)
);
CREATE TABLE task(
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	task_name TEXT NOT NULL,
	owner_id VARCHAR(255) NOT NULL,
	created_date TIMESTAMP NOT NULL,
	last_modified_date TIMESTAMP NOT NULL,
	status TEXT(50) NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	estimate INTEGER,
	remaining INTEGER,
	group_id VARCHAR(255),
	project_id VARCHAR(255) NOT NULL REFERENCES project(id) DELETE CASCADE,
	sprint_id VARCHAR(255) REFERENCES sprint(id)
) -- create test data 
-- INSERT INTO Sprint(id,sprint_name,start_date,end_date)
-- 	VALUES(777, 'sprint_x7', current_timestamp,current_timestamp);
-- INSERT INTO project(id,project_name,created_by, status, created_date, last_modified_date, sprint_id) 
-- 	VALUES(4477,'pname','001','active',current_timestamp,current_timestamp,777);
-- INSERT INTO task(task_name, id, owner_id, status, created_by, project_id, estimate, remaining, sprint_id, created_date, last_modified_date)
--   VALUES('fix progres issue', 5679, 4321, 'Active', 666, 4477, 45, 2, 777, current_timestamp, current_timestamp);