package authmgr

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/config"
	"github.com/webcuss/webcuss/types"
	"net/http"
	"time"
)

func SignUp(c *gin.Context, db *pgxpool.Pool) {
	var req types.SignUpReq
	err := c.Bind(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
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

	// begin txn
	txn, err := db.Begin(context.Background())
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot create db txn")
		return
	}

	// create user
	insertSql := `
	INSERT INTO avatar (uname, pword, "createdOn", pebbles)
	VALUES ($1, CRYPT($2, GEN_SALT('md5')), $3, 179)
	RETURNING id;
	`
	var userId string
	err = db.QueryRow(context.Background(), insertSql, req.Uname, req.Pword, time.Now().UTC()).Scan(&userId)
	if err == nil {
		// create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"aud": userId,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString(config.GetSecret())
		if err == nil {
			// commit txn
			err := txn.Commit(context.Background())
			if err == nil {
				c.JSON(http.StatusCreated, gin.H{
					"token": tokenString,
				})
				return
			} else {
				c.String(http.StatusInternalServerError, "Cannot commit db txn")
				return
			}
		}
	}

	_ = txn.Rollback(context.Background())
	c.String(http.StatusInternalServerError, "Something went wrong")
}
