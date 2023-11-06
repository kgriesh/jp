package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
)

type DbService interface {
	GetConnection() *sql.DB
}

type Database struct {
	Conn *sql.DB
}

func (db Database) GetConnection() *sql.DB {
	return db.Conn
}

func NewDbService(username, password, database, port string) (DbService, error) {
	db := Database{}
	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, port, username, password, database)
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
