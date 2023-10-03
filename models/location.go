package models

import "api.legatodesigns.com/database"

type Country struct {
	ID   int   
	Name string
}

func CountryList() []Country {
	db := database.DBConnect()
	selDB,err := db.Query("SELECT * FROM country")

	country := Country{}
	response :=[] Country{}
	if err !=nil {
		panic("Db get error"+ err.Error())
	}
	defer selDB.Close()
	for selDB.Next() {
		var id int
		var name string
		err := selDB.Scan(&id,&name)
		if err != nil {
			panic("Db scan error"+ err.Error())
		}

		country.ID = id
		country.Name = name

		response = append(response, country)
	}
	return response
	// defer db.Close();

}