package main

import (
	"log"
	"os"
	"strings"

	"github.com/webcuss/webcuss/db/migrate"

	"github.com/webcuss/webcuss/db"
	"github.com/webcuss/webcuss/route"
)

func main() {
	dbConn := db.Connect()
	defer dbConn.Close()

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrate.Migrate(dbConn)
		return
	} else {
		appSecret := os.Getenv("APP_SECRET")
		if len(strings.TrimSpace(appSecret)) < 1 {
			log.Fatalln("APP_SECRET environment variable is missing")
		}
	}

	r := route.SetupRouter(dbConn)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln("Application cannot start, err=", err)
	}
}
