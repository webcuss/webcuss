package tpcmgr

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func createOrGetTopic(dbConn *pgxpool.Pool, userId, urlString, title string) (string, bool, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", false, err
	}

	scheme := u.Scheme
	hostname := u.Hostname()
	path := u.Path
	querySlice := strings.Split(u.RawQuery, "&")
	sort.Strings(querySlice)
	querySorted := strings.Join(querySlice, "&")

	sqlFind := `
	SELECT t.Id
	FROM topic t
	WHERE t.hostname = $1
		AND t.path = $2
		AND t."querySearch" @@ to_tsquery('english', $3);
	`
	var tpcId pgtype.UUID
	err = dbConn.QueryRow(context.Background(), sqlFind, hostname, path, strings.Join(querySlice, " & ")).
		Scan(&tpcId)
	if err != nil {
		// no row
		// insert
		sqlInsert := `
		INSERT INTO topic
		    (
		     "scheme",
		     "hostname",
		     "path",
		     "query",
		     "querySearch",
		     "title",
		     "likes",
		     "createdOn",
		     "userId"
		    )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
		`
		err = dbConn.QueryRow(
			context.Background(), sqlInsert, scheme,
			hostname, path, querySorted, strings.Join(querySlice, " "),
			title, 0, time.Now().UTC(), userId,
		).Scan(&tpcId)
		if err != nil {
			return "", true, err
		}
		return fmt.Sprintf("%x", tpcId.Bytes), true, nil
	} else {
		// already created
		return fmt.Sprintf("%x", tpcId.Bytes), false, nil
	}
}

func addComment(dbConn *pgxpool.Pool, userId, tpcId, comment string) (string, error) {
	sql := `
	INSERT INTO comment ("topicId", "userId", "content", "createdOn")
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var cmtId pgtype.UUID
	err := dbConn.QueryRow(context.Background(), sql, tpcId, userId, comment, time.Now().UTC()).
		Scan(&cmtId)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", cmtId.Bytes), nil
}

func PostTopic(c *gin.Context, dbConn *pgxpool.Pool) {
	user := c.MustGet("user").(types.Avatar)
	userId := fmt.Sprintf("%x", user.Id.Bytes)

	var req types.PostTopicReq
	err := c.Bind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	tpcId, tpcCreated, err := createOrGetTopic(dbConn, userId, req.Url, req.Title)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	res := gin.H{"id": tpcId}

	// handle comment
	if cmt := req.Comment; tpcCreated && cmt != "" {
		cmtId, err := addComment(dbConn, userId, tpcId, cmt)
		if err != nil {
			res["commentId"] = err.Error()
		} else {
			res["commentId"] = cmtId
		}
	}
	c.JSON(http.StatusCreated, res)
}

func GetTopic(c *gin.Context, dbConn *pgxpool.Pool) {
	sql := `
	SELECT t."id",
		MIN(t."hostname") AS "hostname",
		MIN(t."path") AS "path",
		MIN(t."query") AS "query",
		MIN(t."userId"::text)::uuid AS "userId",
		MIN(a."uname") AS "uname",
		COUNT(c."id") AS "commentsCount",
		MIN(t."likes") AS "likes",
		MIN(t."title") AS "title"
	FROM topic t
	INNER JOIN avatar a ON a."id" = t."userId"
	LEFT JOIN comment c ON c."topicId" = t."id"
	GROUP BY t."id"
	ORDER BY "commentsCount" DESC NULLS LAST;
	`

	type TopicWithAvatar struct {
		Id            pgtype.UUID
		Hostname      pgtype.Text
		Path          pgtype.Text
		Query         pgtype.Text
		UserId        pgtype.UUID
		Uname         pgtype.Text
		CommentsCount pgtype.Numeric
		Likes         pgtype.Numeric
		Title         pgtype.Text
	}

	scanned := make([]TopicWithAvatar, 0)
	rows, err := dbConn.Query(context.Background(), sql)
	if err != nil {
		log.Println("Failed to fetch topics, err=", err)
		c.JSON(http.StatusOK, gin.H{
			"data": make([]interface{}, 0),
		})
		return
	}
	for rows.Next() {
		var row TopicWithAvatar
		err = rows.Scan(
			&row.Id, &row.Hostname, &row.Path, &row.Query,
			&row.UserId, &row.Uname, &row.CommentsCount,
			&row.Likes, &row.Title,
		)
		if err != nil {
			log.Println("Error scanning topic row, err=", err)
			continue
		}
		scanned = append(scanned, row)
	}

	result := make([]gin.H, 0)
	for _, v := range scanned {
		m := gin.H{
			"id":            fmt.Sprintf("%x", v.Id.Bytes),
			"hostname":      v.Hostname.String,
			"path":          v.Path.String,
			"query":         v.Query.String,
			"commentsCount": v.CommentsCount.Int,
			"likes":         v.Likes.Int,
			"title":         v.Title.String,
			"user": gin.H{
				"id":    fmt.Sprintf("%x", v.UserId.Bytes),
				"uname": v.Uname.String,
			},
		}
		result = append(result, m)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
