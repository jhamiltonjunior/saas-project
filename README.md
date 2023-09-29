# My SaaS App

This is a SaaS project with a clean architecture, implemented in Go. The project is designed for long-term maintainability and scalability.

## Project Structure

The project has the following directory structure:

```
my-saas-app
├── cmd
│   ├── server
│   │   └── main.go
│   └── migrate
│       └── main.go
├── internal
│   ├── application
│   │   ├── usecases
│   │   └── services
│   ├── domain
│   │   ├── entities
│   │   ├── repositories
│   │   └── value-objects
│   ├── infrastructure
│   │   ├── database
│   │   ├── email
│   │   ├── logging
│   │   ├── payment
│   │   └── web
│   ├── interfaces
│   │   ├── api
│   │   ├── config
│   │   ├── controllers
│   │   ├── middleware
│   │   ├── presenters
│   │   └── routes
│   └── utils
├── pkg
│   ├── errors
│   ├── logger
│   ├── middleware
│   ├── server
│   └── validation
├── scripts
│   ├── db
│   │   ├── migrate-up.sh
│   │   └── migrate-down.sh
│   └── test
│       └── test.sh
├── tests
│   ├── application
│   ├── domain
│   ├── infrastructure
│   └── interfaces
├── go.mod
├── go.sum
├── .env
├── .gitignore
└── README.md
```

The project has the following files:

- `cmd/server/main.go`: This file is the entry point of the server application. It creates an instance of the server and starts it.
- `cmd/migrate/main.go`: This file is the entry point of the migration application. It creates an instance of the migration tool and runs it.
- `internal/application/usecases`: This directory contains the use cases of the application. Each use case is implemented as a function.
- `internal/application/services`: This directory contains the services of the application. Each service is implemented as a struct with methods.
- `internal/domain/entities`: This directory contains the entities of the domain. Each entity is implemented as a struct with properties and methods.
- `internal/domain/repositories`: This directory contains the repositories of the domain. Each repository is implemented as an interface with methods.
- `internal/domain/value-objects`: This directory contains the value objects of the domain. Each value object is implemented as a struct with properties and methods.
- `internal/infrastructure/database`: This directory contains the database infrastructure of the application. It exports functions to connect to the database and execute queries.
- `internal/infrastructure/email`: This directory contains the email infrastructure of the application. It exports functions to send emails.
- `internal/infrastructure/logging`: This directory contains the logging infrastructure of the application. It exports functions to log messages.
- `internal/infrastructure/payment`: This directory contains the payment infrastructure of the application. It exports functions to process payments.
- `internal/infrastructure/web`: This directory contains the web infrastructure of the application. It exports functions to start the server and handle requests.
- `internal/interfaces/api`: This directory contains the API interface of the application. It exports functions to handle API requests.
- `internal/interfaces/config`: This directory contains the configuration interface of the application. It exports functions to load the configuration from environment variables or files.
- `internal/interfaces/controllers`: This directory contains the controllers of the application. Each controller is implemented as a struct with methods to handle requests.
- `internal/interfaces/middleware`: This directory contains the middleware of the application. Each middleware is implemented as a function that takes a handler function and returns a new handler function.
- `internal/interfaces/presenters`: This directory contains the presenters of the application. Each presenter is implemented as a struct with methods to format data for the response.
- `internal/interfaces/routes`: This directory contains the routes of the application. Each route is implemented as a function that takes a router and sets up the handlers and middleware for the route.
- `internal/utils`: This directory contains utility functions used throughout the application.
- `pkg/errors`: This directory contains error types used throughout the application.
- `pkg/logger`: This directory contains the logger used throughout the application. It exports functions to log messages.
- `pkg/middleware`: This directory contains middleware functions used throughout the application.
- `pkg/server`: This directory contains the server implementation used throughout the application. It exports functions to start the server and handle requests.
- `pkg/validation`: This directory contains validation functions used throughout the application.
- `scripts/db/migrate-up.sh`: This file is a shell script to run the database migration tool to apply the migrations.
- `scripts/db/migrate-down.sh`: This file is a shell script to run the database migration tool to rollback the migrations.
- `scripts/test/test.sh`: This file is a shell script to run the tests for the application.
- `tests/application`: This directory contains the tests for the application use cases and services.
- `tests/domain`: This directory contains the tests for the domain entities and repositories.
- `tests/infrastructure`: This directory contains the tests for the infrastructure components of the application.
- `tests/interfaces`: This directory contains the tests for the application controllers, middleware, presenters, and routes.
- `go.mod`: This file is the module definition file for the application. It lists the dependencies and the module path.
- `go.sum`: This file contains the checksums of the dependencies used in the application.
- `.env`: This file contains the environment variables used in the application.
- `.gitignore`: This file lists the files and directories to ignore in the Git repository.
- `README.md`: This file contains the documentation for the project.

## Getting Started

To run the server application, run the following command:

```
go run cmd/server/main.go
```

To run the migration tool, run the following command:

```
go run cmd/migrate/main.go
```

To run the tests, run the following command:

```
./scripts/test/test.sh
```

## Configuration

The application can be configured using environment variables or configuration files. The configuration files should be placed in the `configs` directory.

The following environment variables can be used to configure the application:

- `APP_ENV`: The environment in which the application is running (e.g. `development`, `production`).
- `APP_PORT`: The port on which the server should listen (default: `8080`).
- `DB_HOST`: The hostname of the database server.
- `DB_PORT`: The port on which the database server is listening.
- `DB_NAME`: The name of the database.
- `DB_USER`: The username to use when connecting to the database.
- `DB_PASSWORD`: The password to use when connecting to the database.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.