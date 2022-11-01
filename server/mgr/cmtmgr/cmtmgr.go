package cmtmgr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func PostComment(c *gin.Context, dbConn *pgxpool.Pool) {
	var uri types.TopicCommentUri
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
	var uri types.TopicCommentUri
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
	SELECT c."id", c."content", c."userId", a."uname", c."createdOn"
	FROM comment c
	INNER JOIN avatar a ON a."id" = c."userId"
	WHERE c."topicId" = $1
	ORDER BY c."createdOn" DESC;
	`

	type CommentWithUser struct {
		Id        pgtype.UUID
		Content   pgtype.Text
		UserId    pgtype.UUID
		Uname     pgtype.Text
		CreatedOn pgtype.Timestamp
	}

	rows, err := dbConn.Query(context.Background(), sql, uri.TopicId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]gin.H, 0)
	for rows.Next() {
		var row CommentWithUser
		err = rows.Scan(&row.Id, &row.Content, &row.UserId, &row.Uname, &row.CreatedOn)
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
		if row.CreatedOn.Valid {
			item["createdOn"] = row.CreatedOn.Time.Format(time.RFC3339)
		}
		result = append(result, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func PostReply(c *gin.Context, dbConn *pgxpool.Pool) {
	var uri types.CommentReplyUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var req types.PostReplyReq
	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// validate commentId
	if !isValidUUID(uri.CommentId) {
		c.String(http.StatusBadRequest, "Invalid commentId")
		return
	}
	sqlExists := `
	SELECT EXISTS (
		SELECT c.id
		FROM comment c
		WHERE c.id = $1
	);
	`
	var exists bool
	err = dbConn.QueryRow(context.Background(), sqlExists, uri.CommentId).Scan(&exists)
	if err != nil {
		log.Println("Failed to fetch comment, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.String(http.StatusNotFound, "Comment not found")
		return
	}

	user := c.MustGet("user").(types.Avatar)
	userId := fmt.Sprintf("%x", user.Id.Bytes)

	sqlInsert := `
	INSERT INTO comment ("topicId", "commentId", "userId", "content", "createdOn")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	var commentId pgtype.UUID
	err = dbConn.QueryRow(context.Background(), sqlInsert, nil, uri.CommentId, userId, req.Comment, time.Now().UTC()).
		Scan(&commentId)
	if err != nil {
		log.Println("Failed to insert reply, err=", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": fmt.Sprintf("%x", commentId.Bytes),
	})
}

func GetReplies(c *gin.Context, dbConn *pgxpool.Pool) {
	var uri types.CommentReplyUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// validate commentId
	if !isValidUUID(uri.CommentId) {
		c.String(http.StatusBadRequest, "Invalid commentId")
		return
	}
	sqlExists := `
	SELECT EXISTS (
		SELECT c.id
		FROM comment c
		WHERE c.id = $1
	);
	`
	var exists bool
	err = dbConn.QueryRow(context.Background(), sqlExists, uri.CommentId).Scan(&exists)
	if err != nil {
		log.Println("Failed to fetch comment, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.String(http.StatusNotFound, "Comment not found")
		return
	}

	sql := `
	SELECT c."id",
	c."content",
	c."createdOn",
	c."userId",
	a."uname"
	FROM comment c
	INNER JOIN avatar a ON a."id" = c."userId"
	WHERE c."commentId" = $1;
	`

	type CommentWithUser struct {
		Id        pgtype.UUID
		Content   pgtype.Text
		CreatedOn pgtype.Timestamp
		UserId    pgtype.UUID
		Uname     pgtype.Text
	}

	result := make([]gin.H, 0)
	rows, err := dbConn.Query(context.Background(), sql, uri.CommentId)
	if err != nil {
		log.Println("Failed to fetch comments, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		var row CommentWithUser
		err = rows.Scan(&row.Id, &row.Content, &row.CreatedOn, &row.UserId, &row.Uname)
		if err != nil {
			log.Println("Failed to parse row")
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
		if row.CreatedOn.Valid {
			item["createdOn"] = row.CreatedOn.Time.Format(time.RFC3339)
		}
		result = append(result, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
