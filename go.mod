module github.com/dhis2-sre/im-database-manager

go 1.16

require (
	github.com/aws/aws-sdk-go-v2/config v1.11.1
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.7.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.22.0
	github.com/dhis2-sre/im-user v0.3.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-openapi/runtime v0.21.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/google/wire v0.5.0
	github.com/lestrrat-go/jwx v1.2.14
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.3
)
