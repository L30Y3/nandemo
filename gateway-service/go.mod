module github.com/L30Y3/nandemo/gateway-service

go 1.24.4

require (
	github.com/L30Y3/nandemo/shared v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.2.2
)

require (
	github.com/go-chi/cors v1.2.2
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/L30Y3/nandemo/shared => ../shared
