package app

type Dinosaur struct {
	Id      int64  `json:"id"`
	CageId  int64  `json:"cage_id" validate:"required"`
	Name    string `json:"dino_name" validate:"required"`
	Species string `json:"dino_species" validate:"oneof=Tyrannosaurus Velociraptor Spinosaurus Megalosaurus Brachiosaurus Stegosaurus Ankylosaurus Triceratops"`
}

type Cage struct {
	Id     int64  `json:"id"`
	Name   string `json:"cage_name" validate:"required"`
	Status string `json:"cage_status" validate:"oneof=ACTIVE DOWN"`
}
