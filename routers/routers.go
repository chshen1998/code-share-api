package routers

import (
	"net/http"

	"gin_project/controllers"
	"gin_project/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	postService        services.PostService           = services.New()
	postController     controllers.PostController     = controllers.New(postService)
	responseService    services.ResponseService       = services.NewResponseService()
	responseController controllers.ResponseController = controllers.NewResponseController(responseService)
)

// Initializes all the endpoints in the rest api and set up CORS
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORS())

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Endpoints related to authenication
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.SignUp)
		authRoutes.POST("/login", controllers.Login)
		authRoutes.GET("/logout", controllers.Logout)
		authRoutes.GET("/get-session", controllers.GetSession)
	}

	// Endpoints that do not require user to be authenicated
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/posts", func(context *gin.Context) {
			context.JSON(http.StatusOK, postController.FindAll())
		})

		apiRoutes.GET("/posts/filter", func(context *gin.Context) {
			context.JSON(http.StatusOK, postController.FindFiltered(context))
		})

		apiRoutes.GET("/likes", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"count": responseController.GetLikes(context)})
		})

		apiRoutes.GET("/comments", func(context *gin.Context) {
			context.JSON(http.StatusOK, responseController.GetComments(context))
		})
	}

	// Endpointst that require user to be authenicated
	apiPrivateRoutes := r.Group("/api/private")
	apiPrivateRoutes.Use(controllers.AuthRequired)
	{
		apiPrivateRoutes.GET("/post", func(context *gin.Context) {
			context.JSON(http.StatusOK, postController.Get(context))
		})

		apiPrivateRoutes.POST("/post", func(context *gin.Context) {
			err := postController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusAccepted, postController.FindAll())
			}
		})

		apiPrivateRoutes.DELETE("/post", func(context *gin.Context) {
			err := postController.Delete(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusAccepted, postController.FindAll())
			}
		})

		apiPrivateRoutes.POST("/like", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"count": responseController.AddLike(context)})
		})

		apiPrivateRoutes.POST("/comment", func(context *gin.Context) {
			context.JSON(http.StatusOK, responseController.AddComment(context))
		})
	}

	return r
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
