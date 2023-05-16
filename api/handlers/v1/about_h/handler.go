package aboutsh

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/azizbekdev-blog/go-monolithic/api/models"
	t "gitlab.com/azizbekdev-blog/go-monolithic/api/tokens"
	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage"
)

type AboutI interface {
	Create(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type AboutHandler struct {
	log        *logger.Logger
	cfg        config.Config
	jwthandler t.JWTHandler
	postgres   storage.StorageI
}

func New(c *models.HandlerV1Config) AboutI {
	return &AboutHandler{
		log:        c.Logger,
		cfg:        c.Cfg,
		jwthandler: c.JWTHandler,
		postgres:   c.Postgres,
	}
}
