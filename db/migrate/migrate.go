package migrate

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/db"
	"os"
	"strings"
)

func Migrate(dbConn *pgxpool.Pool) {
	switch {
	case len(os.Args) == 2:
		db.CreateTables(dbConn)
		return
	case strings.ToLower(os.Args[2]) == "clear":
		db.ClearTables(dbConn)
		return
	}
}
