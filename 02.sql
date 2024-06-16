USE eventostec;

CREATE TABLE IF NOT EXISTS event (
    id CHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    imgUrl VARCHAR(2048) NOT NULL,
    eventUrl VARCHAR(2048) NOT NULL,
    remote BOOLEAN NOT NULL,
    date DATE NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO event (id, title, description, imgUrl, eventUrl, remote, date)
VALUES (
    '123e4567-e89b-12d3-a456-426614174000', 'Conferência de Tecnologia',
    'Participe da nossa conferência anual de tecnologia, com palestrantes do mundo todo discutindo as últimas tendências em desenvolvimento de software, segurança e IA.',
    'https://example.com/img/event1.jpg',
    'https://example.com/events/tech2023',
    TRUE,
    '2023-09-15'
);