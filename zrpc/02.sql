CREATE TABLE IF NOT EXISTS coupon (
    id VARCHAR(36) NOT NULL,
    discount INTEGER NOT NULL,
    code VARCHAR(255) NOT NULL,
    valid_until DATE NOT NULL,
    event_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (id)

);

INSERT INTO coupon (id, discount, code, valid_until, event_id)
VALUES (
    '456e7890-f31b-25f4-c789-567834562300',
    20,
    'TECH2023',
    '2023-09-14',
    '123e4567-e89b-12d3-a456-426614174000'
);
