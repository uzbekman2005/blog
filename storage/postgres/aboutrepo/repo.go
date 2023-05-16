package aboutrepo

import (
	"context"
	"time"

	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/db"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type AboutI interface {
	Create(ctx context.Context, req *CreateReq) (*FullResponse, error)
	FindOne(ctx context.Context, req *FindOneReq) (*FullResponse, error)
	Update(ctx context.Context, req *UpdateReq) (*FullResponse, error)
	Delete(ctx context.Context, req *DeleteReq) error
}

type AboutRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg *config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) AboutI {
	return &AboutRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}
