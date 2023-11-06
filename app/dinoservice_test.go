package app

import (
	"context"
	"jp/app/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

// These tests work using a localhost db
func Test_Add_Cage_and_Dino(t *testing.T) {

	ctx := context.Background()

	asserter := assert.New(t)

	client, err := getClient()
	asserter.NoError(err)

	dinoService := NewDinoService(client)

	cage := Cage{
		Name:   "test_cage",
		Status: "ACTIVE",
	}
	err = dinoService.AddCage(ctx, cage)
	asserter.NoError(err)

	var testCageId int64
	row := client.GetConnection().QueryRowContext(ctx, "select id from cage where cage_name = 'test_cage'")
	row.Scan(&testCageId)

	dino := Dinosaur{
		CageId:  testCageId,
		Name:    "test_dino",
		Species: "Brachiosaurus",
	}

	err = dinoService.AddDino(ctx, dino)
	asserter.NoError(err)

	dinos, err := dinoService.GetDinosByCage(ctx, testCageId)
	asserter.NoError(err)
	asserter.Equal("test_dino", dinos[0].Name)
	asserter.Equal(testCageId, dinos[0].CageId)
	asserter.Equal("Brachiosaurus", dinos[0].Species)

	// clean up
	_, err = client.GetConnection().ExecContext(ctx, "DELETE from dinosaur where dino_name = 'test_dino'")
	asserter.NoError(err)

	_, err = client.GetConnection().ExecContext(ctx, "DELETE from cage where cage_name = 'test_cage'")
	asserter.NoError(err)

}

func getClient() (db.DbService, error) {
	return db.NewDbService("localhost", "postgres", "postgres", "postgres", "5432")
}
