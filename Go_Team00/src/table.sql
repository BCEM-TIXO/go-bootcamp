CREATE TABLE IF NOT EXISTS frequency_records (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(100) NOT NULL,
    frequency DOUBLE PRECISION NOT NULL,
    timestamp BIGINT NOT NULL
);
