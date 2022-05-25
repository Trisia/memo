DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         INTEGER PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    username   VARCHAR(64), 
    email      VARCHAR(256), 
    `password` VARCHAR(256), 
    salt       VARCHAR(256)
);


DROP TABLE IF EXISTS documents;
CREATE TABLE documents
(
    id         INTEGER PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    creator    INT, 
    title      VARCHAR(256), 
    content    TEXT
);


DROP TABLE IF EXISTS tags;
CREATE TABLE tags
(
    id         INTEGER PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    creator    INT, 
    `value`    VARCHAR(512)
);


DROP TABLE IF EXISTS doc_tags;
CREATE TABLE doc_tags
(
    id         INTEGER PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    doc_id     INT, 
    tag_id     INT
);

