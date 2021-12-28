//+build wireinject

package di

import (
	"github.com/dhis2-sre/im-database-manager/internal/client"
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	"github.com/google/wire"
	"gorm.io/gorm"
	"log"
)

type Environment struct {
	Config                   config.Config
	AuthenticationMiddleware handler.AuthenticationMiddleware
	DatabaseHandler          database.Handler
}

func ProvideEnvironment(
	config config.Config,
	authenticationMiddleware handler.AuthenticationMiddleware,
	databaseHandler database.Handler,
) Environment {
	return Environment{
		config,
		authenticationMiddleware,
		databaseHandler,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		config.ProvideConfig,

		client.ProvideUser,

		provideDatabase,
		storage.ProvideS3Client,

		database.ProvideRepository,
		database.ProvideService,
		database.ProvideHandler,

		handler.ProvideAuthentication,

		ProvideEnvironment,
	)
	return Environment{}
}

func provideDatabase(c config.Config) *gorm.DB {
	database, err := storage.ProvideDatabase(c)
	if err != nil {
		log.Fatalln(err)
	}
	return database
}
