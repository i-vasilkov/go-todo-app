build:
	go mod download\
	&& GOOS=linux CGO_ENABLED=0 go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans app

swag:
	swag init -g internal/app/app.go
