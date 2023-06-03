CREATE TABLE IF NOT EXISTS USERS (
    id BIGSERIAL PRIMARY KEY,
    discord_id VARCHAR NOT NULL,
    coins INTEGER NOT NULL DEFAULT 0 CHECK(coins >= 0),
    class VARCHAR NOT NULL,
    subclass VARCHAR NOT NULL,
    class_ascension VARCHAR NOT NULL DEFAULT '1',
    created_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW(),
    unique(discord_id)
);

CREATE TABLE IF NOT EXISTS INVENTORIES (
    id BIGSERIAL PRIMARY KEY,
    discord_id VARCHAR NOT NULL,
    item_name VARCHAR NOT NULL,
    item_type VARCHAR NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    durability INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW()
);