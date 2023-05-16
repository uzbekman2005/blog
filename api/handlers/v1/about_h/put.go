package aboutsh

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	errorh "gitlab.com/azizbekdev-blog/go-monolithic/api/handlers/v1/error_h"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage/postgres/aboutrepo"
)

// @Summary		Update about
// @Tags        About
// @Description	Here about can be Updated.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body      UpdateReq true "post info"
// @Success		200 	{object}  models.Response
// @Failure     default {object}  FullResponse
// @Router		/about [PUT]
func (h *AboutHandler) Update(c *gin.Context) {
	body := &UpdateReq{}
	err := c.ShouldBindJSON(&body)
	if errorh.HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.postgres.About().Update(context.Background(), (*aboutrepo.UpdateReq)(body))
	if errorh.HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.About().Update()") {
		return
	}

	c.JSON(http.StatusOK, &FullResponse{
		ErrorCode:    errorh.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         (*FullResponseBody)(res),
	})
}
