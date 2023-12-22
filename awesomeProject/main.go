package main

import (
	"awesomeProject/inventory"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	inv := inventory.New(db) //new inventory qaytaradi

	//c := country.Country{ //insert qilinshi kere bogan data
	//	ID:       uuid.New(),
	//	Name:     "china",
	//	Currency: "yang",
	//}
	//if err = inv.InsertCountry(c); err != nil {
	//	panic(err)
	//}
	//fmt.Println("data successfully added")

	//update data in main
	idU := "7e24db9d-e0c6-4828-9d49-40647da4ded0"
	if err = inv.UpdateByID(idU); err != nil {
		fmt.Println("error is while updating data in main function", err)
	}
	fmt.Println("data updated")
	//countries, err := inv.GetAllCountries()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(countries)

	//get by id
	//id := "8693bd5c-7bc5-4494-b151-7c61da3452f9"
	//country, err := inv.GetByIDCountry(id)
	//if err != nil {
	//	fmt.Println("error in main.go while using getbyid method")
	//}
	//fmt.Println(country)

	//deleting by id
	//if err = inv.DeleteByID(id); err != nil {
	//	fmt.Println("error is while deleting data in main.go", err)
	//}
	//fmt.Println("data deleted successfully")

}
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port = 5432 password = 1218 database = country sslmode = disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
