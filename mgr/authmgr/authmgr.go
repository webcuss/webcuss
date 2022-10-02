package authmgr

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/types"
	"net/http"
	"time"
)

func SignUp(c *gin.Context, db *pgxpool.Pool) {
	var req types.SignUpReq
	err := c.Bind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// check if username taken
	doesExistSql := `
	SELECT EXISTS (
	    SELECT a.id
	    FROM avatar a
	    WHERE a.uname = $1
	);
	`
	var taken bool
	err = db.QueryRow(context.Background(), doesExistSql, req.Uname).Scan(&taken)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if taken {
		c.String(http.StatusConflict, "Username already taken")
		return
	}

	// create user
	insertSql := `
	INSERT INTO avatar (uname, pword, "createdOn", pebbles)
	VALUES ($1, CRYPT($2, GEN_SALT('md5')), $3, 179);
	`
	_, err = db.Exec(context.Background(), insertSql, req.Uname, req.Pword, time.Now().UTC())
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": "********",
	})
}
