module github.com/fastenhealth/fasten-sources

go 1.18

require (
	github.com/fastenhealth/gofhir-models v0.0.6
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/go-playground/validator/v10 v10.16.0
	github.com/golang/mock v1.6.0
	github.com/samber/lo v1.35.0
	github.com/seborama/govcr v4.5.0+incompatible
	github.com/sirupsen/logrus v1.9.0
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/stretchr/testify v1.8.4
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17
	golang.org/x/net v0.17.0
	golang.org/x/oauth2 v0.2.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

//replace github.com/fastenhealth/gofhir-models => ../gofhir-models
