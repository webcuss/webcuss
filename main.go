package main

import (
	"github.com/webcuss/webcuss/db"
	"github.com/webcuss/webcuss/route"
)

func main() {
	dbConn := db.SetupDatabase("webcuss")
	defer dbConn.Close()
	db.ShouldResetDatabase(dbConn)
	db.CreateDatabaseTables(dbConn)

	r := route.SetupRouter(dbConn)

	_ = r.Run(":8080")
}
