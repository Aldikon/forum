CREATE TABLE Users (
    id INTEGER PRIMARY KEY,
    user_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    UNIQUE (user_name),
    UNIQUE (email)
);
