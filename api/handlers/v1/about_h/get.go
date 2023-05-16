package aboutsh

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	errorh "gitlab.com/azizbekdev-blog/go-monolithic/api/handlers/v1/error_h"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage/postgres/aboutrepo"
)

// @Summary		Get about by id
// @Tags        About
// @Description	Here about can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  FullResponse
// @Failure     default {object}  models.Response
// @Router		/about/{id} [GET]
func (h *AboutHandler) FindOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if errorh.HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.postgres.About().FindOne(context.Background(), &aboutrepo.FindOneReq{
		Id: id,
	})
	if errorh.HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.About().FindOne()") {
		return
	}

	c.JSON(http.StatusOK, FullResponse{
		ErrorCode:    errorh.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         (*FullResponseBody)(res),
	})
}
