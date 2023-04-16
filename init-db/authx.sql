
CREATE TABLE IF NOT EXISTS credentials(
	id VARCHAR(16) NOT NULL,
	password_hash BYTEA NOT NULL,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,
	CONSTRAINT PK PRIMARY KEY(id)
);

INSERT INTO credentials(id, password_hash, created_at, updated_at) VALUES('AST-00340', SHA256('abc123'), NOW(), NOW());