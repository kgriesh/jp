package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jp/app/db"
	"slices"
	"strings"

	validate "github.com/go-playground/validator/v10"
)

type DinoService interface {
	GetDinos(ctx context.Context) ([]Dinosaur, error)
	GetDinoById(ctx context.Context, dinoId int64) (Dinosaur, error)
	GetCageById(ctx context.Context, cageId int64) (Cage, error)
	GetDinosByCage(ctx context.Context, cageId int64) ([]Dinosaur, error)
	AddDino(ctx context.Context, dino Dinosaur) error
	UpdateDino(ctx context.Context, dino Dinosaur) error
	GetCages(ctx context.Context) ([]Cage, error)
	AddCage(ctx context.Context, cage Cage) error
	UpdateCage(ctx context.Context, cage Cage) error
}

type dinoServiceImpl struct {
	dbService db.DbService
}

// NewDinoService return a new DinoService
func NewDinoService(db db.DbService) dinoServiceImpl {
	return dinoServiceImpl{
		dbService: db,
	}
}

// GetDinos get all dinos regardless of cage
func (s dinoServiceImpl) GetDinos(ctx context.Context) ([]Dinosaur, error) {
	dinos := []Dinosaur{}
	rows, err := s.
		dbService.
		GetConnection().
		QueryContext(ctx, "SELECT id, dino_name, dino_species, cage_id FROM dinosaur ORDER BY ID ASC")
	if err != nil {
		return dinos, err
	}
	for rows.Next() {
		var dino Dinosaur
		err := rows.Scan(&dino.Id, &dino.Name, &dino.Species, &dino.CageId)
		if err != nil {
			return dinos, err
		}
		dinos = append(dinos, dino)
	}
	return dinos, nil
}

// GetDinoById get a cage by id
func (s dinoServiceImpl) GetDinoById(ctx context.Context, dinoId int64) (Dinosaur, error) {
	dino := Dinosaur{}
	row := s.dbService.GetConnection().QueryRowContext(ctx, "SELECT id, dino_name, cage_id, dino_species FROM dinosaur where id=$1", dinoId)
	err := row.Scan(&dino.Id, &dino.Name, &dino.CageId, &dino.Species)
	if err != nil {
		return dino, err
	}
	return dino, nil
}

// GetCageById get a cage by id
func (s dinoServiceImpl) GetCageById(ctx context.Context, cageId int64) (Cage, error) {
	cage := Cage{}
	row := s.dbService.GetConnection().QueryRowContext(ctx, "SELECT id, cage_name, cage_status FROM cage where id=$1", cageId)
	err := row.Scan(&cage.Id, &cage.Name, &cage.Status)
	if err != nil {
		return cage, err
	}
	return cage, nil
}

// GetDinosByCage get all dinos in a cage
func (s dinoServiceImpl) GetDinosByCage(ctx context.Context, cageId int64) ([]Dinosaur, error) {
	dinos := []Dinosaur{}
	rows, err := s.
		dbService.
		GetConnection().
		QueryContext(ctx, "SELECT id, dino_name, dino_species, cage_id FROM dinosaur where cage_id = $1", cageId)
	if err != nil {
		return dinos, err
	}
	for rows.Next() {
		var dino Dinosaur
		err := rows.Scan(&dino.Id, &dino.Name, &dino.Species, &dino.CageId)
		if err != nil {
			return dinos, err
		}
		dinos = append(dinos, dino)
	}
	if len(dinos) == 0 {
		return dinos, sql.ErrNoRows
	}
	return dinos, nil

}

