-- +goose Up
CREATE TABLE IF NOT EXISTS requests (
                                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        contractor TEXT,
                                        contract TEXT,
                                        contract_date TEXT,
                                        subject TEXT,
                                        amount REAL,
                                        contract_amount REAL,
                                        article TEXT,
                                        start_date TEXT,
                                        deadline TEXT,
                                        payment_account TEXT,
                                        deadline_date TEXT,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS requests;