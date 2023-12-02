package main

import "gps-data-receiver/database"

func main() {
	db, err := database.NewGPSDatabaseConnection("mongodb://root:password@localhost:27017")
	if err != nil {
		panic(err)
	}

	defer db.Disconnect()

}
