CREATE TABLE IF NOT EXISTS links (
    id bigSerial PRIMARY KEY,
    url varchar(30),
    short_url varchar(30) UNIQUE
);