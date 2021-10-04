CREATE TABLE users
(
    id       SERIAL       NOT NULL UNIQUE,
    login    VARCHAR(256) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    timezone VARCHAR(256)
);

CREATE TABLE schedule_events
(
    id         SERIAL                                      NOT NULL UNIQUE,
    user_id    INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    name       VARCHAR(256),
    time       INT,
    start_at   BIGINT,
    created_at BIGINT,
    updated_at BIGINT
);

INSERT INTO users (login, password, timezone)
VALUES ('admin', '$2a$04$kpo7WmqRJgOQ2gw4PR.AhugcwwHcm6CF9sqFGkJuCZWy/AUXcmVuu', 'Europe/Kiev'); /*password = password*/
