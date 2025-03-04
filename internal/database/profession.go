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
const sqlFindProfessionAll = "SELECT id, description FROM profession"
const sqlFindProfessionById = "SELECT id, description FROM profession WEHRE id = $1"

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

func (p *Profession) FindAll() ([]Profession, error) {
	rows, err := p.db.Query(sqlFindProfessionAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	profession := []Profession{}
	for rows.Next() {
		var id, description string
		if err := rows.Scan(&id, &description); err != nil {
			return nil, err
		}
		profession = append(profession, Profession{ID: id, Description: description})
	}

	return profession, nil
}

func (p *Profession) Find(idProfession string) (Profession, error) {
	var id, description string
	err := p.db.QueryRow(sqlFindProfessionById, idProfession).Scan(&id, &description)
	if err != nil {
		return Profession{}, nil
	}
	return Profession{ID: id, Description: description}, nil
}
