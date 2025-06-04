CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE urls (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      url TEXT NOT NULL,
                      count INT DEFAULT 0
);
