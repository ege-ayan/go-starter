CREATE TABLE IF NOT EXISTS greetings (
    id         SERIAL PRIMARY KEY,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO greetings (name) VALUES ('World'), ('Go');
