CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id),
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    closed_at TIMESTAMP,
    is_completed BOOLEAN DEFAULT FALSE
);
