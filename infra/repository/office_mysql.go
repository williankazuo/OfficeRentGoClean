package repository

import (
	"database/sql"
	"officerent/entity"
	"time"
)

const insertOffice = "insert into office (id, title, description, people, price, created_at) values (?,?,?,?,?,?)"
const selectOffice = "select id, title, description, people, price, created_at, updated_at from office"

type OfficeMySQL struct {
	db *sql.DB
}

func NewOfficeMySQL(db *sql.DB) *OfficeMySQL {
	return &OfficeMySQL{
		db: db,
	}
}

func (o *OfficeMySQL) Create(e *entity.Office) (entity.ID, error) {
	stmt, err := o.db.Prepare(insertOffice)
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(e.ID, e.Title, e.Description, e.People, e.Price, time.Now())
	if err != nil {
		return e.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (o *OfficeMySQL) List() ([]*entity.Office, error) {
	stmt, err := o.db.Prepare(selectOffice)
	if err != nil {
		return nil, err
	}

	var offices []*entity.Office
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var office entity.Office
		err = rows.Scan(&office.ID, &office.Title, &office.Description, &office.People, &office.Price, &office.CreatedAt, &office.UpdatedAt)
		if err != nil {
			return nil, err
		}
		offices = append(offices, &office)
	}

	return offices, nil
}
