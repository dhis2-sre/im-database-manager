module github.com/dhis2-sre/im-database-manager

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/config v1.11.1
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.7.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.22.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/dhis2-sre/im-job v0.1.0
	github.com/dhis2-sre/im-user v0.4.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-openapi/errors v0.20.1
	github.com/go-openapi/runtime v0.21.0
	github.com/go-openapi/strfmt v0.21.1
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.3
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/google/wire v0.5.0
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/lestrrat-go/jwx v1.2.14
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.4
)
