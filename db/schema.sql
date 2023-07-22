CREATE TABLE transactions (
	id varchar(100) PRIMARY KEY NOT NULL,
	amount BIGINT NOT NULL,
	transaction_type varchar(100) NOT NULL,
	transaction_date DATETIME NOT NULL,
  description text NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;
