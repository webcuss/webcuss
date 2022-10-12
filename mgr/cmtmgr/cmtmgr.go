package cmtmgr

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
	"log"
	"net/http"
	"time"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func PostComment(c *gin.Context, dbConn *pgxpool.Pool) {
	var uri types.PostCommentUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var req types.PostCommentReq
	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// validate topicId
	if !isValidUUID(uri.TopicId) {
		c.String(http.StatusBadRequest, "Invalid topicId")
		return
	}
	sqlExists := `
	SELECT EXISTS (
		SELECT t.id
		FROM topic t
		WHERE t.id = $1
	);
	`
	var exists bool
	err = dbConn.QueryRow(context.Background(), sqlExists, uri.TopicId).Scan(&exists)
	if err != nil {
		log.Println("Failed to fetch topic, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.String(http.StatusNotFound, "Topic not found")
		return
	}

	user := c.MustGet("user").(types.Avatar)
	userId := fmt.Sprintf("%x", user.Id.Bytes)

	sqlInsert := `
	INSERT INTO comment ("topicId", "commentId", "userId", "content", "createdOn")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`
	var cmtId pgtype.UUID
	err = dbConn.QueryRow(context.Background(), sqlInsert, uri.TopicId, nil, userId, req.Comment, time.Now().UTC()).
		Scan(&cmtId)
	if err != nil {
		log.Println("Failed to insert topic, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": fmt.Sprintf("%x", cmtId.Bytes),
	})
}

func GetComments(c *gin.Context, dbConn *pgxpool.Pool) {
	var uri types.PostCommentUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// validate topicId
	if !isValidUUID(uri.TopicId) {
		c.String(http.StatusBadRequest, "Invalid topicId")
		return
	}
	sqlExists := `
	SELECT EXISTS (
		SELECT t.id
		FROM topic t
		WHERE t.id = $1
	);
	`
	var exists bool
	err = dbConn.QueryRow(context.Background(), sqlExists, uri.TopicId).Scan(&exists)
	if err != nil {
		log.Println("Failed to fetch topic, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.String(http.StatusNotFound, "Topic not found")
		return
	}

	sql := `
	SELECT c."id", c."content", c."userId", a."uname"
	FROM comment c
	INNER JOIN avatar a ON a."id" = c."userId"
	WHERE c."topicId" = $1
	ORDER BY c."createdOn" DESC;
	`

	type CommentWithUser struct {
		Id      pgtype.UUID
		Content pgtype.Text
		UserId  pgtype.UUID
		Uname   pgtype.Text
	}

	rows, err := dbConn.Query(context.Background(), sql, uri.TopicId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]gin.H, 0)
	for rows.Next() {
		var row CommentWithUser
		err = rows.Scan(&row.Id, &row.Content, &row.UserId, &row.Uname)
		if err != nil {
			log.Println("Failed to parse row, err=", err)
			continue
		}
		item := gin.H{
			"id":      fmt.Sprintf("%x", row.Id.Bytes),
			"content": row.Content.String,
			"user": gin.H{
				"id":    fmt.Sprintf("%x", row.UserId.Bytes),
				"uname": row.Uname.String,
			},
		}
		result = append(result, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"pg":   1,
		"data": result,
	})
}
