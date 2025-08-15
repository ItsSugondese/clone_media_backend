package route

import (
	"clone_media/internal/user/controller"
	authentication_middleware "clone_media/pkg/middleware/authentication-middleware"
	paseto_token "clone_media/pkg/utils/paseto-token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoutes(r *gin.Engine, validate *validator.Validate) {
	users := r.Group("/user/")
	{
		users.POST("", func(c *gin.Context) {
			controller.RegisterUser(c, validate)
		})
		users.Use(authentication_middleware.PasetoAuthMiddleware(*paseto_token.TokenMaker))
		users.PUT("", func(c *gin.Context) {
			controller.RegisterUser(c, validate)
		})
		users.GET("doc/:id", controller.GetUserImage)
	}
}
