CREATE TABLE
if NOT EXISTS dinosaur
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL
);

--seed
INSERT INTO dinosaur
    (name)
VALUES
    ('dino one');
INSERT INTO dinosaur
    (name)
VALUES
    ('dino two');
INSERT INTO dinosaur
    (name)
VALUES
    ('dino three');