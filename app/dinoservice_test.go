package app

import (
	"context"
	"jp/app/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddCage(t *testing.T) {

	ctx := context.Background()

	asserter := assert.New(t)

	client, err := getClient()
	asserter.NoError(err)

	dinoService := NewDinoService(client)

	cage := Cage{
		Name:   "test cage",
		Status: "ACTIVE",
	}
	err = dinoService.AddCage(ctx, cage)
	asserter.NoError(err)

}

func getClient() (db.DbService, error) {

	return db.NewDbService("postgres", "postgres", "postgres", "5432")
}
