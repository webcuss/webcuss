package authmgr

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/config"
	"github.com/webcuss/webcuss/types"
)

func createAuthToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": strings.ReplaceAll(userId, "-", ""),
		"exp": time.Now().Add(time.Hour * 8760).Unix(), //8760 total hours of a year
	})
	tokenString, err := token.SignedString(config.GetSecret())
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

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
		tokenString, err := createAuthToken(userId)
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

func SignIn(c *gin.Context, db *pgxpool.Pool) {
	var req types.SignInReq
	err := c.Bind(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// db login
	var userId pgtype.UUID
	q := `
	SELECT a.Id
	FROM avatar a
	WHERE a.uname = $1
		AND a.pword = crypt($2, a.pword);
	`
	err = db.QueryRow(context.Background(), q, req.Uname, req.Pword).Scan(&userId)
	if err != nil {
		c.String(http.StatusUnauthorized, "Incorrect credentials")
		return
	}
	tokenString, err := createAuthToken(fmt.Sprintf("%x", userId.Bytes))
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
