CREATE TABLE IF NOT EXISTS address (
    id VARCHAR(36) NOT NULL,
    uf VARCHAR(2) NOT NULL,
    city VARCHAR(255) NOT NULL,
    event_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (id)

);

INSERT INTO address (id, uf, city, event_id)
VALUES (
    '789e0123-g62c-36h7-i890-678912345600',
    'SP',
    'São Paulo',
    '123e4567-e89b-12d3-a456-426614174000'
);
