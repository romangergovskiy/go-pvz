-- users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('client', 'moderator')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- pvz
CREATE TABLE pvz (
    id SERIAL PRIMARY KEY,
    city TEXT NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- receivings
CREATE TABLE receivings (
    id SERIAL PRIMARY KEY,
    pvz_id INTEGER NOT NULL REFERENCES pvz(id),
    status TEXT NOT NULL CHECK (status IN ('in_progress', 'close')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- items
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    receiving_id INTEGER NOT NULL REFERENCES receivings(id),
    type TEXT NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
