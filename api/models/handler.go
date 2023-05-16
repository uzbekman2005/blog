package models

import (
	t "gitlab.com/azizbekdev-blog/go-monolithic/api/tokens"
	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage"
)

type HandlerV1Config struct {
	Logger     *logger.Logger
	Cfg        config.Config
	JWTHandler t.JWTHandler
	Postgres   storage.StorageI
}
