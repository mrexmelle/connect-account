
CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS credentials(
	employee_id TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,
	CONSTRAINT PK_CREDENTIALS PRIMARY KEY(employee_id)
);

INSERT INTO credentials(employee_id, password_hash, created_at, updated_at) VALUES(
	'AST-00340',
	CRYPT('abc123', GEN_SALT('bf', 8)),
	NOW(),
	NOW()
);

CREATE TABLE IF NOT EXISTS basic_profiles(
	employee_id_hash TEXT NOT NULL,
	name TEXT NOT NULL,
	dob DATE,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,
	CONSTRAINT PK_BASIC_PROFILES PRIMARY KEY(employee_id_hash)
);

CREATE TABLE IF NOT EXISTS work_profiles(
	employee_id_hash TEXT NOT NULL,
	employee_id TEXT NOT NULL,
	work_email_address TEXT,
	start_date DATE NOT NULL,
	end_date DATE,
	employment_type TEXT,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,
	CONSTRAINT PK_WORK_PROFILES PRIMARY KEY(employee_id_hash)
);
