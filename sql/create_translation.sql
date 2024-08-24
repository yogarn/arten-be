CREATE TABLE translations (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    origin_language VARCHAR(255) NOT NULL,
    target_language VARCHAR(255) NOT NULL,
    word TEXT NOT NULL,
    translate VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
