package repository

import (
	"database/sql"
	"officerent/entity"
	"strings"
	"time"
)

const insertOffice = "insert into office (id, title, description, people, price, country, state, city, district, created_at) values (?,?,?,?,?,?,?,?,?,?)"
const selectOffice = "select id, title, description, people, price, country, state, city, district, created_at, updated_at from office"
const selectOfficeByID = "select id, title, description, people, price, country, state, city, district, created_at, updated_at from office where id = UNHEX(?) LIMIT 1"

// OfficeMySQL Struct
type OfficeMySQL struct {
	db *sql.DB
}

// NewOfficeMySQL Create MySQL Office
func NewOfficeMySQL(db *sql.DB) *OfficeMySQL {
	return &OfficeMySQL{
		db: db,
	}
}

// Create an office
func (o *OfficeMySQL) Create(e *entity.Office) (entity.ID, error) {
	stmt, err := o.db.Prepare(insertOffice)
	if err != nil {
		return e.ID, err
	}

	idBinary, _ := e.ID.MarshalBinary()
	_, err = stmt.Exec(idBinary, e.Title, e.Description, e.People, e.Price, e.Country, e.State, e.City, e.District, time.Now())
	if err != nil {
		return e.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// List all offices
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
		err = rows.Scan(&office.ID, &office.Title, &office.Description, &office.People, &office.Price, &office.Country,
			&office.State, &office.City, &office.District, &office.CreatedAt, &office.UpdatedAt)
		if err != nil {
			return nil, err
		}
		offices = append(offices, &office)
	}

	return offices, nil
}

// Get an office by id.
func (o *OfficeMySQL) Get(id string) (*entity.Office, error) {
	stmt, err := o.db.Prepare(selectOfficeByID)
	if err != nil {
		return nil, err
	}

	var offices []entity.Office
	id = strings.ReplaceAll(id, "-", "")
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var office entity.Office
		err = rows.Scan(&office.ID, &office.Title, &office.Description, &office.People, &office.Price, &office.Country,
			&office.State, &office.City, &office.District, &office.CreatedAt, &office.UpdatedAt)
		if err != nil {
			return nil, err
		}
		offices = append(offices, office)
	}

	if len(offices) <= 0 {
		return nil, err
	}

	return &offices[0], nil
}
