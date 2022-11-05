package rctnmgr

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
	"github.com/webcuss/webcuss/util"
)

func PostReaction(c *gin.Context, dbConn *pgxpool.Pool) {
	user := c.MustGet("user").(types.Avatar)
	userId := fmt.Sprintf("%x", user.Id.Bytes)

	var uri types.PostReactionUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var req types.PostReactionReq
	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// validate commentId
	if !util.IsValidUUID(uri.CommentId) {
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

	// insert
	sqlInsert := `
	INSERT into reaction ("userId", "commentId", "reaction", "createdOn")
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT unique_userid_commentid_reaction
	DO UPDATE
		SET "reaction" = $3
	RETURNING id;
	`

	var reactionId pgtype.UUID
	err = dbConn.QueryRow(context.Background(), sqlInsert, userId, uri.CommentId, req.Reaction, time.Now().UTC()).
		Scan(&reactionId)
	if err != nil {
		log.Println("Failed to insert reaction, err=", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !reactionId.Valid {
		log.Println("Invalid reactionId")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": fmt.Sprintf("%x", reactionId.Bytes),
	})
}

func GetReaction(c *gin.Context, dbConn *pgxpool.Pool) {
	user := c.MustGet("user").(types.Avatar)
	userId := fmt.Sprintf("%x", user.Id.Bytes)

	var uri types.GetReactionUri
	err := c.BindUri(&uri)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	sql := `
	SELECT r."reaction", COUNT(r."id")
	FROM reaction r
	WHERE r."commentId" = $1
	GROUP BY r."reaction";
	`

	type ReactionCounts struct {
		Reaction pgtype.Numeric
		Count    pgtype.Numeric
	}

	rows, err := dbConn.Query(context.Background(), sql, uri.CommentId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resultAll := make([]gin.H, 0)
	for rows.Next() {
		var row ReactionCounts
		err = rows.Scan(&row.Reaction, &row.Count)
		if err != nil {
			log.Println("Failed to parse row, err=", err)
			continue
		}
		item := gin.H{
			"reaction": row.Reaction.Int,
			"count":    row.Count.Int,
		}
		resultAll = append(resultAll, item)
	}

	userSql := `
	SELECT DISTINCT(r."reaction")
	FROM reaction r
	WHERE r."commentId" = $1
		AND r."userId" = $2
	`
	rows, err = dbConn.Query(context.Background(), userSql, uri.CommentId, userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resultUser := make([]*big.Int, 0)
	for rows.Next() {
		var reaction pgtype.Numeric
		err = rows.Scan(&reaction)
		if err != nil || !reaction.Valid {
			log.Println("Failed to parse row, err=", err)
			continue
		}
		resultUser = append(resultUser, reaction.Int)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": resultUser,
		"all":  resultAll,
	})
}
