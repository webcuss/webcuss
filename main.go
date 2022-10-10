package main

import (
	"github.com/webcuss/webcuss/db/migrate"
	"os"

	"github.com/webcuss/webcuss/db"
	"github.com/webcuss/webcuss/route"
)

func main() {
	dbConn := db.Connect("webcuss_test")
	defer dbConn.Close()

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrate.Migrate(dbConn)
		return
	}

	r := route.SetupRouter(dbConn)

	err := r.Run(":8080")
	if err != nil {
		os.Exit(1)
	}
}
