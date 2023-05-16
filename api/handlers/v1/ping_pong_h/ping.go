package pingpongh

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errorh "gitlab.com/azizbekdev-blog/go-monolithic/api/handlers/v1/error_h"
	"gitlab.com/azizbekdev-blog/go-monolithic/api/models"
)

// @Summary		Ping pong
// @Description	Just to ensure that server is running
// @Tags		Ping
// @Produce		json
// @Success		200 				{object}  models.Response
// @Router		/ [get]
func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		ErrorCode:    errorh.ErrorSuccessCode,
		ErrorMessage: "Server is successfuly running",
	})
}
