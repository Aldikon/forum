CREATE TABLE Posts (
    id INTEGER PRIMARY KEY,
    create_att TEXT  NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
    -- CHECK (user_id != 0)
);
