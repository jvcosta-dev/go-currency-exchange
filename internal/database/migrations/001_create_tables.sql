CREATE TABLE IF NOT EXISTS rates (
    currency TEXT PRIMARY KEY,
    rate REAL
);

CREATE TABLE  IF NOT EXISTS metadata (
    key TEXT PRIMARY KEY,
    value TEXT
);
