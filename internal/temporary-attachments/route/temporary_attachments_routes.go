package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"clone_media/internal/temporary-attachments/controller"
)

func TempAttachmentsRoutes(r *gin.Engine, validate *validator.Validate) {
	tempAttachent := r.Group("/temp-attachments/")
	{
		tempAttachent.POST("", func(c *gin.Context) {
			controller.CreateTemporaryAttachments(c)
		})
	}
}
