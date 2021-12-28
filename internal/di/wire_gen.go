// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/dhis2-sre/im-database-manager/internal/client"
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	"gorm.io/gorm"
	"log"
)

// Injectors from wire.go:

func GetEnvironment() Environment {
	configConfig := config.ProvideConfig()
	authenticationMiddleware := handler.ProvideAuthentication(configConfig)
	clientClient := client.ProvideUser(configConfig)
	s3Client := storage.ProvideS3Client()
	db := provideDatabase(configConfig)
	repository := database.ProvideRepository(db)
	service := database.ProvideService(configConfig, s3Client, repository)
	databaseHandler := database.ProvideHandler(clientClient, service)
	environment := ProvideEnvironment(configConfig, authenticationMiddleware, databaseHandler)
	return environment
}

// wire.go:

type Environment struct {
	Config                   config.Config
	AuthenticationMiddleware handler.AuthenticationMiddleware
	DatabaseHandler          database.Handler
}

func ProvideEnvironment(config2 config.Config,

	authenticationMiddleware handler.AuthenticationMiddleware,
	databaseHandler database.Handler,
) Environment {
	return Environment{config2, authenticationMiddleware,
		databaseHandler,
	}
}

func provideDatabase(c config.Config) *gorm.DB {
	database2, err := storage.ProvideDatabase(c)
	if err != nil {
		log.Fatalln(err)
	}
	return database2
}
