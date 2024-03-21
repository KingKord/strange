CREATE TABLE IF NOT EXISTS schedule
(
    id          SERIAL PRIMARY KEY,
    user_id     INT          NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    date_from   TIMESTAMP    NOT NULL,
    date_to     TIMESTAMP    NOT NULL
);