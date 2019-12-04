module github.com/kconde2/vote-app/api

go 1.13

// Generated by doing
// `go mod init github.com/kconde2/vote-app/api`
// `go get -u github.com/gin-gonic/gin`
// `go get -u github.com/jinzhu/gorm`

require (
	github.com/appleboy/gin-jwt/v2 v2.6.2
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/bearbin/go-age v0.0.0-20140407072555-316d0c1e7cd1
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/itsjamie/gin-cors v0.0.0-20160420130702-97b4a9da7933
	github.com/jinzhu/gorm v1.9.10
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/qor/validations v0.0.0-20171228122639-f364bca61b46
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/sys v0.0.0-20190919044723-0c1ff786ef13 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
