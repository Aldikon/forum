CREATE TABLE Session_User (
    user_id INTEGER NOT NULL,
    session_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (session_id) REFERENCES Session(id)
);  
