CREATE TABLE Comments (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    parent_id INTEGER,
    description TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id), 
    FOREIGN KEY (post_id) REFERENCES Posts(id),
    FOREIGN KEY (parent_id) REFERENCES Comments(id)
);
