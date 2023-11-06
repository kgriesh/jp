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


-- seed data
INSERT INTO cage
    (cage_name, cage_status)
VALUES
    ('Cage One', 'ACTIVE');
INSERT INTO cage
    (cage_name, cage_status)
VALUES
    ('Cage Two', 'ACTIVE');   
INSERT INTO dinosaur
    (dino_name, dino_species, cage_id)
VALUES
    ('Maggie', 'Tyrannosaurus', 1);
INSERT INTO dinosaur
    (dino_name, dino_species, cage_id)
VALUES
    ('Lisa', 'Tyrannosaurus', 1);
INSERT INTO dinosaur
    (dino_name, dino_species, cage_id)
VALUES
    ('Bart', 'Brachiosaurus', 2);
INSERT INTO dinosaur
    (dino_name, dino_species, cage_id)
VALUES
    ('Homer', 'Stegosaurus', 2);
INSERT INTO dinosaur
    (dino_name, dino_species, cage_id)
VALUES
    ('Marge', 'Ankylosaurus', 2);    