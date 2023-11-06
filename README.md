# jp

This application provides the APIs for Jurassic Park Operations.

Rules implemented in the MVP:

- All dinos must be in a cage
- Carnivores can only be in a cage with other carnivores of the same species
- Carnivores and herbivores cannot live in the same cage
- Allowed carnivor types: Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus
- Allowed herbivor types: Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops

## Endpoints

The app runs at <http://localhost:8000>

GET /dinosaurs - returns all dinosaurs in the park
GET /dinosaurs/cage/{id} - returns all dinos in a given cage
GET /dinosaurs/{id} - returns one dino matching the provided id
GET /cages = returns all cages
GET /cage/{id} - returns one cage matching the provided id
POST /dinosaur - creates a new dino and puts it in the provided cage
POST /cage - creates a new cage
PUT /dinosaur/{id} - updates a dino matching the provided id
PUT /cage/{id} - updates a cage matching the provided id

## Running locally

Requires docker running. The application with run the app and a postgres database in docker.

To build:

- Run `make build`

To run:

- Run `make start-local`s
