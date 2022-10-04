package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/m8e"
	"github.com/webcuss/webcuss/mgr/authmgr"
	"net/http"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// middleware
	r.Use(m8e.A11r(db))

	r.POST("/sup", func(c *gin.Context) {
		authmgr.SignUp(c, db)
	})

	r.POST("/sin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	})

	r.POST("/sout", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	r.GET("/tpc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pg":   0,
			"data": make([]interface{}, 0),
		})
	})

	r.POST("/tpc", func(c *gin.Context) {
		c.String(http.StatusCreated, "Ok")
	})

	r.POST("/tpc/:topicId/cmt", func(c *gin.Context) {
		c.String(http.StatusCreated, "Ok")
	})

	r.GET("/tpc/:topicId/cmt", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":  "",
			"url": "",
			"pg":  0,
			"data": []gin.H{
				{
					"id":      "",
					"comment": "",
					"user": gin.H{
						"id":   "",
						"name": "",
					},
				},
			},
		})
	})

	r.POST("/cmt/:commentId", func(c *gin.Context) {
		c.String(http.StatusCreated, "Ok")
	})

	r.GET("/cmt/:commentId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pg": 0,
			"data": []gin.H{
				{
					"comment": "",
					"user":    gin.H{},
				},
			},
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}
