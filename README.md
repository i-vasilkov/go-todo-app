# go-todo-app

Todo application on Golang

## Build & Run (Locally)

### Prerequisites

- Docker
- Go (in project used v1.16)
- golangci-lint (<i>optional</i>, used to run code checks)

### Quick start

Copy the file `.env-example` to `.env` file and add values.

Run `make run` command for build&run application

### Postgres

If you want use project with PostgreSQL then:
- replace connection in /internal/app/app.go
- After `make run` command run `make migrate-up` command 
