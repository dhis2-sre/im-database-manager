package client

import (
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	jobClient "github.com/dhis2-sre/im-job/pkg/client"
)

func ProvideJobService(config config.Config) jobClient.Client {
	return jobClient.ProvideClient(config.JobService.Host, config.JobService.BasePath)
}
