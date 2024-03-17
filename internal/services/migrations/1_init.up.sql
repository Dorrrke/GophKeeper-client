CREATE TABLE IF NOT EXISTS users (
    uId INTEGER PRIMARY KEY,
    login TEXT,
    hash TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_login ON users (login);

CREATE TABLE IF NOT EXISTS cards (
    cId INTEGER PRIMARY KEY, 
    name TEXT, 
    number TEXT, 
    date TEXT, 
    cvv INTEGER,
    uId INTEGER,
    deleted INTEGER,
    last_update TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_card_name ON cards(name);

CREATE TABLE IF NOT EXISTS logins (
    lId INTEGER PRIMARY KEY, 
    name TEXT, 
    login TEXT, 
    password TEXT,
    uId INTEGER,
    deleted INTEGER,
    last_update TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_login_name ON logins (name);

CREATE TABLE IF NOT EXISTS text_data (
    tId INTEGER PRIMARY KEY, 
    name TEXT, 
    data TEXT,
    uId INTEGER,
    deleted INTEGER,
    last_update TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_text_name ON text_data (name);

CREATE TABLE IF NOT EXISTS binares_data (
    bId INTEGER PRIMARY KEY, 
    name TEXT, 
    data TEXT,
    uId INTEGER,
    deleted INTEGER,
    last_update TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_binares_name ON binares_data (name);