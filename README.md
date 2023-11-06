# jp

This application provides the APIs for Jurassic Park Operations.

Rules implemented in the MVP:

- All dinos must be in a cage
- Carnivores can only be in a cage with other carnivores of the same species
- Carnivores and herbivores cannot live in the same cage
- Allowed carnivor types: Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus
- Allowed herbivor types: Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops

## Notable items missing

- Many more tests are needed
- Possibly refactor to separate http logic from the handler and separate business logic from the db queries
- Implement other items in the Bonus Points section of the requirements
- Security - endpoinst are currently unsecured

## Concurrent Enviroment Considerations

If this app was required to support high volume and concurrency, it would be important to handle scalability and load balancing in the production environment. From the code side, using threads where needed, using a database connection pool and data caching would go a long way. It would also be important to load test the application to test real world scenarios and see what's needed.

## Endpoints

The app runs at <http://localhost:8000/v1>

GET /dinosaurs - returns all dinosaurs in the park
GET /dinosaurs/cage/{id} - returns all dinos for a given cageId
GET /dinosaurs/{id} - returns one dino matching the provided id
GET /cages = returns all cages
GET /cage/{id} - returns one cage matching the provided id
PUT /dinosaur/{id} - updates a dino name and/or cage (changing the cage_id will move the dino, if allowed)
PUT /cage/{id} - updates cage attributes for the matching cageId
POST /dinosaur - creates a new dino and puts it in the provided cage
    - example:
        {
            "cage_id": 1,
            "dino_name": "Brac",
            "dino_species": "Brachiosaurus"
        }
POST /cage - creates a new cage
    - example:
        {
            "cage_name": "Cage One",
            "cage_status": "ACTIVE"
        }

## Running locally

Requires docker running, the application runs the app and a postgres database in docker. There is currently seed data added during start up to provide example data.

To build:

- Run `make build`

To run:

- Run `make start-local`s
