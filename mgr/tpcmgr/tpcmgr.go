package tpcmgr

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func createOrGetTopic(dbConn *pgxpool.Pool, userId, urlString string) (string, bool, error) {
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
		     "createdOn",
		     "userId"
		    )
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
		`
		err = dbConn.QueryRow(context.Background(), sqlInsert, scheme, hostname, path, querySorted, strings.Join(querySlice, " "), time.Now().UTC(), userId).
			Scan(&tpcId)
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

	tpcId, tpcCreated, err := createOrGetTopic(dbConn, userId, req.Url)
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
