//+build wireinject

package di

import (
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	"github.com/google/wire"
	"gorm.io/gorm"
	"log"
)

type Environment struct {
	Config                   config.Config
	AuthenticationMiddleware handler.AuthenticationMiddleware
}

func ProvideEnvironment(
	config config.Config,
	authenticationMiddleware handler.AuthenticationMiddleware,
) Environment {
	return Environment{
		config,
		authenticationMiddleware,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		config.ProvideConfig,

		//		provideDatabase,

		//		client.ProvideUser,

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
