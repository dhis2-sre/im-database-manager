// Package classification Instance Database Manager Service.
//
// Database Manager Service as part of the Instance Manager environment
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk we take no responsibility
//
//    Version: 0.1.0
//    License: TODO
//    Contact: <info@dhis2.org> https://github.com/dhis2-sre/im-database-manager
//
//    Consumes:
//      - application/json
//
//    Produces:
//      - application/json
//
//    SecurityDefinitions:
//      oauth2:
//        type: oauth2
//        tokenUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//        refreshUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//        flow: password
// swagger:meta
package main

import (
	"fmt"
	"os"

	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/internal/server"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	jobClient "github.com/dhis2-sre/im-job/pkg/client"
	instanceClient "github.com/dhis2-sre/im-manager/pkg/client"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err) // nolint:errcheck
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("error getting config: %v", err)
	}

	usrSvc := userClient.New(cfg.UserService.Host, cfg.UserService.BasePath)
	instanceSvc := instanceClient.New(cfg.InstanceService.Host, cfg.InstanceService.BasePath)

	s3Client, err := storage.NewS3Client()
	if err != nil {
		return err
	}

	jobSvc := jobClient.ProvideClient(cfg.JobService.Host, cfg.JobService.BasePath)

	db, err := storage.NewDatabase(cfg)
	if err != nil {
		return err
	}
	dbRepo := database.NewRepository(db)
	dbSvc := database.NewService(cfg, usrSvc, s3Client, jobSvc, dbRepo)

	dbHandler := database.New(usrSvc, dbSvc, instanceSvc)

	authMiddleware, err := handler.NewAuthentication(cfg)
	if err != nil {
		return err
	}

	r := server.GetEngine(cfg.BasePath, dbHandler, authMiddleware)
	return r.Run()
}
