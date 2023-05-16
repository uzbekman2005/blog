CREATE TABLE IF NOT EXISTS about_me (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    linkedin TEXT NOT NULL DEFAULT '', 
    youtube TEXT NOT NULL DEFAULT '', 
    faceebook TEXT NOT NULL DEFAULT '', 
    telegram TEXT NOT NULL DEFAULT '', 
    resume_link TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL DEFAULT '',
    intro TEXT NOT NULL DEFAULT ''
);