CREATE TABLE IF NOT EXISTS USERS (
    id BIGSERIAL PRIMARY KEY,
    discord_id VARCHAR NOT NULL,
    coins INTEGER NOT NULL DEFAULT 0,
    level INTEGER NOT NULL DEFAULT 1,
    xp INTEGER NOT NULL DEFAULT 0,
    kills INTEGER NOT NULL DEFAULT 0,
    class VARCHAR NOT NULL,
    subclass VARCHAR NOT NULL,
    class_ascension VARCHAR NOT NULL DEFAULT '1',
    strength INTEGER NOT NULL DEFAULT 1,
    agility INTEGER NOT NULL DEFAULT 1,
    mana INTEGER NOT NULL DEFAULT 1,
    health INTEGER NOT NULL DEFAULT 1,
    defence INTEGER NOT NULL DEFAULT 1,
    luck INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW(),
    unique(discord_id)
);

CREATE TABLE IF NOT EXISTS FLOORS (
    id BIGSERIAL PRIMARY KEY,
    discord_id VARCHAR NOT NULL,
    floor INTEGER NOT NULL DEFAULT 1,
    exploration REAL NOT NULL DEFAULT 0,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW()
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

CREATE TABLE IF NOT EXISTS SERVERS (
    id BIGSERIAL PRIMARY KEY,
    guild_id VARCHAR NOT NULL,
    spawnchannel VARCHAR,
    unique(guild_id)
)