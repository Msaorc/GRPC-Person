package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Profession struct {
	db          *sql.DB
	ID          string
	Description string
}

const sqlInsertProfession = "INSERT INTO profession (id, description) VALUES ($1,$2)"

func NewProfession(db *sql.DB) *Profession {
	return &Profession{db: db}
}

func (p *Profession) Create(description string) (Profession, error) {
	id := uuid.New().String()
	_, err := p.db.Exec(sqlInsertProfession, id, description)

	if err != nil {
		return Profession{}, err
	}

	return Profession{ID: id, Description: description}, nil
}
