CREATE TABLE requests (
                          id SERIAL PRIMARY KEY,
                          contractor TEXT NOT NULL,
                          contract TEXT NOT NULL,
                          contract_date DATE NOT NULL,
                          subject TEXT NOT NULL,
                          amount DECIMAL(10, 2) NOT NULL,
                          contract_amount DECIMAL(10, 2) NOT NULL,
                          article TEXT NOT NULL,
                          start_date DATE NOT NULL,
                          deadline TEXT NOT NULL,
                          payment_account TEXT NOT NULL,
                          deadline_date DATE NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);