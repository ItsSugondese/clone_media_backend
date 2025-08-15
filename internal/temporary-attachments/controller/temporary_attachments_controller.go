package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	response_crud_enum "clone_media/enums/interface-enums/response/response-crud-enum"
	localization_enums "clone_media/enums/struct-enums/localization-enums"
	"clone_media/enums/struct-enums/project_module"
	generic_controller "clone_media/generics/generic-controller"
	"clone_media/internal/temporary-attachments/dto"
	"clone_media/internal/temporary-attachments/service"
	"clone_media/pkg/common/localization"
)

// @Summary Get list of attachment id
// @Schemes
// @Description do ping
// @Tags Temporary Attachments
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} int
// @Router /temporary-attachments [post]
// post /temporary-attachments
func CreateTemporaryAttachments(c *gin.Context) {
	var attachmentsDetailRequest dto.TemporaryAttachmentsDetailRequest
	if err := c.ShouldBind(&attachmentsDetailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the uploaded file

	// calling service

	savedData := service.SaveTemporaryAttachmentsService(c)

	generic_controller.GenericControllerSuccessResponseHandler(c,
		localization.GetLocalizedMessage(localization_enums.MessageCodeEnums.API_OPERATION, map[string]interface{}{
			"First":  project_module.ModuleNameEnums.BASE_USER,
			"Second": strings.ToLower(response_crud_enum.Create().String()),
		}), savedData)
}
