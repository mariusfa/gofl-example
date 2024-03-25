module todo

go 1.22.0

// For local dev
replace github.com/mariusfa/gofl => ../../gofl

require (
	github.com/golang-migrate/migrate/v4 v4.17.0
	github.com/google/uuid v1.4.0
	github.com/lib/pq v1.10.9
	github.com/mariusfa/gofl/v2 v2.1.1
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
)
