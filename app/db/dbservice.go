package db

import (
	"database/sql"
	"fmt"
	"jp/app/model"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

type DbService interface {
	GetDinos() (model.Dinosaurs, error)
	GetConnection() *sql.DB
}

type Database struct {
	Conn *sql.DB
}

func (db Database) GetConnection() *sql.DB {
	return db.Conn
}

func NewDbService(username, password, database string) (DbService, error) {
	db := Database{}
	ds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", ds)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func (db Database) GetDinos() (model.Dinosaurs, error) {
	dinos := model.Dinosaurs{}
	rows, err := db.Conn.Query("SELECT * FROM dinosaur ORDER BY ID DESC")
	if err != nil {
		return dinos, err
	}
	for rows.Next() {
		var dino model.Dinosaur
		err := rows.Scan(&dino.ID, &dino.Name)
		if err != nil {
			return dinos, err
		}
		dinos.Dinosaur = append(dinos.Dinosaur, dino)
	}
	return dinos, nil

}
