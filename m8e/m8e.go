package m8e

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/config"
	"github.com/webcuss/webcuss/types"
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
		// assert prefixed with `Bearer `
		if !strings.HasPrefix(auth, "Bearer ") {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		authToken := strings.Split(auth, " ")[1]
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.GetSecret(), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// get user info
			q := `
			SELECT a.id, a.uname, a."createdOn", a.pebbles, a."verifiedOn", a.email
			FROM avatar a
			WHERE a.id = $1;
			`
			var usr types.Avatar
			err := db.QueryRow(context.Background(), q, claims["aud"]).
				Scan(&usr.Id, &usr.Uname, &usr.CreatedOn, &usr.Pebbles, &usr.VerifiedOn, &usr.Email)
			if err != nil {
				// no result
				c.String(http.StatusUnauthorized, "Unauthorized")
				c.Abort()
				return
			}
			// set user info
			c.Set("user", usr)
			c.Next()
		} else {
			log.Println(err) // could be token expired
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}
	}
}
