package pingpongh

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/azizbekdev-blog/go-monolithic/api/models"
	t "gitlab.com/azizbekdev-blog/go-monolithic/api/tokens"
	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage"
)

type PingI interface {
	Ping(c *gin.Context)
}

type PingHandler struct {
	log        *logger.Logger
	cfg        config.Config
	jwthandler t.JWTHandler
	postgres   storage.StorageI
}

func New(c *models.HandlerV1Config) PingI {
	return &PingHandler{
		log:        c.Logger,
		cfg:        c.Cfg,
		jwthandler: c.JWTHandler,
		postgres:   c.Postgres,
	}
}
