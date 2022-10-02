package m8e

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/config"
	"github.com/webcuss/webcuss/types"
	"log"
	"net/http"
	"strings"
)

// A11r Authenticator middleware
func A11r(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/sup") ||
			strings.HasPrefix(c.Request.URL.Path, "/sin") {
			c.Next()
			return
		}
		auth := c.GetHeader("Authorization")
		if strings.TrimSpace(auth) == "" {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		//TODO: Assert prefixed with `Bearer `
		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.GetSecret(), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			// get user info
			getUserSql := `
			SELECT a.id, a.uname, a."createdOn", a.pebbles, a."verifiedOn", a.email
			FROM avatar a
			WHERE a.id = $1;
			`
			var user types.Avatar
			err := db.QueryRow(context.Background(), getUserSql, claims["aud"]).
				Scan(&user.Id, &user.Uname, &user.CreatedOn, &user.Pebbles, &user.VerifiedOn, &user.Email)
			// TODO: Assert row > 0
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to get user info")
				c.Abort()
				return
			}
			// set user info
			c.Next()
		} else {
			log.Println(err)
			c.String(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
