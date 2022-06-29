DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         BIGINT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    username   VARCHAR(64), 
    email      VARCHAR(256), 
    `password` VARCHAR(256), 
    salt       VARCHAR(256),
    typ        TINYINT, -- 用户类型： 0 - 普通用户；1 - 管理员；2 - 引用
    avatar     BLOB     -- 头像
);


DROP TABLE IF EXISTS documents;
CREATE TABLE documents
(
    id         BIGINT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    creator    BIGINT, 
    title      VARCHAR(256), 
    content    TEXT,
    brief      VARCHAR(516)
);


DROP TABLE IF EXISTS tags;
CREATE TABLE tags
(
    id         BIGINT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    creator    BIGINT, 
    `value`    VARCHAR(512)
);


DROP TABLE IF EXISTS doc_tags;
CREATE TABLE doc_tags
(
    id         BIGINT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    doc_id     BIGINT, 
    tag_id     BIGINT
);

