package route

import (
	"github.com/webcuss/webcuss/mgr/cmtmgr"
	"github.com/webcuss/webcuss/mgr/tpcmgr"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webcuss/webcuss/m8e"
	"github.com/webcuss/webcuss/mgr/authmgr"
)

func SetupRouter(dbConn *pgxpool.Pool) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// middleware
	r.Use(m8e.A11r(dbConn))

	r.POST("/sup", func(c *gin.Context) {
		authmgr.SignUp(c, dbConn)
	})

	r.POST("/sin", func(c *gin.Context) {
		authmgr.SignIn(c, dbConn)
	})

	r.POST("/sout", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	r.GET("/tpc", func(c *gin.Context) {
		tpcmgr.GetTopic(c, dbConn)
	})

	r.POST("/tpc", func(c *gin.Context) {
		tpcmgr.PostTopic(c, dbConn)
	})

	r.POST("/tpc/:topicId/cmt", func(c *gin.Context) {
		cmtmgr.PostComment(c, dbConn)
	})

	r.GET("/tpc/:topicId/cmt", func(c *gin.Context) {
		cmtmgr.GetComments(c, dbConn)
	})

	r.POST("/cmt/:commentId", func(c *gin.Context) {
		cmtmgr.PostReply(c, dbConn)
	})

	r.GET("/cmt/:commentId", func(c *gin.Context) {
		cmtmgr.GetReplies(c, dbConn)
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}
