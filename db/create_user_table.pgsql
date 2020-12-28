DROP TABLE app_user;
CREATE TABLE app_user (
    Id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL
);

CREATE TABLE sprint(
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	sprint_name VARCHAR(255) NOT NULL,
	start_date TIMESTAMP NOT NULL,
	end_date TIMESTAMP NOT NULL,
	percentage_of_completed_task INTEGER
);

CREATE TABLE project(
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	project_name VARCHAR(255) NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	status VARCHAR(255) NOT NULL,
	created_date TIMESTAMP NOT NULL,
	last_modified_date TIMESTAMP NOT NULL,
	sprint_id VARCHAR(255) REFERENCES sprint(id)
)

CREATE TABLE task(
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	task_name VARCHAR(255) NOT NULL,
	owner_id VARCHAR(255) NOT NULL,
	created_date TIMESTAMP NOT NULL,
	last_modified_date TIMESTAMP NOT NULL,
	status VARCHAR(255) NOT NULL,  
	created_by VARCHAR(255) NOT NULL,
	project_id VARCHAR(255) NOT NULL,
	estimate INTEGER,
	remaining INTEGER,
	sprint_id VARCHAR(255) REFERENCES sprint(id)
)

CREATE TABLE project(
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	project_name VARCHAR(255) NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	status VARCHAR(255) NOT NULL,
	created_date TIMESTAMP NOT NULL,
	last_modified_date TIMESTAMP NOT NULL,
	sprint_id VARCHAR(255) REFERENCES sprint(id)
)