package client

import (
	"fmt"
	"github.com/dhis2-sre/im-database-manager/internal/di"
	"github.com/dhis2-sre/im-database-manager/internal/server"
	"github.com/dhis2-sre/im-database-manager/swagger/sdk/models"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"net/url"
	"testing"
)

const AccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3OTYzODQwNzcsImlhdCI6MTY0MDY5MTQ3NywidXNlciI6eyJJRCI6NCwiQ3JlYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiVXBkYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiRGVsZXRlZEF0IjpudWxsLCJFbWFpbCI6InNvbWVvbmVAc29tZXRoaW5nLm9yZyIsIkdyb3VwcyI6bnVsbCwiQWRtaW5Hcm91cHMiOm51bGx9fQ.FnPIu36kV1T-Jix5Wy-HsZeqxQI6Q_7HQ14C1DWKHETIBSk-vLQ_sCMHVPKA42utEDFI3Xpmf6Gyzv9aPU_Cvg-JDazRprfrZBqn4LzSmT6K3HGoKoQ0b5G8exxz0Ote8NQDB1NBZmYvD1gpVVisCvzaewJTRAvRA3DS0n_O4kU5QENdLNPfWFo0rXOC83sLBsEIe2Ce4TiRrepOCSQKE-_rwQQSA3w30MhFmhAY7Ozcd9i69mtfcvqjORdNJ-zREgiw8B2g9oh7byE1h2oxjvoKC3WRfPeSYoRY6GuMHSSWJdzFKIswlZHdWU1GicPJASBbkKGbP5n5O6FXyeo0bw"

func TestFindDatabaseById(t *testing.T) {
	// TODO: fix test... im-user is needed
	return

	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	ts := httptest.NewServer(r)
	defer ts.Close()

	parsedUrl, err := url.Parse(ts.URL)
	assert.NoError(t, err)
	host := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())
	c := ProvideClient(host, environment.Config.BasePath)

	create, err := c.Create(AccessToken, &models.CreateDatabaseRequest{
		GroupID: 1,
		Name:    "whatever",
	})
	assert.NoError(t, err)

	db, err := c.FindById(AccessToken, uint(create.ID))
	log.Printf("%+v", db)
	//	assert.NoError(t, err)

	//	assert.Equal(t, uint64(1), u.ID)
}
