package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Person struct {
	db   *sql.DB
	ID   string
	Name string
	Year int32
}

const sqlInsertPerson = "INSERT INTO person (id, name, year) VALUES ($1,$2,$3)"

func NewPerson(db *sql.DB) *Person {
	return &Person{db: db}
}

func (p *Person) Create(name string, year int32) (Person, error) {
	id := uuid.New().String()
	_, err := p.db.Exec(sqlInsertPerson, id, name, year)

	if err != nil {
		return Person{}, err
	}

	return Person{ID: id, Name: name, Year: year}, nil
}
