package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"clone_media/internal/role/controller"
)

func RoleRoutes(r *gin.Engine, validate *validator.Validate) {
	roles := r.Group("/role/")
	{
		roles.POST("", func(c *gin.Context) {
			controller.CreateRole(c, validate)
		})
	}
}
