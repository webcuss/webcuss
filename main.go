package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/sup", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"token": "",
		})
	})

	r.POST("/sin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	})

	r.POST("/sout", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}

func setupDatabase() *pgxpool.Pool {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "postgres://postgres:postgres@localhost:5432/webcuss"
	}
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func createDatabaseTables(db *pgxpool.Pool) {
	createTables := `
		CREATE TABLE IF NOT EXISTS avatar (
			"id" UUID NOT NULL DEFAULT gen_random_uuid(),
			"uname" VARCHAR(255) NOT NULL,
			"pword" VARCHAR(255) NOT NULL,
			"createdOn" TIMESTAMP NOT NULL,
			"pebbles" INTEGER NOT NULL,
			"verifiedOn" TIMESTAMP NULL,
			"email" VARCHAR(255) NULL,
			PRIMARY KEY (id)
		);

		-- indexes
		CREATE INDEX IF NOT EXISTS "idx_avatar_uname" ON avatar (
			"uname" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_avatar_createdOn" ON avatar (
			"createdOn" DESC
		);
		CREATE INDEX IF NOT EXISTS "idx_avatar_pebbles" ON avatar (
			"pebbles" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_avatar_verifiedOn" ON avatar (
			"verifiedOn" ASC NULLS LAST
		);

		CREATE TABLE IF NOT EXISTS topic (
			"id" UUID NOT NULL DEFAULT gen_random_uuid(),
			"url" TEXT NOT NULL,
			"search" TEXT NOT NULL,
			"createdOn" TIMESTAMP NOT NULL,
			"userId" UUID NOT NULL,
			PRIMARY KEY (id),
			CONSTRAINT fk_user
				FOREIGN KEY ("userId")
					REFERENCES avatar("id")
					ON DELETE CASCADE
		);

		-- indexes
		CREATE INDEX IF NOT EXISTS "idx_topic_url" ON topic (
			"url" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_topic_createdOn" ON topic (
			"createdOn" DESC
		);
		CREATE INDEX IF NOT EXISTS "idx_topic_userId" ON topic (
			"userId" ASC
		);

		CREATE TABLE IF NOT EXISTS comment (
			"id" UUID NOT NULL DEFAULT gen_random_uuid(),
			"topicId" UUID NOT NULL,
			"commentId" UUID NULL,
			"userId" UUID NOT NULL,
			"content" TEXT NOT NULL,
			"createdOn" TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			CONSTRAINT fk_topic
				FOREIGN KEY ("topicId")
					REFERENCES topic("id")
					ON DELETE CASCADE,
			CONSTRAINT fk_user
				FOREIGN KEY ("userId")
					REFERENCES avatar("id")
					ON DELETE CASCADE,
			CONSTRAINT fk_comment
				FOREIGN KEY ("commentId")
					REFERENCES comment("id")
					ON DELETE CASCADE
		);

		-- indexes
		CREATE INDEX IF NOT EXISTS "idx_comment_topicId" ON comment (
			"topicId" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_comment_commentId" ON comment (
			"commentId" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_comment_userId" ON comment (
			"userId" ASC
		);
		CREATE INDEX IF NOT EXISTS "idx_comment_createdOn" ON comment (
			"userId" DESC
		);
		`
	_, err := db.Exec(context.Background(), createTables)
	if err != nil {
		fmt.Printf("Create tables failed. %v", err)
		os.Exit(1)
	}
}

func main() {
	db := setupDatabase()
	defer db.Close()
	createDatabaseTables(db)

	r := setupRouter()
	r.Run(":8080")
}
