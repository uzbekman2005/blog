package storage

import (
	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/db"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage/postgres/aboutrepo"
)

type StorageI interface {
	About() aboutrepo.AboutI
}

type StoragePg struct {
	aboutRepo aboutrepo.AboutI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) StorageI {
	return &StoragePg{
		aboutRepo: aboutrepo.New(db, log, cfg),
	}
}

func (s *StoragePg) About() aboutrepo.AboutI {
	return s.aboutRepo
}
