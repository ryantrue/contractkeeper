CREATE TABLE contracts (
                           id SERIAL PRIMARY KEY,
                           contractor_id INT NOT NULL,
                           type TEXT NOT NULL,
                           number TEXT NOT NULL,
                           date DATE NOT NULL,
                           initiator TEXT NOT NULL,
                           amount DECIMAL(10, 2) NOT NULL,
                           subject TEXT NOT NULL,
                           status TEXT NOT NULL,
                           start_date DATE NOT NULL,
                           end_date DATE NOT NULL,
                           payment_procedure TEXT NOT NULL,
                           is_regular BOOLEAN NOT NULL,
                           article TEXT NOT NULL,
                           payment_accounts TEXT[] NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           FOREIGN KEY (contractor_id) REFERENCES contractors(id)
);