-- POSTGRES

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    pass TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    room_id UUID NOT NULL REFERENCES rooms(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);