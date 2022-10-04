package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"strings"
)

func SetupDatabase(dbName string) *pgxpool.Pool {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "postgres://postgres:postgres@localhost:5432/" + dbName
	}
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalln("Unable to create connection pool: ", err)
	}
	return dbpool
}

func ShouldResetDatabase(db *pgxpool.Pool) {
	switch v := strings.ToLower(os.Getenv("RESET_DATABASE")); v {
	case "true", "1":
		sql := `
		DROP TABLE IF EXISTS comment;
		DROP TABLE IF EXISTS topic;
		DROP TABLE IF EXISTS avatar;
		`
		_, err := db.Exec(context.Background(), sql)
		if err != nil {
			log.Fatalln("Failed to reset database", err)
		}
		log.Println("Tables have been reset")
	}
}

func CreateDatabaseTables(db *pgxpool.Pool) {
	createTables := `
		CREATE EXTENSION IF NOT EXISTS pgcrypto;

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
		log.Fatalln("Create tables failed", err)
	}
}
