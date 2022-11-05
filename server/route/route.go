package route

import (
	"net/http"

	"github.com/webcuss/webcuss/mgr/cmtmgr"
	"github.com/webcuss/webcuss/mgr/rctnmgr"
	"github.com/webcuss/webcuss/mgr/tpcmgr"

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
	r.Use(m8e.CORS())
	r.Use(m8e.A11r(dbConn))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Webcuss API! 🚀")
	})

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

	r.POST("/rctn/:commentId", func(c *gin.Context) {
		rctnmgr.PostReaction(c, dbConn)
	})

	r.GET("/rctn/:commentId", func(c *gin.Context) {
		rctnmgr.GetReaction(c, dbConn)
	})

	r.DELETE("/rctn/:commentId", func(c *gin.Context) {
		rctnmgr.DeleteReaction(c, dbConn)
	})
	r.GET("/support", func(c *gin.Context) {
		content := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas nec fringilla orci, sed imperdiet turpis. Sed ac semper magna. Sed malesuada fringilla commodo. Integer suscipit neque et mauris gravida, a interdum ante tincidunt. Nunc et mi sed turpis eleifend porttitor at eget justo. Sed quis elit dolor. Sed gravida erat eget nisi laoreet pretium. Vivamus hendrerit volutpat faucibus. Nullam vestibulum dui sed lectus bibendum sodales. Cras ac diam tristique, pulvinar mauris vitae, tempus tortor. Pellentesque elit eros, consequat sit amet quam at, dictum luctus leo. Proin varius dolor diam, eget accumsan turpis accumsan a. Nulla facilisi. Aenean a urna augue. Suspendisse ultrices odio non turpis accumsan vehicula.

Nullam auctor, tellus vitae mollis suscipit, nibh arcu facilisis nulla, ac iaculis metus neque eget lacus. Vivamus luctus risus vel lorem mollis, a lacinia neque cursus. Morbi ultrices risus ut orci congue sodales. Sed dignissim id sapien at pretium. Donec tristique, diam sed blandit condimentum, orci urna luctus dolor, eu tempus tortor nibh ut urna. Sed quis augue vitae ex dictum molestie vitae quis est. Proin posuere eros in metus porttitor, sed tincidunt diam hendrerit. Integer varius, dolor ullamcorper iaculis scelerisque, elit ex scelerisque metus, tempus congue velit nisi nec justo. Morbi sapien eros, rutrum ac scelerisque lacinia, molestie ac justo. Cras in vehicula augue. Suspendisse sem neque, rutrum non neque eu, mollis viverra urna. Proin id nunc at augue egestas accumsan. Cras semper odio ut nibh malesuada iaculis. Suspendisse et eleifend erat.

Donec bibendum luctus viverra. Nunc congue erat risus, at tristique erat interdum vitae. Nam rutrum ultrices ullamcorper. Curabitur quis fermentum quam. Etiam bibendum urna ac purus aliquet eleifend. Ut consequat leo eu nisl ornare, eu pharetra diam feugiat. Suspendisse luctus, turpis id lacinia hendrerit, quam lorem dignissim quam, sed semper nulla tortor at dolor. Morbi consectetur euismod rhoncus. Suspendisse potenti. Etiam mattis placerat commodo. Donec hendrerit et diam ac eleifend. Curabitur convallis ullamcorper sem, ut gravida lectus suscipit et.

Cras libero eros, accumsan vel interdum cursus, vulputate nec neque. Sed sodales odio tortor, non tempor felis posuere in. Suspendisse egestas cursus ullamcorper. Morbi eu lorem a nibh porta commodo a id nulla. In tortor neque, tristique at mi id, ornare luctus lacus. Phasellus gravida velit in ornare condimentum. Donec porttitor scelerisque orci, a pulvinar odio iaculis sit amet. Maecenas non velit eget sapien eleifend ornare. Sed luctus pulvinar imperdiet. Morbi sed nulla iaculis, laoreet felis eu, ultricies magna. Integer sed risus eget turpis ultrices dapibus id vitae nunc. Nulla at arcu ultrices, maximus ante vel, tempor turpis. Etiam fringilla, lectus eget vulputate accumsan, est sem dignissim risus, eget efficitur libero turpis vel leo. Aliquam et turpis sit amet nulla fringilla cursus at in nulla. Pellentesque a fringilla ipsum, eu convallis ipsum. Suspendisse purus orci, blandit eget pellentesque sit amet, venenatis interdum velit.

Sed gravida quis odio in convallis. Integer vitae vestibulum est, nec lacinia elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aenean vestibulum, urna in laoreet egestas, libero nisl volutpat odio, in dictum quam purus malesuada nunc. Phasellus id pharetra enim. Donec blandit, sapien non ultricies scelerisque, urna velit rhoncus leo, nec volutpat sapien enim a metus. Aliquam dapibus eleifend nisi, ut iaculis risus porta ut.
		`
		c.String(http.StatusOK, content)
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}
