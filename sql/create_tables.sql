CREATE TABLE IF NOT EXISTS dinosaur (
    id BIGSERIAL PRIMARY KEY,
    dino_name text NOT NULL,
    dino_species text NOT NULL,
    cage_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS cage (
    id BIGSERIAL PRIMARY KEY,
    cage_name text NOT NULL,
    cage_status text NOT NULL,
    UNIQUE ("cage_name" )
);

ALTER TABLE dinosaur ADD FOREIGN KEY ("cage_id") REFERENCES cage ("id");