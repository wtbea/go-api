CREATE TABLE IF NOT EXISTS characters (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    age text NOT NULL,
    horror_genre text NOT NULL,
    first_appearance text NOT NULL
);