// AddDino add a new dinosaur
func (s dinoServiceImpl) AddDino(ctx context.Context, dino Dinosaur) error {

	v := validate.New()
	err := v.Struct(dino)
	if err != nil {
		var errString strings.Builder
		vErrors := err.(validate.ValidationErrors)
		for _, validationError := range vErrors {
			errString.WriteString(fmt.Sprintf("Invalid entry for %s. ", validationError.Field()))
		}
		return &ServiceRequestError{
			err:      err.Error(),
			response: errString.String(),
		}
	}

	// get the exising dinos in the cage and check if new dino is allowed
	existingDinos, err := s.GetDinosByCage(ctx, dino.CageId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if dinoIsAllowed(dino, existingDinos) {
		_, err = s.dbService.
			GetConnection().
			ExecContext(ctx, "INSERT INTO dinosaur ( dino_name, dino_species, cage_id) VALUES ($1, $2, $3)", dino.Name, dino.Species, dino.CageId)
		if err != nil {
			return err
		}
	} else {
		return &ServiceRequestError{
			err:      "error adding dinosaur to cage",
			response: "This dinosaur is not allowed to be put in this cage",
		}
	}

	return nil
}

// UpdateDino updates a dinosaur
func (s dinoServiceImpl) UpdateDino(ctx context.Context, dino Dinosaur) error {

	v := validate.New()
	err := v.Struct(dino)
	if err != nil {
		var errString strings.Builder
		vErrors := err.(validate.ValidationErrors)
		for _, validationError := range vErrors {
			errString.WriteString(fmt.Sprintf("Invalid entry for %s. ", validationError.Field()))
		}
		return &ServiceRequestError{
			err:      err.Error(),
			response: errString.String(),
		}
	}

	// get the exising dinos for the cage_id and check if the dino is allowed
	existingDinos, err := s.GetDinosByCage(ctx, dino.CageId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if dinoIsAllowed(dino, existingDinos) {
		query := `UPDATE dinosaur set 
		dino_name = $1,
		cage_id = $2
		where id = $3`

		_, err = s.
			dbService.
			GetConnection().
			ExecContext(ctx, query, dino.Name, dino.CageId, dino.Id)
		if err != nil {
			return err
		}
	} else {
		return &ServiceRequestError{
			err:      "error adding dinosaur to cage",
			response: "This dinosaur is not allowed to be put in this cage",
		}
	}

	return nil
}

// AddCage add a new cage
func (s dinoServiceImpl) AddCage(ctx context.Context, cage Cage) error {

	v := validate.New()
	err := v.Struct(cage)
	if err != nil {
		var errString strings.Builder
		vErrors := err.(validate.ValidationErrors)
		for _, validationError := range vErrors {
			errString.WriteString(fmt.Sprintf("Invalid entry for %s. ", validationError.Field()))
		}
		return &ServiceRequestError{
			err:      err.Error(),
			response: errString.String(),
		}
	}

	_, err = s.
		dbService.
		GetConnection().
		ExecContext(ctx, "INSERT INTO cage ( cage_name, cage_status) VALUES ($1, $2)", cage.Name, cage.Status)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCage updates a cage
func (s dinoServiceImpl) UpdateCage(ctx context.Context, cage Cage) error {

	v := validate.New()
	err := v.Struct(cage)
	if err != nil {
		var errString strings.Builder
		vErrors := err.(validate.ValidationErrors)
		for _, validationError := range vErrors {
			errString.WriteString(fmt.Sprintf("Invalid entry for %s. ", validationError.Field()))
		}
		return &ServiceRequestError{
			err:      err.Error(),
			response: errString.String(),
		}
	}

	query := `UPDATE cage set 
		cage_status = $1,
		cage_name = $2
		where id = $3`

	_, err = s.dbService.GetConnection().ExecContext(ctx, query, cage.Status, cage.Name, cage.Id)
	if err != nil {
		return err
	}

	return nil
}

// GetCages get all cages
func (s dinoServiceImpl) GetCages(ctx context.Context) ([]Cage, error) {
	cages := []Cage{}
	rows, err := s.
		dbService.
		GetConnection().
		QueryContext(ctx, "SELECT id, cage_name, cage_status FROM cage")
	if err != nil {
		return cages, err
	}
	for rows.Next() {
		var cage Cage
		err := rows.Scan(&cage.Id, &cage.Name, &cage.Status)
		if err != nil {
			return cages, err
		}
		cages = append(cages, cage)
	}
	return cages, nil
}

/*
dinoIsAllowed rules:
- carnivores can only be in same cage as same species
- herbivores cannot be in same cage as carnivores
*/
func dinoIsAllowed(newDino Dinosaur, currentDinos []Dinosaur) bool {

	if len(currentDinos) == 0 {
		return true
	}

	carnivores := []string{"Tyrannosaurus", "Velociraptor", "Spinosaurus", "Megalosaurus"}

	newDinoIsCarn := false
	currentDinosAreCarn := false

	if slices.Contains(carnivores, newDino.Species) {
		newDinoIsCarn = true
	}

	// we can compare the new dino to just one example in the cage
	existingDino := currentDinos[0]
	if slices.Contains(carnivores, existingDino.Species) {
		currentDinosAreCarn = true
	}

	if newDinoIsCarn && currentDinosAreCarn {
		return newDino.Species == existingDino.Species
	} else if !newDinoIsCarn && !currentDinosAreCarn {
		return true
	} else {
		return false
	}
}

type ServiceRequestError struct {
	err      string
	response string
}

func (s ServiceRequestError) Error() string {
	return s.err
}
