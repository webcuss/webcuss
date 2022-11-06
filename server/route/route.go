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
For any concern and support, please contact us at notaprefix@gmail.com
		`
		c.String(http.StatusOK, content)
	})

	r.GET("/privacy-policy", func(c *gin.Context) {
		content := `
<!DOCTYPE html>
<html>
<body>
	<h2>Data Privacy Policy</h2>
	<h4>
		Who We Are
	</h4>
	<p>
		This Privacy Policy explains how Webcuss ("Webcuss," "we," or "us") collects, uses, and discloses information
		about you. This Privacy Policy applies when you use our browser extensions that link to this Privacy Policy
		(collectively, our "Services"), contact us through our contacts found at the bottom of this page.
	</p>
	<p>
		We may change this Privacy Policy from time to time. If we make changes, we will notify you through the browser
		extension app by showing a notification inside the app. We encourage you to review this Privacy Policy regularly
		to stay informed about our information practices and the choices available to you.
	</p>
	<h4>
		Collection of Personal Data
	</h4>
	<p>
		We collect information you provided to us e.g., when you create an account by filling out registration form,
		submit or post content through our Services, or communicate with us. The types of personal information we may
		collect include your email address, username, your content (comments, replies, and reactions), and browser
		activity (website title and URL).
	</p>
	<p>
		We DO NOT collect data that are not mentioned above. You will find the data being collected and processed by
		directly looking into the source code of this project on GitHub (https://github.com/webcuss/webcuss)
	</p>
	<h4>
		Use of Personal Data
	</h4>
	<p>
		We use the information we collect to provide, maintain, and improve our Services, which includes publishing and
		distributing user-generated contents. We also use the information we collect to:
	<ol>
		<li>Create and maintain your Webcuss account.</li>
		<li>Send you technical notices and security alert messages.</li>
		<li>Debug to identify and repair errors in our Services.</li>
	</ol>
	</p>
	<h4>
		Sharing of Personal Data
	</h4>
	<ol>
		<li>Your data is only used by Webcuss and its Services.</li>
		<li>We DO NOT share or sell your information to third-parties.</li>
		<li>All your information is stored in our cloud service provider (AWS) with only the technical staff (developers and
			system admins) have access to.</li>
		<li>We may disclose personal information if we believe that disclosure is in accordance with, or required by, any
			applicable law or legal process, including lawful requests by public authorities to meet national security or law
			enforcement requirements.</li>
	</ol>
	<p>
	</p>
	<h4>
		Retention of Personal Data
	</h4>
	<p>
		We keep your personal data associated with your account for as long as your account remain active. You can request
		to delete your account by contacting us (we may ask for additional informations to verify account ownership), see
		contacts at the bottom of this page.
	</p>
	<h4>
		Contact Us
	</h4>
	<p>
		You can reach us on the following email
		<ul>
			<li>notaprefix@gmail.com</li>
		</ul>
	</p>
</body>
</html>
		`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, content)
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}
