// Package classification Instance Database Manager Service.
//
// Database Manager Service is part of the Instance Manager environment
//
//	Version: 0.1.0
//	License: TODO
//	Contact: <info@dhis2.org> https://github.com/dhis2-sre/im-database-manager
//
//	Consumes:
//	  - application/json
//
//	Produces:
//	  - application/json
//
//	SecurityDefinitions:
//	  oauth2:
//	    type: oauth2
//	    tokenUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//	    refreshUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//	    flow: password
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"

	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/internal/server"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
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

	s3Config, err := s3config.LoadDefaultConfig(context.TODO(), s3config.WithRegion("eu-west-1"))
	if err != nil {
		return err
	}
	s3AWSClient := s3.NewFromConfig(s3Config)
	uploader := manager.NewUploader(s3AWSClient)
	s3Client := storage.NewS3Client(s3AWSClient, uploader)

	db, err := storage.NewDatabase(cfg)
	if err != nil {
		return err
	}
	dbRepo := database.NewRepository(db)
	dbSvc := database.NewService(cfg, usrSvc, s3Client, dbRepo)

	dbHandler := database.New(usrSvc, dbSvc, instanceSvc)

	authMiddleware, err := handler.NewAuthentication(cfg)
	if err != nil {
		return err
	}

	err = handler.RegisterValidation()
	if err != nil {
		return err
	}

	r := server.GetEngine(cfg.BasePath, dbHandler, authMiddleware)
	return r.Run()
}
