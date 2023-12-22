package inventory

import (
	"awesomeProject/country"
	"database/sql"
	"fmt"
)

type Inventory struct {
	db *sql.DB //access to database which you created
}

// New to access Inventory struct in main
func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}

// insert data
func (i Inventory) InsertCountry(c country.Country) error { //country oziga oladi
	//database connection
	if _, err := i.db.Exec(`insert into countries values($1,$2,$3)`, c.ID, c.Name, c.Currency); err != nil {
		return err
	}
	return nil
}

// get all countries
func (i Inventory) GetAllCountries() ([]country.Country, error) {
	rows, err := i.db.Query(
		`select * from countries`,
	)
	if err != nil {
		return nil, err
	}
	cs := []country.Country{}
	for rows.Next() {
		c := country.Country{}
		if err = rows.Scan(&c.ID, &c.Name, &c.Currency); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}

// get by id
func (i Inventory) GetByIDCountry(id string) (country.Country, error) {
	c := country.Country{}
	row := i.db.QueryRow(`select * from countries where id = $1`, id)
	if err := row.Scan(&c.ID, &c.Name, &c.Currency); err != nil {
		fmt.Println("error is while selectiong by id", err)
	}
	return c, nil
}

// delete by id
func (i Inventory) DeleteByID(id string) error {
	if _, err := i.db.Exec(`delete from countries where id = $1`, id); err != nil {
		return err
	}
	return nil
}

// update by id
func (i Inventory) UpdateByID(id string) error {
	if _, err := i.db.Exec(`update countries set name = $1 where id = $2`, "saudiya", id); err != nil {
		fmt.Println("error is while updating data", err)
	}
	return nil
}
