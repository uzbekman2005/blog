package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"gitlab.com/azizbekdev-blog/go-monolithic/api"
	"gitlab.com/azizbekdev-blog/go-monolithic/config"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/db"
	"gitlab.com/azizbekdev-blog/go-monolithic/pkg/logger"
	"gitlab.com/azizbekdev-blog/go-monolithic/storage"
)

func main() {
	var (
		casbinEnforcer *casbin.Enforcer
	)
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Error("casbin enforcer error", err)
		return
	}
	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", err)
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	dbConn, err := db.New(pgxUrl)
	if err != nil {
		panic(err)
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		CasbinEnforcer: casbinEnforcer,
		Postgres:       storage.New(dbConn, log, &cfg),
	})
	fmt.Println(cfg.PostgresDatabase, cfg.PostgresHost, cfg.PostgresPassword, cfg.PostgresUser, cfg.PostgresPort)
	if err := server.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", err)
		panic(err)
	}
}
