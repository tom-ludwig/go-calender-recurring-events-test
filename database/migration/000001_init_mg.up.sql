CREATE TABLE events (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    rrule TEXT, -- Stores the recurrence rule (optional, NULL for non-recurring events)
    exrule TEXT, -- Stores the exception rule (optional, NULL if no exceptions)
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    events REFERENCES events(id)
);

