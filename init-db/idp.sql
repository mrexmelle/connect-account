CREATE DATABASE idp;
CREATE USER idp WITH PASSWORD 'idp123';
GRANT CONNECT ON DATABASE idp TO idp;

\c idp;

CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS credentials(
	employee_id TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ DEFAULT NULL,
	CONSTRAINT pk_credentials PRIMARY KEY(employee_id)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON credentials TO idp;

CREATE TABLE IF NOT EXISTS profiles(
	ehid TEXT NOT NULL,
	employee_id TEXT NOT NULL,
	email_address TEXT NOT NULL,
	name TEXT NOT NULL,
	dob DATE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT pk_profiles PRIMARY KEY(ehid)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON profiles TO idp;

CREATE TABLE IF NOT EXISTS organizations(
	id TEXT NOT NULL,
	hierarchy TEXT NOT NULL,
	name TEXT NOT NULL,
	lead_ehid TEXT,
	email_address TEXT,
	private_slack_channel TEXT,
	public_slack_channel TEXT,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ DEFAULT NULL,
	CONSTRAINT pk_organizations PRIMARY KEY(id)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON organizations TO idp;

CREATE TABLE IF NOT EXISTS tenures(
	id SERIAL,
	ehid TEXT NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE DEFAULT NULL,
	employment_type TEXT,
	organization_id TEXT NOT NULL,
	title_grade TEXT DEFAULT '',
	title_name TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT pk_tenures PRIMARY KEY(id),
	CONSTRAINT fk1_tenures FOREIGN KEY(ehid) REFERENCES profiles(ehid),
	CONSTRAINT fk2_tenures FOREIGN KEY(organization_id) REFERENCES organizations(id),
	CONSTRAINT uk_tenures UNIQUE(ehid, start_date)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON tenures TO idp;
GRANT USAGE ON SEQUENCE tenures_id_seq TO idp;

CREATE TABLE IF NOT EXISTS titles(
	id SERIAL,
	grade TEXT NOT NULL,
	career_type TEXT NOT NULL,
	domain TEXT NOT NULL, 
	name TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT pk_grades PRIMARY KEY(id)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON titles TO idp;
GRANT USAGE ON SEQUENCE titles_id_seq TO idp;
