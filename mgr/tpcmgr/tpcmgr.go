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
		     "createdOn",
		     "userId"
		    )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
		`
		err = dbConn.QueryRow(context.Background(), sqlInsert, scheme, hostname, path, querySorted, strings.Join(querySlice, " "), title, time.Now().UTC(), userId).
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
		t."scheme",
		t."hostname",
		t."path",
		t."query",
		t."createdOn",
		t."userId",
		a."uname",
		a."createdOn" AS "avatarCreatedOn",
		a."pebbles",
		a."verifiedOn" AS "avatarVerifiedOn",
		a."email"
	FROM topic t
	INNER JOIN avatar a ON a."id" = t."userId"
	ORDER BY t."createdOn" DESC;
	`

	type TopicWithAvatar struct {
		Id               pgtype.UUID
		Scheme           pgtype.Text
		Hostname         pgtype.Text
		Path             pgtype.Text
		Query            pgtype.Text
		CreatedOn        pgtype.Timestamp
		UserId           pgtype.UUID
		Uname            pgtype.Text
		AvatarCreatedOn  pgtype.Timestamp
		Pebbles          pgtype.Numeric
		AvatarVerifiedOn pgtype.Timestamp
		Email            pgtype.Text
	}

	scanned := make([]TopicWithAvatar, 0)
	rows, err := dbConn.Query(context.Background(), sql)
	if err != nil {
		log.Println("Failed to fetch topics, err=", err)
		c.JSON(http.StatusOK, gin.H{
			"pg":   1,
			"data": make([]interface{}, 0),
		})
		return
	}
	for rows.Next() {
		var row TopicWithAvatar
		err = rows.Scan(
			&row.Id, &row.Scheme, &row.Hostname, &row.Path, &row.Query,
			&row.CreatedOn, &row.UserId, &row.Uname, &row.AvatarCreatedOn,
			&row.Pebbles, &row.AvatarVerifiedOn, &row.Email,
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
			"id":        fmt.Sprintf("%x", v.Id.Bytes),
			"scheme":    v.Scheme.String,
			"hostname":  v.Hostname.String,
			"path":      v.Path.String,
			"query":     v.Query.String,
			"createdOn": v.CreatedOn.Time.Format(time.RFC3339),
			"userId":    fmt.Sprintf("%x", v.UserId.Bytes),
			"uname":     v.Uname.String,
			"pebbles":   v.Pebbles.Int,
		}
		if v.AvatarCreatedOn.Valid {
			m["avatarCreatedOn"] = v.AvatarCreatedOn.Time.Format(time.RFC3339)
		}
		if v.AvatarVerifiedOn.Valid {
			m["avatarVerifiedOn"] = v.AvatarVerifiedOn.Time.Format(time.RFC3339)
		}
		if v.Email.Valid {
			m["email"] = v.Email.String
		}
		result = append(result, m)
	}
	c.JSON(http.StatusOK, gin.H{
		"pg":   1,
		"data": result,
	})
}
