CREATE TABLE contractors (
                             id SERIAL PRIMARY KEY,
                             name TEXT NOT NULL,
                             inn TEXT NOT NULL,
                             ogrn TEXT NOT NULL,
                             requisites TEXT NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);