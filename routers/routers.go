package routers

import (
	"net/http"

	"gin_project/controllers"
	"gin_project/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	postService        services.PostService           = services.New()
	postController     controllers.PostController     = controllers.New(postService)
	responseService    services.ResponseService       = services.NewResponseService()
	responseController controllers.ResponseController = controllers.NewResponseController(responseService)
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.SignUp)
		authRoutes.POST("/login", controllers.Login)
		authRoutes.GET("/logout", controllers.Logout)
		authRoutes.GET("/get-session", controllers.GetSession)
	}

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
