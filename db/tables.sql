CREATE TABLE IF NOT EXISTS USERS (
    id BIGSERIAL PRIMARY KEY,
    discord_id VARCHAR,
    coins INTEGER,
    class VARCHAR,
    subclass VARCHAR,
    class_ascension INTEGER,
    weapon VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW(),
    unique(discord_id));